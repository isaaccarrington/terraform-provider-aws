// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_directory_service_directory": DataSourceDirectory,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_directory_service_conditional_forwarder":     ResourceConditionalForwarder,
		"aws_directory_service_directory":                 ResourceDirectory,
		"aws_directory_service_log_subscription":          ResourceLogSubscription,
		"aws_directory_service_radius_settings":           ResourceRadiusSettings,
		"aws_directory_service_region":                    ResourceRegion,
		"aws_directory_service_shared_directory":          ResourceSharedDirectory,
		"aws_directory_service_shared_directory_accepter": ResourceSharedDirectoryAccepter,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DS
}

var ServicePackage = &servicePackage{}
