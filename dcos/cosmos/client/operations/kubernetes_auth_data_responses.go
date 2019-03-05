// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/cosmos/models"
)

// KubernetesAuthDataReader is a Reader for the KubernetesAuthData structure.
type KubernetesAuthDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesAuthDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewKubernetesAuthDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewKubernetesAuthDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKubernetesAuthDataOK creates a KubernetesAuthDataOK with default headers values
func NewKubernetesAuthDataOK() *KubernetesAuthDataOK {
	return &KubernetesAuthDataOK{}
}

/*KubernetesAuthDataOK handles this case with default header values.

KubernetesAuthDataOK kubernetes auth data o k
*/
type KubernetesAuthDataOK struct {
	Payload *models.KubernetesAuthDataResponse
}

func (o *KubernetesAuthDataOK) Error() string {
	return fmt.Sprintf("[GET /service/{appId}/v1/auth/data][%d] kubernetesAuthDataOK  %+v", 200, o.Payload)
}

func (o *KubernetesAuthDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesAuthDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesAuthDataNotFound creates a KubernetesAuthDataNotFound with default headers values
func NewKubernetesAuthDataNotFound() *KubernetesAuthDataNotFound {
	return &KubernetesAuthDataNotFound{}
}

/*KubernetesAuthDataNotFound handles this case with default header values.

Service not found.
*/
type KubernetesAuthDataNotFound struct {
	Payload *models.Error
}

func (o *KubernetesAuthDataNotFound) Error() string {
	return fmt.Sprintf("[GET /service/{appId}/v1/auth/data][%d] kubernetesAuthDataNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesAuthDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
