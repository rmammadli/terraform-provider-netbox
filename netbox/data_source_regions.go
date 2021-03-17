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

func dataSourceDcimRegions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDcimRegionsRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent": {
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

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent": {
							Type:     schema.TypeInt,
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

func dataSourceDcimRegionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &dcim.DcimRegionsListParams{
		Context: ctx,
	}

	if val, ok := d.GetOk("description"); ok {
		description := val.(string)
		params.Description = &description
	}

	if val, ok := d.GetOk("name"); ok {
		name := val.(string)
		params.Name = &name
	}

	if val, ok := d.GetOk("id"); ok {
		id := val.(string)
		params.ID = &id
	}

	if val, ok := d.GetOk("parent"); ok {
		parent := val.(string)
		params.Parent = &parent
	}

	if val, ok := d.GetOk("slug"); ok {
		slug := val.(string)
		params.Slug = &slug
	}

	r, err := c.Dcim.DcimRegionsList(params, nil)
	if err != nil {
		return diag.Errorf("Could not get the list of regions: %v", err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenDcimRegionsResults(r.Payload.Results))

	return diags
}

func flattenDcimRegionsResults(input []*models.Region) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, value := range input {
		values := make(map[string]interface{})

		values["id"] = value.ID
		values["description"] = value.Description
		values["name"] = value.Name
		values["parent"] = value.Parent
		values["slug"] = value.Slug

		result = append(result, values)
	}
	return result
}
