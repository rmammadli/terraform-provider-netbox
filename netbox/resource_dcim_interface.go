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

func resourceDcimInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimInterfaceCreate,
		ReadContext:   resourceDcimInterfaceRead,
		UpdateContext: resourceDcimInterfaceUpdate,
		DeleteContext: resourceDcimInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"connected_endpoint_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conection_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"count_ipaddresses": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lag": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"untagged_vlan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}

}

func resourceDcimInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	device := int64(d.Get("device").(int))
	interface_type := d.Get("type").(string)

	params := &dcim.DcimInterfacesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableInterface{
		Name:        &name,
		Device:      &device,
		Type:        &interface_type,
		TaggedVlans: []int64{},
	}

	if val, ok := d.GetOk("conection_status"); ok {
		connection_status := val.(bool)
		params.Data.ConnectionStatus = &connection_status
	}

	if val, ok := d.GetOk("description"); ok {
		params.Data.Description = val.(string)
	}

	if val, ok := d.GetOk("enabled"); ok {
		params.Data.Enabled = val.(bool)
	}

	if val, ok := d.GetOk("label"); ok {
		params.Data.Description = val.(string)
	}

	if val, ok := d.GetOk("lag"); ok {
		lag := int64(val.(int))
		params.Data.Lag = &lag
	}

	if val, ok := d.GetOk("mac_address"); ok {
		mac_address := val.(string)
		params.Data.MacAddress = &mac_address
	}

	if val, ok := d.GetOk("mgmt_only"); ok {
		params.Data.MgmtOnly = val.(bool)
	}

	if val, ok := d.GetOk("mode"); ok {
		params.Data.Mode = val.(string)
	}

	if val, ok := d.GetOk("mtu"); ok {
		mtu := int64(val.(int))
		params.Data.Mtu = &mtu
	}

	if val, ok := d.GetOk("tags"); ok {
		params.Data.Tags = expandTags(val.([]interface{}))
	}

	if val, ok := d.GetOk("untagged_vlan"); ok {
		untagged_vlan := int64(val.(int))
		params.Data.UntaggedVlan = &untagged_vlan
	}

	r, err := c.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Could not create an interface: %v", err)
	}

	d.SetId(strconv.FormatInt(r.Payload.ID, 10))

	resourceDcimInterfaceRead(ctx, d, m)

	return diags
}

func resourceDcimInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	params := &dcim.DcimInterfacesReadParams{
		Context: ctx,
		ID:      id,
	}

	r, err := c.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Could not get an interface: %v", err)
	}

	if r.Payload.ConnectionStatus != nil {
		d.Set("connection_status", r.Payload.ConnectionStatus)
	}

	d.Set("description", r.Payload.Description)
	d.Set("enabled", r.Payload.Enabled)
	d.Set("label", r.Payload.Label)
	d.Set("lag", r.Payload.Lag)
	d.Set("mac_address", r.Payload.MacAddress)
	d.Set("mgmt_only", r.Payload.MgmtOnly)
	d.Set("mode", r.Payload.Mode)
	d.Set("mtu", r.Payload.Mtu)
	d.Set("tags", flattenTags(r.Payload.Tags))
	d.Set("untagged_vlan", r.Payload.UntaggedVlan)

	return diags
}

func resourceDcimInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse interface id: %v", err)
	}

	params := &dcim.DcimInterfacesPartialUpdateParams{
		Context: ctx,
		ID:      id,
	}

	name := d.Get("name").(string)
	device := int64(d.Get("device").(int))
	interface_type := d.Get("type").(string)

	params.Data = &models.WritableInterface{
		Name:        &name,
		Device:      &device,
		Type:        &interface_type,
		TaggedVlans: []int64{},
	}

	if d.HasChange("conection_status") {
		connection_status := d.Get("conection_status").(bool)
		params.Data.ConnectionStatus = &connection_status
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("enabled") {
		params.Data.Enabled = d.Get("enabled").(bool)
	}

	if d.HasChange("label") {
		params.Data.Label = d.Get("label").(string)
	}

	if d.HasChange("lag") {
		lag := int64(d.Get("lag").(int))
		params.Data.Lag = &lag
	}

	if d.HasChange("mac_address") {
		mac_address := d.Get("mac_address").(string)
		params.Data.MacAddress = &mac_address
	}

	if d.HasChange("mgmt_only") {
		params.Data.MgmtOnly = d.Get("mgmt_only").(bool)
	}

	if d.HasChange("mode") {
		params.Data.Label = d.Get("label").(string)
	}

	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		params.Data.Mtu = &mtu
	}

	if d.HasChange("tags") {
		params.Data.Tags = expandTags(d.Get("tags").([]interface{}))
	}

	if d.HasChange("untagged_vlan") {
		untagged_vlan := int64(d.Get("untagged_vlan").(int))
		params.Data.UntaggedVlan = &untagged_vlan
	}

	_, err = c.Dcim.DcimInterfacesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Could not update interface: %v", err)
	}

	return resourceDcimInterfaceRead(ctx, d, m)
}

func resourceDcimInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Could not parse interface ID: %v", err)
	}

	params := &dcim.DcimInterfacesDeleteParams{
		Context: ctx,
		ID:      id,
	}

	_, err = c.Dcim.DcimInterfacesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Could not to delete interface: %v", err)
	}

	d.SetId("")

	return diags
}
