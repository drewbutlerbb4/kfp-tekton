// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIPipelineSpec api pipeline spec
// swagger:model apiPipelineSpec
type APIPipelineSpec struct {

	// The parameter user provide to inject to the pipeline JSON.
	// If a default value of a parameter exist in the JSON,
	// the value user provided here will replace.
	Parameters []*APIParameter `json:"parameters"`

	// Optional input field. The ID of the pipeline user uploaded before.
	PipelineID string `json:"pipeline_id,omitempty"`

	// Optional input field. The raw pipeline JSON spec.
	PipelineManifest string `json:"pipeline_manifest,omitempty"`

	// Optional output field. The name of the pipeline.
	// Not empty if the pipeline id is not empty.
	PipelineName string `json:"pipeline_name,omitempty"`

	// Optional input field. The marshalled raw argo JSON workflow.
	// This will be deprecated when pipeline_manifest is in use.
	WorkflowManifest string `json:"workflow_manifest,omitempty"`
}

// Validate validates this api pipeline spec
func (m *APIPipelineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPipelineSpec) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPipelineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPipelineSpec) UnmarshalBinary(b []byte) error {
	var res APIPipelineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
