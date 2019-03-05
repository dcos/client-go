// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetAclsRidGroupsGidReader is a Reader for the GetAclsRidGroupsGid structure.
type GetAclsRidGroupsGidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAclsRidGroupsGidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAclsRidGroupsGidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAclsRidGroupsGidOK creates a GetAclsRidGroupsGidOK with default headers values
func NewGetAclsRidGroupsGidOK() *GetAclsRidGroupsGidOK {
	return &GetAclsRidGroupsGidOK{}
}

/*GetAclsRidGroupsGidOK handles this case with default header values.

Success.
*/
type GetAclsRidGroupsGidOK struct {
	Payload *GetAclsRidGroupsGidOKBody
}

func (o *GetAclsRidGroupsGidOK) Error() string {
	return fmt.Sprintf("[GET /acls/{rid}/groups/{gid}][%d] getAclsRidGroupsGidOK  %+v", 200, o.Payload)
}

func (o *GetAclsRidGroupsGidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAclsRidGroupsGidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAclsRidGroupsGidOKBody get acls rid groups gid o k body
swagger:model GetAclsRidGroupsGidOKBody
*/
type GetAclsRidGroupsGidOKBody struct {

	// array
	Array []*models.Action `json:"array"`
}

// Validate validates this get acls rid groups gid o k body
func (o *GetAclsRidGroupsGidOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAclsRidGroupsGidOKBody) validateArray(formats strfmt.Registry) error {

	if swag.IsZero(o.Array) { // not required
		return nil
	}

	for i := 0; i < len(o.Array); i++ {
		if swag.IsZero(o.Array[i]) { // not required
			continue
		}

		if o.Array[i] != nil {
			if err := o.Array[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAclsRidGroupsGidOK" + "." + "array" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAclsRidGroupsGidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAclsRidGroupsGidOKBody) UnmarshalBinary(b []byte) error {
	var res GetAclsRidGroupsGidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
