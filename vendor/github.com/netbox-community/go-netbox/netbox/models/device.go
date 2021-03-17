// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Device device
//
// swagger:model Device
type Device struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// device role
	// Required: true
	DeviceRole *NestedDeviceRole `json:"device_role"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// face
	Face *DeviceFace `json:"face,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// Name
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// status
	Status *DeviceStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// virtual chassis
	VirtualChassis *NestedVirtualChassis `json:"virtual_chassis,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(*m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	if m.DeviceRole != nil {
		if err := m.DeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateFace(formats strfmt.Registry) error {

	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateParentDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(*m.Position), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", int64(*m.Position), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Device) validatePrimaryIP(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validatePrimaryIp4(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validatePrimaryIp6(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateRack(formats strfmt.Registry) error {

	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateVcPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", int64(*m.VcPosition), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", int64(*m.VcPosition), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateVcPriority(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", int64(*m.VcPriority), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", int64(*m.VcPriority), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateVirtualChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualChassis) { // not required
		return nil
	}

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceFace Face
//
// swagger:model DeviceFace
type DeviceFace struct {

	// label
	// Required: true
	// Enum: [Front Rear]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [front rear]
	Value *string `json:"value"`
}

// Validate validates this device face
func (m *DeviceFace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceFaceTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Front","Rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceFaceTypeLabelPropEnum = append(deviceFaceTypeLabelPropEnum, v)
	}
}

const (

	// DeviceFaceLabelFront captures enum value "Front"
	DeviceFaceLabelFront string = "Front"

	// DeviceFaceLabelRear captures enum value "Rear"
	DeviceFaceLabelRear string = "Rear"
)

// prop value enum
func (m *DeviceFace) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceFaceTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("face"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceFaceTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceFaceTypeValuePropEnum = append(deviceFaceTypeValuePropEnum, v)
	}
}

const (

	// DeviceFaceValueFront captures enum value "front"
	DeviceFaceValueFront string = "front"

	// DeviceFaceValueRear captures enum value "rear"
	DeviceFaceValueRear string = "rear"
)

// prop value enum
func (m *DeviceFace) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceFaceTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("face"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceFace) UnmarshalBinary(b []byte) error {
	var res DeviceFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceStatus Status
//
// swagger:model DeviceStatus
type DeviceStatus struct {

	// label
	// Required: true
	// Enum: [Offline Active Planned Staged Failed Inventory Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [offline active planned staged failed inventory decommissioning]
	Value *string `json:"value"`
}

// Validate validates this device status
func (m *DeviceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Active","Planned","Staged","Failed","Inventory","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceStatusTypeLabelPropEnum = append(deviceStatusTypeLabelPropEnum, v)
	}
}

const (

	// DeviceStatusLabelOffline captures enum value "Offline"
	DeviceStatusLabelOffline string = "Offline"

	// DeviceStatusLabelActive captures enum value "Active"
	DeviceStatusLabelActive string = "Active"

	// DeviceStatusLabelPlanned captures enum value "Planned"
	DeviceStatusLabelPlanned string = "Planned"

	// DeviceStatusLabelStaged captures enum value "Staged"
	DeviceStatusLabelStaged string = "Staged"

	// DeviceStatusLabelFailed captures enum value "Failed"
	DeviceStatusLabelFailed string = "Failed"

	// DeviceStatusLabelInventory captures enum value "Inventory"
	DeviceStatusLabelInventory string = "Inventory"

	// DeviceStatusLabelDecommissioning captures enum value "Decommissioning"
	DeviceStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *DeviceStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","inventory","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceStatusTypeValuePropEnum = append(deviceStatusTypeValuePropEnum, v)
	}
}

const (

	// DeviceStatusValueOffline captures enum value "offline"
	DeviceStatusValueOffline string = "offline"

	// DeviceStatusValueActive captures enum value "active"
	DeviceStatusValueActive string = "active"

	// DeviceStatusValuePlanned captures enum value "planned"
	DeviceStatusValuePlanned string = "planned"

	// DeviceStatusValueStaged captures enum value "staged"
	DeviceStatusValueStaged string = "staged"

	// DeviceStatusValueFailed captures enum value "failed"
	DeviceStatusValueFailed string = "failed"

	// DeviceStatusValueInventory captures enum value "inventory"
	DeviceStatusValueInventory string = "inventory"

	// DeviceStatusValueDecommissioning captures enum value "decommissioning"
	DeviceStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *DeviceStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatus) UnmarshalBinary(b []byte) error {
	var res DeviceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
