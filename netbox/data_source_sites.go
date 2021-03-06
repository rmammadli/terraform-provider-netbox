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

func dataSourceDcimInterfaces() *schema.Resource {
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
					},
				},
			},
		},
	}
}

func dataSourceDcimInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &dcim.DcimInterfacesListParams{
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

	r, err := c.Dcim.DcimInterfacesList(params, nil)
	if err != nil {
		return diag.Errorf("Could not get the list of interfaces: %v", err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenDcimInterfacesResults(r.Payload.Results))

	return diags
}

func flattenDcimInterfacesResults(input []*models.Interface) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, value := range input {
		values := make(map[string]interface{})

		values["id"] = value.ID
		values["name"] = value.Name

		result = append(result, values)
	}
	return result
}
