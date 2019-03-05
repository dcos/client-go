// Code generated by go-swagger; DO NOT EDIT.

package permissions

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

// NewDeleteAclsRidGroupsGidActionParams creates a new DeleteAclsRidGroupsGidActionParams object
// with the default values initialized.
func NewDeleteAclsRidGroupsGidActionParams() *DeleteAclsRidGroupsGidActionParams {
	var ()
	return &DeleteAclsRidGroupsGidActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAclsRidGroupsGidActionParamsWithTimeout creates a new DeleteAclsRidGroupsGidActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAclsRidGroupsGidActionParamsWithTimeout(timeout time.Duration) *DeleteAclsRidGroupsGidActionParams {
	var ()
	return &DeleteAclsRidGroupsGidActionParams{

		timeout: timeout,
	}
}

// NewDeleteAclsRidGroupsGidActionParamsWithContext creates a new DeleteAclsRidGroupsGidActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAclsRidGroupsGidActionParamsWithContext(ctx context.Context) *DeleteAclsRidGroupsGidActionParams {
	var ()
	return &DeleteAclsRidGroupsGidActionParams{

		Context: ctx,
	}
}

// NewDeleteAclsRidGroupsGidActionParamsWithHTTPClient creates a new DeleteAclsRidGroupsGidActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAclsRidGroupsGidActionParamsWithHTTPClient(client *http.Client) *DeleteAclsRidGroupsGidActionParams {
	var ()
	return &DeleteAclsRidGroupsGidActionParams{
		HTTPClient: client,
	}
}

/*DeleteAclsRidGroupsGidActionParams contains all the parameters to send to the API endpoint
for the delete acls rid groups gid action operation typically these are written to a http.Request
*/
type DeleteAclsRidGroupsGidActionParams struct {

	/*Action
	  action name

	*/
	Action string
	/*Gid
	  group ID.

	*/
	Gid string
	/*Rid
	  resource ID.

	*/
	Rid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithTimeout(timeout time.Duration) *DeleteAclsRidGroupsGidActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithContext(ctx context.Context) *DeleteAclsRidGroupsGidActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithHTTPClient(client *http.Client) *DeleteAclsRidGroupsGidActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithAction(action string) *DeleteAclsRidGroupsGidActionParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetAction(action string) {
	o.Action = action
}

// WithGid adds the gid to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithGid(gid string) *DeleteAclsRidGroupsGidActionParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetGid(gid string) {
	o.Gid = gid
}

// WithRid adds the rid to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) WithRid(rid string) *DeleteAclsRidGroupsGidActionParams {
	o.SetRid(rid)
	return o
}

// SetRid adds the rid to the delete acls rid groups gid action params
func (o *DeleteAclsRidGroupsGidActionParams) SetRid(rid string) {
	o.Rid = rid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAclsRidGroupsGidActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", o.Action); err != nil {
		return err
	}

	// path param gid
	if err := r.SetPathParam("gid", o.Gid); err != nil {
		return err
	}

	// path param rid
	if err := r.SetPathParam("rid", o.Rid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
