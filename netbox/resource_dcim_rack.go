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

func resourceDcimRack() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimRackCreate,
		ReadContext:   resourceDcimRackRead,
		UpdateContext: resourceDcimRackUpdate,
		DeleteContext: resourceDcimRackDelete,

		Schema: map[string]*schema.Schema{
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"desc_units": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"device_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"facility_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeInt,
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
			"outer_depth": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Enum: [mm in]
			"outer_unit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"outer_width": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"powerfeed_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// Enum: [reserved available planned active deprecated]
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"color": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tenant": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Enum: [2-post-frame 4-post-frame 4-post-cabinet wall-frame wall-cabinet]
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"u_height": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Enum: [10 19 21 23]
			"width": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}

}

func resourceDcimRackCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	site := int64(d.Get("site").(int))

	params := &dcim.DcimRacksCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableRack{
		Name: &name,
		Site: &site,
	}

	if val, ok := d.GetOk("asset_tag"); ok {
		assetTag := val.(string)
		params.Data.AssetTag = &assetTag
	}

	if val, ok := d.GetOk("comments"); ok {
		params.Data.Comments = val.(string)
	}

	if val, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = val.(map[string]interface{})
	}

	if val, ok := d.GetOk("desc_units"); ok {
		params.Data.DescUnits = val.(bool)
	}

	if val, ok := d.GetOk("facility_id"); ok {
		facilityID := val.(string)
		params.Data.FacilityID = &facilityID
	}

	if val, ok := d.GetOk("group"); ok {
		group := int64(val.(int))
		params.Data.Group = &group
	}

	if val, ok := d.GetOk("outer_depth"); ok {
		outerDepth := int64(val.(int))
		params.Data.OuterDepth = &outerDepth
	}

	if val, ok := d.GetOk("outer_unit"); ok {
		params.Data.OuterUnit = val.(string)
	}

	if val, ok := d.GetOk("outer_width"); ok {
		outerWidth := int64(val.(int))
		params.Data.OuterWidth = &outerWidth
	}

	if val, ok := d.GetOk("role"); ok {
		role := int64(val.(int))
		params.Data.Role = &role
	}

	if val, ok := d.GetOk("serial"); ok {
		params.Data.Serial = val.(string)
	}

	if val, ok := d.GetOk("status"); ok {
		params.Data.Status = val.(string)
	}

	if val, ok := d.GetOk("tags"); ok {
		params.Data.Tags = expandTags(val.([]interface{}))
	}

	if val, ok := d.GetOk("tenant"); ok {
		tenant := int64(val.(int))
		params.Data.Tenant = &tenant
	}

	if val, ok := d.GetOk("type"); ok {
		params.Data.Type = val.(string)
	}

	if val, ok := d.GetOk("u_height"); ok {
		params.Data.UHeight = int64(val.(int))
	}

	if val, ok := d.GetOk("width"); ok {
		params.Data.Width = int64(val.(int))
	}

	r, err := c.Dcim.DcimRacksCreate(params, nil)
	if err != nil {
		return diag.Errorf("Could not create a rack: %v", err)
	}

	d.SetId(strconv.FormatInt(r.Payload.ID, 10))

	resourceDcimRackRead(ctx, d, m)

	return diags
}

func resourceDcimRackRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	params := &dcim.DcimRacksReadParams{
		Context: ctx,
		ID:      id,
	}

	r, err := c.Dcim.DcimRacksRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Could not get a rack: %v", err)
	}

	if r.Payload.AssetTag != nil {
		d.Set("asset_tag", r.Payload.AssetTag)
	}

	d.Set("comments", r.Payload.Comments)
	d.Set("custom_fields", r.Payload.CustomFields)
	d.Set("desc_units", r.Payload.DescUnits)
	d.Set("device_count", r.Payload.DeviceCount)
	d.Set("display_name", r.Payload.DisplayName)

	if r.Payload.FacilityID != nil {
		d.Set("facility_id", r.Payload.FacilityID)
	}

	if r.Payload.Group != nil {
		d.Set("group", r.Payload.Group)
	}

	d.Set("id", r.Payload.ID)

	if r.Payload.Name != nil {
		d.Set("name", r.Payload.Name)
	}

	if r.Payload.OuterDepth != nil {
		d.Set("outer_depth", r.Payload.OuterDepth)
	}

	d.Set("outer_unit", r.Payload.OuterUnit)

	if r.Payload.OuterWidth != nil {
		d.Set("outer_width", r.Payload.OuterWidth)
	}

	d.Set("powerfeed_count", r.Payload.PowerfeedCount)

	if r.Payload.Role != nil {
		d.Set("role", r.Payload.Role)
	}

	d.Set("serial", r.Payload.Serial)

	if r.Payload.Site != nil {
		d.Set("site", r.Payload.Site)
	}

	d.Set("status", r.Payload.Status)
	d.Set("tags", flattenTags(r.Payload.Tags))

	if r.Payload.Tenant != nil {
		d.Set("tenant", r.Payload.Tenant.ID)
	}

	if &r.Payload.UHeight != nil {
		d.Set("u_height", r.Payload.UHeight)
	}

	if r.Payload.Width != nil {
		d.Set("width", r.Payload.Width)
	}

	return diags
}

func resourceDcimRackUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse rack id: %v", err)
	}

	params := &dcim.DcimRacksPartialUpdateParams{
		Context: ctx,
		ID:      id,
	}

	name := d.Get("name").(string)
	site := int64(d.Get("site").(int))

	params.Data = &models.WritableRack{
		Name: &name,
		Site: &site,
	}

	if d.HasChange("asset_tag") {
		assetTag := d.Get("asset_tag").(string)
		params.Data.AssetTag = &assetTag
	}

	if d.HasChange("comments") {
		params.Data.Comments = d.Get("comments").(string)
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	if d.HasChange("desc_units") {
		params.Data.DescUnits = d.Get("desc_units").(bool)
	}

	if d.HasChange("facility_id") {
		facilityID := d.Get("facility_id").(string)
		params.Data.FacilityID = &facilityID
	}

	if d.HasChange("group") {
		group := int64(d.Get("group").(int))
		params.Data.Group = &group
	}

	if d.HasChange("outer_depth") {
		outerDepth := int64(d.Get("outer_depth").(int))
		params.Data.OuterDepth = &outerDepth
	}

	if d.HasChange("outer_unit") {
		params.Data.OuterUnit = d.Get("outer_unit").(string)
	}

	if d.HasChange("outer_width") {
		outerWidth := int64(d.Get("outer_width").(int))
		params.Data.OuterWidth = &outerWidth
	}

	if d.HasChange("role") {
		role := int64(d.Get("role").(int))
		params.Data.Role = &role
	}

	if d.HasChange("serial") {
		params.Data.Serial = d.Get("serial").(string)
	}

	if d.HasChange("status") {
		params.Data.Serial = d.Get("status").(string)
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("tenant") {
		tenant := int64(d.Get("tenant").(int))
		params.Data.Tenant = &tenant
	}

	if d.HasChange("type") {
		params.Data.Type = d.Get("type").(string)
	}

	if d.HasChange("u_height") {
		params.Data.UHeight = int64(d.Get("u_height").(int))
	}

	if d.HasChange("width") {
		params.Data.UHeight = int64(d.Get("width").(int))
	}

	_, err = c.Dcim.DcimRacksPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Could not update rack: %v", err)
	}

	return resourceDcimRackRead(ctx, d, m)
}

func resourceDcimRackDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse rack ID: %v", err)
	}

	params := &dcim.DcimRacksDeleteParams{
		Context: ctx,
		ID:      id,
	}

	_, err = c.Dcim.DcimRacksDelete(params, nil)
	if err != nil {
		return diag.Errorf("Could not to delete rack: %v", err)
	}

	d.SetId("")

	return diags
}
