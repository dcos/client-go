// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAuthSamlProvidersProviderIDSpMetadataReader is a Reader for the GetAuthSamlProvidersProviderIDSpMetadata structure.
type GetAuthSamlProvidersProviderIDSpMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthSamlProvidersProviderIDSpMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthSamlProvidersProviderIDSpMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthSamlProvidersProviderIDSpMetadataOK creates a GetAuthSamlProvidersProviderIDSpMetadataOK with default headers values
func NewGetAuthSamlProvidersProviderIDSpMetadataOK() *GetAuthSamlProvidersProviderIDSpMetadataOK {
	return &GetAuthSamlProvidersProviderIDSpMetadataOK{}
}

/*GetAuthSamlProvidersProviderIDSpMetadataOK handles this case with default header values.

The response body contains the metadata in UTF-8 encoding, setting the Content-Type to `application/samlmetadata+xml`.
*/
type GetAuthSamlProvidersProviderIDSpMetadataOK struct {
}

func (o *GetAuthSamlProvidersProviderIDSpMetadataOK) Error() string {
	return fmt.Sprintf("[GET /auth/saml/providers/{provider-id}/sp-metadata][%d] getAuthSamlProvidersProviderIdSpMetadataOK ", 200)
}

func (o *GetAuthSamlProvidersProviderIDSpMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}