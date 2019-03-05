// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

// NewGetSecretStorePathToSecretParams creates a new GetSecretStorePathToSecretParams object
// with the default values initialized.
func NewGetSecretStorePathToSecretParams() *GetSecretStorePathToSecretParams {
	var ()
	return &GetSecretStorePathToSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecretStorePathToSecretParamsWithTimeout creates a new GetSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecretStorePathToSecretParamsWithTimeout(timeout time.Duration) *GetSecretStorePathToSecretParams {
	var ()
	return &GetSecretStorePathToSecretParams{

		timeout: timeout,
	}
}

// NewGetSecretStorePathToSecretParamsWithContext creates a new GetSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecretStorePathToSecretParamsWithContext(ctx context.Context) *GetSecretStorePathToSecretParams {
	var ()
	return &GetSecretStorePathToSecretParams{

		Context: ctx,
	}
}

// NewGetSecretStorePathToSecretParamsWithHTTPClient creates a new GetSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecretStorePathToSecretParamsWithHTTPClient(client *http.Client) *GetSecretStorePathToSecretParams {
	var ()
	return &GetSecretStorePathToSecretParams{
		HTTPClient: client,
	}
}

/*GetSecretStorePathToSecretParams contains all the parameters to send to the API endpoint
for the get secret store path to secret operation typically these are written to a http.Request
*/
type GetSecretStorePathToSecretParams struct {

	/*List
	  If set to true, list only secret keys.


	*/
	List *string
	/*PathToSecret
	  The path to store the secret in.

	*/
	PathToSecret string
	/*Store
	  The backend store from which to get the secret.

	*/
	Store string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithTimeout(timeout time.Duration) *GetSecretStorePathToSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithContext(ctx context.Context) *GetSecretStorePathToSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithHTTPClient(client *http.Client) *GetSecretStorePathToSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithList adds the list to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithList(list *string) *GetSecretStorePathToSecretParams {
	o.SetList(list)
	return o
}

// SetList adds the list to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetList(list *string) {
	o.List = list
}

// WithPathToSecret adds the pathToSecret to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithPathToSecret(pathToSecret string) *GetSecretStorePathToSecretParams {
	o.SetPathToSecret(pathToSecret)
	return o
}

// SetPathToSecret adds the pathToSecret to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetPathToSecret(pathToSecret string) {
	o.PathToSecret = pathToSecret
}

// WithStore adds the store to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) WithStore(store string) *GetSecretStorePathToSecretParams {
	o.SetStore(store)
	return o
}

// SetStore adds the store to the get secret store path to secret params
func (o *GetSecretStorePathToSecretParams) SetStore(store string) {
	o.Store = store
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecretStorePathToSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.List != nil {

		// query param list
		var qrList string
		if o.List != nil {
			qrList = *o.List
		}
		qList := qrList
		if qList != "" {
			if err := r.SetQueryParam("list", qList); err != nil {
				return err
			}
		}

	}

	// path param path-to-secret
	if err := r.SetPathParam("path-to-secret", o.PathToSecret); err != nil {
		return err
	}

	// path param store
	if err := r.SetPathParam("store", o.Store); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}