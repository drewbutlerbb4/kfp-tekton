// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobMode Required input.
//
//  - DISABLED: The job won't schedule any run if disabled.
// swagger:model JobMode
type JobMode string

const (

	// JobModeUNKNOWNMODE captures enum value "UNKNOWN_MODE"
	JobModeUNKNOWNMODE JobMode = "UNKNOWN_MODE"

	// JobModeENABLED captures enum value "ENABLED"
	JobModeENABLED JobMode = "ENABLED"

	// JobModeDISABLED captures enum value "DISABLED"
	JobModeDISABLED JobMode = "DISABLED"
)

// for schema
var jobModeEnum []interface{}

func init() {
	var res []JobMode
	if err := json.Unmarshal([]byte(`["UNKNOWN_MODE","ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobModeEnum = append(jobModeEnum, v)
	}
}

func (m JobMode) validateJobModeEnum(path, location string, value JobMode) error {
	if err := validate.Enum(path, location, value, jobModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this job mode
func (m JobMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
