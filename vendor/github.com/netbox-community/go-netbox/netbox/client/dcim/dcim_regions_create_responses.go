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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimRegionsCreateReader is a Reader for the DcimRegionsCreate structure.
type DcimRegionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRegionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsCreateCreated creates a DcimRegionsCreateCreated with default headers values
func NewDcimRegionsCreateCreated() *DcimRegionsCreateCreated {
	return &DcimRegionsCreateCreated{}
}

/*DcimRegionsCreateCreated handles this case with default header values.

DcimRegionsCreateCreated dcim regions create created
*/
type DcimRegionsCreateCreated struct {
	Payload *models.Region
}

func (o *DcimRegionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/regions/][%d] dcimRegionsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRegionsCreateCreated) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsCreateDefault creates a DcimRegionsCreateDefault with default headers values
func NewDcimRegionsCreateDefault(code int) *DcimRegionsCreateDefault {
	return &DcimRegionsCreateDefault{
		_statusCode: code,
	}
}

/*DcimRegionsCreateDefault handles this case with default header values.

DcimRegionsCreateDefault dcim regions create default
*/
type DcimRegionsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions create default response
func (o *DcimRegionsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRegionsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/regions/][%d] dcim_regions_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
