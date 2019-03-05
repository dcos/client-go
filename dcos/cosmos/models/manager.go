// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Manager manager
// swagger:model manager
type Manager struct {

	// min package version
	MinPackageVersion string `json:"minPackageVersion,omitempty"`

	// package name
	// Required: true
	PackageName *string `json:"packageName"`
}

// Validate validates this manager
func (m *Manager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager) validatePackageName(formats strfmt.Registry) error {

	if err := validate.Required("packageName", "body", m.PackageName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Manager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Manager) UnmarshalBinary(b []byte) error {
	var res Manager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
