// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAuthSamlProvidersProviderIDAcsCallbackURLParams creates a new GetAuthSamlProvidersProviderIDAcsCallbackURLParams object
// with the default values initialized.
func NewGetAuthSamlProvidersProviderIDAcsCallbackURLParams() *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	var ()
	return &GetAuthSamlProvidersProviderIDAcsCallbackURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithTimeout creates a new GetAuthSamlProvidersProviderIDAcsCallbackURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithTimeout(timeout time.Duration) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	var ()
	return &GetAuthSamlProvidersProviderIDAcsCallbackURLParams{

		timeout: timeout,
	}
}

// NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithContext creates a new GetAuthSamlProvidersProviderIDAcsCallbackURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithContext(ctx context.Context) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	var ()
	return &GetAuthSamlProvidersProviderIDAcsCallbackURLParams{

		Context: ctx,
	}
}

// NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithHTTPClient creates a new GetAuthSamlProvidersProviderIDAcsCallbackURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthSamlProvidersProviderIDAcsCallbackURLParamsWithHTTPClient(client *http.Client) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	var ()
	return &GetAuthSamlProvidersProviderIDAcsCallbackURLParams{
		HTTPClient: client,
	}
}

/*GetAuthSamlProvidersProviderIDAcsCallbackURLParams contains all the parameters to send to the API endpoint
for the get auth saml providers provider ID acs callback URL operation typically these are written to a http.Request
*/
type GetAuthSamlProvidersProviderIDAcsCallbackURLParams struct {

	/*ProviderID
	  The ID of the SAML provider to retrieve the ACS callback URL for.

	*/
	ProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) WithTimeout(timeout time.Duration) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) WithContext(ctx context.Context) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) WithHTTPClient(client *http.Client) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderID adds the providerID to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) WithProviderID(providerID string) *GetAuthSamlProvidersProviderIDAcsCallbackURLParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the get auth saml providers provider ID acs callback URL params
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthSamlProvidersProviderIDAcsCallbackURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider-id
	if err := r.SetPathParam("provider-id", o.ProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
