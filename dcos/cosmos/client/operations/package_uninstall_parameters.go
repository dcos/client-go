// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/dcos/client-go/dcos/cosmos/models"
)

// NewPackageUninstallParams creates a new PackageUninstallParams object
// with the default values initialized.
func NewPackageUninstallParams() *PackageUninstallParams {
	var ()
	return &PackageUninstallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackageUninstallParamsWithTimeout creates a new PackageUninstallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackageUninstallParamsWithTimeout(timeout time.Duration) *PackageUninstallParams {
	var ()
	return &PackageUninstallParams{

		timeout: timeout,
	}
}

// NewPackageUninstallParamsWithContext creates a new PackageUninstallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackageUninstallParamsWithContext(ctx context.Context) *PackageUninstallParams {
	var ()
	return &PackageUninstallParams{

		Context: ctx,
	}
}

// NewPackageUninstallParamsWithHTTPClient creates a new PackageUninstallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackageUninstallParamsWithHTTPClient(client *http.Client) *PackageUninstallParams {
	var ()
	return &PackageUninstallParams{
		HTTPClient: client,
	}
}

/*PackageUninstallParams contains all the parameters to send to the API endpoint
for the package uninstall operation typically these are written to a http.Request
*/
type PackageUninstallParams struct {

	/*Accept*/
	Accept *string
	/*Body*/
	Body *models.UninstallRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the package uninstall params
func (o *PackageUninstallParams) WithTimeout(timeout time.Duration) *PackageUninstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package uninstall params
func (o *PackageUninstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package uninstall params
func (o *PackageUninstallParams) WithContext(ctx context.Context) *PackageUninstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package uninstall params
func (o *PackageUninstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package uninstall params
func (o *PackageUninstallParams) WithHTTPClient(client *http.Client) *PackageUninstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package uninstall params
func (o *PackageUninstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the package uninstall params
func (o *PackageUninstallParams) WithAccept(accept *string) *PackageUninstallParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the package uninstall params
func (o *PackageUninstallParams) SetAccept(accept *string) {
	o.Accept = accept
}

// WithBody adds the body to the package uninstall params
func (o *PackageUninstallParams) WithBody(body *models.UninstallRequest) *PackageUninstallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the package uninstall params
func (o *PackageUninstallParams) SetBody(body *models.UninstallRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PackageUninstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accept != nil {

		// header param Accept
		if err := r.SetHeaderParam("Accept", *o.Accept); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
