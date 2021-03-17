package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimSiteCreate,
		ReadContext:   resourceDcimSiteRead,
		UpdateContext: resourceDcimSiteUpdate,
		DeleteContext: resourceDcimSiteDelete,

		Schema: map[string]*schema.Schema{
			"asn": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"circuit_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_phone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"facility": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latitude": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"longitude": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"physical_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rack_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shipping_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Enum: [planned staging active decommissioning retired]
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
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Enum: [2-post-frame 4-post-frame 4-post-cabinet wall-frame wall-cabinet]
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtualmachine_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}

}

func resourceDcimSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimSitesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
	}

	if val, ok := d.GetOk("asn"); ok {
		asn := int64(val.(int))
		params.Data.Asn = &asn
	}

	if val, ok := d.GetOk("comments"); ok {
		params.Data.Comments = val.(string)
	}

	if val, ok := d.GetOk("contact_email"); ok {
		params.Data.ContactEmail = strfmt.Email(val.(string))
	}

	if val, ok := d.GetOk("contact_name"); ok {
		params.Data.ContactName = val.(string)
	}

	if val, ok := d.GetOk("contact_phone"); ok {
		params.Data.Comments = val.(string)
	}

	if val, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = val.(map[string]interface{})
	}

	if val, ok := d.GetOk("description"); ok {
		params.Data.Comments = val.(string)
	}

	if val, ok := d.GetOk("facility"); ok {
		params.Data.Facility = val.(string)
	}

	if val, ok := d.GetOk("latitude"); ok {
		latitude := val.(string)
		params.Data.Latitude = &latitude
	}

	if val, ok := d.GetOk("longitude"); ok {
		longitude := val.(string)
		params.Data.Longitude = &longitude
	}

	if val, ok := d.GetOk("region"); ok {
		region := int64(val.(int))
		params.Data.Region = &region
	}

	if val, ok := d.GetOk("shipping_address"); ok {
		params.Data.ShippingAddress = val.(string)
	}

	if val, ok := d.GetOk("physical_address"); ok {
		params.Data.PhysicalAddress = val.(string)
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

	if val, ok := d.GetOk("time_zone"); ok {
		params.Data.TimeZone = val.(string)
	}

	r, err := c.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Could not create a site: %v", err)
	}

	d.SetId(strconv.FormatInt(r.Payload.ID, 10))

	resourceDcimSiteRead(ctx, d, m)

	return diags
}

func resourceDcimSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	params := &dcim.DcimSitesReadParams{
		Context: ctx,
		ID:      id,
	}

	r, err := c.Dcim.DcimSitesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Could not get a site: %v", err)
	}

	if r.Payload.Asn != nil {
		d.Set("asn", r.Payload.Asn)
	}

	d.Set("comments", r.Payload.Comments)
	d.Set("contact_email", r.Payload.ContactEmail)
	d.Set("contact_name", r.Payload.ContactName)
	d.Set("contact_phone", r.Payload.ContactPhone)
	d.Set("custom_fields", r.Payload.CustomFields)
	d.Set("facility", r.Payload.Facility)
	d.Set("latitude", r.Payload.Latitude)
	d.Set("longitude", r.Payload.Longitude)
	d.Set("region", r.Payload.Region)
	d.Set("shipping_address", r.Payload.ShippingAddress)
	d.Set("physical_address", r.Payload.PhysicalAddress)
	d.Set("status", r.Payload.Status)
	d.Set("tags", flattenTags(r.Payload.Tags))
	d.Set("tenant", r.Payload.Tenant)
	d.Set("time_zone", r.Payload.TimeZone)

	return diags
}

func resourceDcimSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse site id: %v", err)
	}

	params := &dcim.DcimSitesPartialUpdateParams{
		Context: ctx,
		ID:      id,
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
	}

	if d.HasChange("asn") {
		asn := int64(d.Get("asn").(int))
		params.Data.Asn = &asn
	}

	if d.HasChange("comments") {
		params.Data.Comments = d.Get("comments").(string)
	}

	if d.HasChange("contact_email") {
		params.Data.ContactEmail = strfmt.Email(d.Get("contact_email").(string))
	}

	if d.HasChange("contact_name") {
		params.Data.ContactName = d.Get("contact_name").(string)
	}

	if d.HasChange("contact_phone") {
		params.Data.ContactPhone = d.Get("contact_phone").(string)
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("facility") {
		params.Data.Facility = d.Get("facility").(string)
	}

	if d.HasChange("latitude") {
		latitude := d.Get("latitude").(string)
		params.Data.Latitude = &latitude
	}

	if d.HasChange("longitude") {
		longitude := d.Get("longitude").(string)
		params.Data.Longitude = &longitude
	}

	if d.HasChange("region") {
		region := int64(d.Get("region").(int))
		params.Data.Region = &region
	}

	if d.HasChange("shipping_address") {
		params.Data.ShippingAddress = d.Get("shipping_address").(string)
	}

	if d.HasChange("physical_address") {
		params.Data.PhysicalAddress = d.Get("physical_address").(string)
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("tenant") {
		tenant := int64(d.Get("tenant").(int))
		params.Data.Tenant = &tenant
	}

	if d.HasChange("time_zone") {
		params.Data.TimeZone = d.Get("time_zone").(string)
	}

	_, err = c.Dcim.DcimSitesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Could not update site: %v", err)
	}

	return resourceDcimRackRead(ctx, d, m)
}

func resourceDcimSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse site ID: %v", err)
	}

	params := &dcim.DcimSitesDeleteParams{
		Context: ctx,
		ID:      id,
	}

	_, err = c.Dcim.DcimSitesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Could not to delete site: %v", err)
	}

	d.SetId("")

	return diags
}
