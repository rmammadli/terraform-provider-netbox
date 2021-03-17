package netbox

import (
	"context"
	"fmt"
	"net/url"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
)

//Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"netbox_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_HOST", nil),
				Description: "Netbox host",
			},

			"netbox_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_TOKEN", nil),
				Description: "Netbox API token",
			},

			"connection_type": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONNECTION_TYPE", "https"),
				Description: "Connection type used for transport",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"netbox_dcim_region":  dataSourceDcimRegion(),
			"netbox_dcim_regions": dataSourceDcimRegions(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"netbox_dcim_region": resourceDcimRegion(),
			"netbox_dcim_rack":   resourceDcimRack(),
			"netbox_dcim_site":   resourceDcimSite(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	netboxHost := d.Get("netbox_host").(string)
	netboxToken := d.Get("netbox_token").(string)
	connType := d.Get("connection_type").(string)

	var diags diag.Diagnostics

	if (netboxHost == "") || (netboxToken == "") || (connType == "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Empty netbox host or password or connection type",
			Detail:   "Unable to auth user for authenticated netbox client",
		})
	}

	_, err := url.Parse(netboxHost)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to parse netbox host",
			Detail:   "Unable to auth user for authenticated netbox client",
		})
		return nil, diags
	}

	transport := runtimeclient.New(netboxHost, client.DefaultBasePath, []string{"http"})
	transport.Transport = logging.NewTransport("Netbox", transport.Transport)

	if netboxToken != "" {
		transport.DefaultAuthentication = runtimeclient.APIKeyAuth("Authorization", "header",
			fmt.Sprintf("Token %v", netboxToken))
	}

	return client.New(transport, strfmt.Default), diags
}
