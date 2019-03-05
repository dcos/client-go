// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchSecretStorePathToSecretReader is a Reader for the PatchSecretStorePathToSecret structure.
type PatchSecretStorePathToSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSecretStorePathToSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchSecretStorePathToSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPatchSecretStorePathToSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSecretStorePathToSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSecretStorePathToSecretNoContent creates a PatchSecretStorePathToSecretNoContent with default headers values
func NewPatchSecretStorePathToSecretNoContent() *PatchSecretStorePathToSecretNoContent {
	return &PatchSecretStorePathToSecretNoContent{}
}

/*PatchSecretStorePathToSecretNoContent handles this case with default header values.

Secret updated.
*/
type PatchSecretStorePathToSecretNoContent struct {
}

func (o *PatchSecretStorePathToSecretNoContent) Error() string {
	return fmt.Sprintf("[PATCH /secret/{store}/{path-to-secret}][%d] patchSecretStorePathToSecretNoContent ", 204)
}

func (o *PatchSecretStorePathToSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSecretStorePathToSecretForbidden creates a PatchSecretStorePathToSecretForbidden with default headers values
func NewPatchSecretStorePathToSecretForbidden() *PatchSecretStorePathToSecretForbidden {
	return &PatchSecretStorePathToSecretForbidden{}
}

/*PatchSecretStorePathToSecretForbidden handles this case with default header values.

Forbidden.
*/
type PatchSecretStorePathToSecretForbidden struct {
}

func (o *PatchSecretStorePathToSecretForbidden) Error() string {
	return fmt.Sprintf("[PATCH /secret/{store}/{path-to-secret}][%d] patchSecretStorePathToSecretForbidden ", 403)
}

func (o *PatchSecretStorePathToSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSecretStorePathToSecretNotFound creates a PatchSecretStorePathToSecretNotFound with default headers values
func NewPatchSecretStorePathToSecretNotFound() *PatchSecretStorePathToSecretNotFound {
	return &PatchSecretStorePathToSecretNotFound{}
}

/*PatchSecretStorePathToSecretNotFound handles this case with default header values.

Secret not found.
*/
type PatchSecretStorePathToSecretNotFound struct {
}

func (o *PatchSecretStorePathToSecretNotFound) Error() string {
	return fmt.Sprintf("[PATCH /secret/{store}/{path-to-secret}][%d] patchSecretStorePathToSecretNotFound ", 404)
}

func (o *PatchSecretStorePathToSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
