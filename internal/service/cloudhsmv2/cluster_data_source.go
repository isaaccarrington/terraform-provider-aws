package cloudhsmv2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

// @SDKDataSource("aws_cloudhsm_v2_cluster")
func DataSourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceClusterRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"cluster_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"security_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_csr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_hardware_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hsm_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"manufacturer_hardware_certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"subnet_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func dataSourceClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).CloudHSMV2Conn()

	clusterId := d.Get("cluster_id").(string)
	filters := []*string{&clusterId}
	log.Printf("[DEBUG] Reading CloudHSM v2 Cluster %s", clusterId)
	result := int64(1)
	input := &cloudhsmv2.DescribeClustersInput{
		Filters: map[string][]*string{
			"clusterIds": filters,
		},
		MaxResults: &result,
	}
	state := d.Get("cluster_state").(string)
	states := []*string{&state}
	if len(state) > 0 {
		input.Filters["states"] = states
	}
	out, err := conn.DescribeClustersWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "describing CloudHSM v2 Cluster: %s", err)
	}

	var cluster *cloudhsmv2.Cluster
	for _, c := range out.Clusters {
		if aws.StringValue(c.ClusterId) == clusterId {
			cluster = c
			break
		}
	}

	if cluster == nil {
		return sdkdiag.AppendErrorf(diags, "cluster with id %s not found", clusterId)
	}

	d.SetId(clusterId)
	d.Set("vpc_id", cluster.VpcId)
	d.Set("security_group_id", cluster.SecurityGroup)
	d.Set("cluster_state", cluster.State)
	if err := d.Set("cluster_certificates", readClusterCertificates(cluster)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting cluster_certificates: %s", err)
	}

	var subnets []string
	for _, sn := range cluster.SubnetMapping {
		subnets = append(subnets, *sn)
	}

	if err := d.Set("subnet_ids", subnets); err != nil {
		return sdkdiag.AppendErrorf(diags, "[DEBUG] Error saving Subnet IDs to state for CloudHSM v2 Cluster (%s): %s", d.Id(), err)
	}

	return diags
}
