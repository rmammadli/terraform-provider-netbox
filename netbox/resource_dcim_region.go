package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimRegion() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimRegionCreate,
		ReadContext:   resourceDcimRegionRead,
		UpdateContext: resourceDcimRegionUpdate,
		DeleteContext: resourceDcimRegionDelete,

		Schema: map[string]*schema.Schema{
			"depth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func resourceDcimRegionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimRegionsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableRegion{
		Name: &name,
		Slug: &slug,
	}

	if val, ok := d.GetOk("description"); ok {
		params.Data.Description = val.(string)
	}

	if val, ok := d.GetOk("parent"); ok {
		parent := int64(val.(int))
		params.Data.Parent = &parent
	}

	r, err := c.Dcim.DcimRegionsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Could not create region: %v", err)
	}

	d.SetId(strconv.FormatInt(r.Payload.ID, 10))

	resourceDcimRegionRead(ctx, d, m)

	return diags
}

func resourceDcimRegionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	params := &dcim.DcimRegionsReadParams{
		Context: ctx,
		ID:      id,
	}

	r, err := c.Dcim.DcimRegionsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Could not get region: %v", err)
	}

	d.Set("depth", r.Payload.Depth)
	d.Set("description", r.Payload.Description)
	d.Set("id", r.Payload.ID)
	d.Set("name", r.Payload.Name)
	d.Set("parent", r.Payload.Parent)
	d.Set("site_count", r.Payload.SiteCount)
	d.Set("slug", r.Payload.Slug)
	d.Set("url", r.Payload.URL)

	return diags
}

func resourceDcimRegionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse id: %v", err)
	}

	params := &dcim.DcimRegionsPartialUpdateParams{
		Context: ctx,
		ID:      id,
	}

	name := d.Get("name").(string)

	params.Data = &models.WritableRegion{
		Name: &name,
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("parent") {
		parent := int64(d.Get("parent").(int))
		params.Data.Parent = &parent
	}

	if d.HasChange("site_count") {
		siteCount := int64(d.Get("site_count").(int))
		params.Data.SiteCount = siteCount
	}

	_, err = c.Dcim.DcimRegionsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Could not update region: %v", err)
	}

	return resourceDcimRegionRead(ctx, d, m)
}

func resourceDcimRegionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse ID: %v", err)
	}

	params := &dcim.DcimRegionsDeleteParams{
		Context: ctx,
		ID:      id,
	}

	_, err = c.Dcim.DcimRegionsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Could not to delete region: %v", err)
	}

	d.SetId("")

	return diags
}
