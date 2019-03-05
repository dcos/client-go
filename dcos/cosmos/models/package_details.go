// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PackageDetails package details
// swagger:model packageDetails
type PackageDetails struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// True if this package installs a new Mesos framework.
	Framework *bool `json:"framework,omitempty"`

	// licenses
	Licenses Licences `json:"licenses,omitempty"`

	// maintainer
	// Required: true
	Maintainer *string `json:"maintainer"`

	// name
	// Required: true
	Name *string `json:"name"`

	// packaging version
	// Required: true
	PackagingVersion *string `json:"packagingVersion"`

	// Post installation notes that would be useful to the user of this package.
	PostInstallNotes string `json:"postInstallNotes,omitempty"`

	// Post uninstallation notes that would be useful to the user of this package.
	PostUninstallNotes string `json:"postUninstallNotes,omitempty"`

	// Pre installation notes that would be useful to the user of this package.
	PreInstallNotes string `json:"preInstallNotes,omitempty"`

	// Corresponds to the revision index from the universe directory structure
	// Minimum: 0
	ReleaseVersion *int64 `json:"releaseVersion,omitempty"`

	// scm
	Scm string `json:"scm,omitempty"`

	// Flag indicating if the package is selected in search results
	Selected *bool `json:"selected,omitempty"`

	// tags
	// Required: true
	Tags []string `json:"tags"`

	// version
	// Required: true
	Version *string `json:"version"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this package details
func (m *PackageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackagingVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetails) validateLicenses(formats strfmt.Registry) error {

	if swag.IsZero(m.Licenses) { // not required
		return nil
	}

	if err := m.Licenses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("licenses")
		}
		return err
	}

	return nil
}

func (m *PackageDetails) validateMaintainer(formats strfmt.Registry) error {

	if err := validate.Required("maintainer", "body", m.Maintainer); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetails) validatePackagingVersion(formats strfmt.Registry) error {

	if err := validate.Required("packagingVersion", "body", m.PackagingVersion); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetails) validateReleaseVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseVersion) { // not required
		return nil
	}

	if err := validate.MinimumInt("releaseVersion", "body", int64(*m.ReleaseVersion), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetails) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.Pattern("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), `^[^\s]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *PackageDetails) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDetails) UnmarshalBinary(b []byte) error {
	var res PackageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}