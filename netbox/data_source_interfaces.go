package netbox

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func dataSourceDcimSites() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDcimSitesRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDcimSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &dcim.DcimSitesListParams{
		Context: ctx,
	}

	if val, ok := d.GetOk("name"); ok {
		name := val.(string)
		params.Name = &name
	}

	if val, ok := d.GetOk("id"); ok {
		id := val.(string)
		params.ID = &id
	}

	if val, ok := d.GetOk("slug"); ok {
		slug := val.(string)
		params.Slug = &slug
	}

	r, err := c.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return diag.Errorf("Could not get the list of sites: %v", err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenDcimSitesResults(r.Payload.Results))

	return diags
}

func flattenDcimSitesResults(input []*models.Site) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, value := range input {
		values := make(map[string]interface{})

		values["id"] = value.ID
		values["name"] = value.Name
		values["slug"] = value.Slug

		result = append(result, values)
	}
	return result
}
