package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func dataSourceDcimRegion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDcimRegionRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:     schema.TypeInt,
				Required: true,
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
	}
}

func dataSourceDcimRegionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &dcim.DcimRegionsReadParams{
		Context: ctx,
		ID:      int64(d.Get("region_id").(int)),
	}

	r, err := c.Dcim.DcimRegionsRead(params, nil)
	if err != nil {
		return diag.Errorf("Could not get a region, please check if provided region_id exists in your netbox: %v", err)
	}

	d.SetId(strconv.FormatInt(r.Payload.ID, 10))
	d.Set("description", r.Payload.Description)
	d.Set("name", r.Payload.Name)
	d.Set("parent", r.Payload.Parent)
	d.Set("slug", r.Payload.Slug)

	return diags
}
