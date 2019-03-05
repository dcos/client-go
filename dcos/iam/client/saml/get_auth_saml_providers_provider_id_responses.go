// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetAuthSamlProvidersProviderIDReader is a Reader for the GetAuthSamlProvidersProviderID structure.
type GetAuthSamlProvidersProviderIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthSamlProvidersProviderIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthSamlProvidersProviderIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthSamlProvidersProviderIDOK creates a GetAuthSamlProvidersProviderIDOK with default headers values
func NewGetAuthSamlProvidersProviderIDOK() *GetAuthSamlProvidersProviderIDOK {
	return &GetAuthSamlProvidersProviderIDOK{}
}

/*GetAuthSamlProvidersProviderIDOK handles this case with default header values.

Success.
*/
type GetAuthSamlProvidersProviderIDOK struct {
	Payload *models.SAMLProviderConfig
}

func (o *GetAuthSamlProvidersProviderIDOK) Error() string {
	return fmt.Sprintf("[GET /auth/saml/providers/{provider-id}][%d] getAuthSamlProvidersProviderIdOK  %+v", 200, o.Payload)
}

func (o *GetAuthSamlProvidersProviderIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAMLProviderConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
