// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetGroupsGidPermissionsReader is a Reader for the GetGroupsGidPermissions structure.
type GetGroupsGidPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupsGidPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGroupsGidPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGroupsGidPermissionsOK creates a GetGroupsGidPermissionsOK with default headers values
func NewGetGroupsGidPermissionsOK() *GetGroupsGidPermissionsOK {
	return &GetGroupsGidPermissionsOK{}
}

/*GetGroupsGidPermissionsOK handles this case with default header values.

Success.
*/
type GetGroupsGidPermissionsOK struct {
	Payload *models.GroupPermissions
}

func (o *GetGroupsGidPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /groups/{gid}/permissions][%d] getGroupsGidPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetGroupsGidPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPermissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
