// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchAclsRidReader is a Reader for the PatchAclsRid structure.
type PatchAclsRidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAclsRidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchAclsRidNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAclsRidNoContent creates a PatchAclsRidNoContent with default headers values
func NewPatchAclsRidNoContent() *PatchAclsRidNoContent {
	return &PatchAclsRidNoContent{}
}

/*PatchAclsRidNoContent handles this case with default header values.

Success.
*/
type PatchAclsRidNoContent struct {
}

func (o *PatchAclsRidNoContent) Error() string {
	return fmt.Sprintf("[PATCH /acls/{rid}][%d] patchAclsRidNoContent ", 204)
}

func (o *PatchAclsRidNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
