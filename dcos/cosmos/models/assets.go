// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Assets assets
// swagger:model assets
type Assets struct {

	// container
	Container *AssetsContainer `json:"container,omitempty"`

	// uris
	Uris Uris `json:"uris,omitempty"`
}

// Validate validates this assets
func (m *Assets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUris(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Assets) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.Container) { // not required
		return nil
	}

	if m.Container != nil {
		if err := m.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container")
			}
			return err
		}
	}

	return nil
}

func (m *Assets) validateUris(formats strfmt.Registry) error {

	if swag.IsZero(m.Uris) { // not required
		return nil
	}

	if err := m.Uris.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("uris")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Assets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Assets) UnmarshalBinary(b []byte) error {
	var res Assets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssetsContainer assets container
// swagger:model AssetsContainer
type AssetsContainer struct {

	// docker
	Docker Docker `json:"docker,omitempty"`
}

// Validate validates this assets container
func (m *AssetsContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetsContainer) validateDocker(formats strfmt.Registry) error {

	if swag.IsZero(m.Docker) { // not required
		return nil
	}

	if err := m.Docker.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("container" + "." + "docker")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetsContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetsContainer) UnmarshalBinary(b []byte) error {
	var res AssetsContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
