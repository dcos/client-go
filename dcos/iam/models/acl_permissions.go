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

// ACLPermissions ACL permissions
// swagger:model ACLPermissions
type ACLPermissions struct {

	// groups
	Groups []*ACLPermissionsGroupsItems0 `json:"groups"`

	// users
	Users []*ACLPermissionsUsersItems0 `json:"users"`
}

// Validate validates this ACL permissions
func (m *ACLPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACLPermissions) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ACLPermissions) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACLPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACLPermissions) UnmarshalBinary(b []byte) error {
	var res ACLPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ACLPermissionsGroupsItems0 ACL permissions groups items0
// swagger:model ACLPermissionsGroupsItems0
type ACLPermissionsGroupsItems0 struct {

	// actions
	// Required: true
	Actions []*Action `json:"actions"`

	// gid
	// Required: true
	Gid *string `json:"gid"`

	// groupurl
	// Required: true
	Groupurl *string `json:"groupurl"`
}

// Validate validates this ACL permissions groups items0
func (m *ACLPermissionsGroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupurl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACLPermissionsGroupsItems0) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ACLPermissionsGroupsItems0) validateGid(formats strfmt.Registry) error {

	if err := validate.Required("gid", "body", m.Gid); err != nil {
		return err
	}

	return nil
}

func (m *ACLPermissionsGroupsItems0) validateGroupurl(formats strfmt.Registry) error {

	if err := validate.Required("groupurl", "body", m.Groupurl); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACLPermissionsGroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACLPermissionsGroupsItems0) UnmarshalBinary(b []byte) error {
	var res ACLPermissionsGroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ACLPermissionsUsersItems0 ACL permissions users items0
// swagger:model ACLPermissionsUsersItems0
type ACLPermissionsUsersItems0 struct {

	// actions
	// Required: true
	Actions []*Action `json:"actions"`

	// uid
	// Required: true
	UID *string `json:"uid"`

	// userurl
	// Required: true
	Userurl *string `json:"userurl"`
}

// Validate validates this ACL permissions users items0
func (m *ACLPermissionsUsersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserurl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACLPermissionsUsersItems0) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ACLPermissionsUsersItems0) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

func (m *ACLPermissionsUsersItems0) validateUserurl(formats strfmt.Registry) error {

	if err := validate.Required("userurl", "body", m.Userurl); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACLPermissionsUsersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACLPermissionsUsersItems0) UnmarshalBinary(b []byte) error {
	var res ACLPermissionsUsersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
