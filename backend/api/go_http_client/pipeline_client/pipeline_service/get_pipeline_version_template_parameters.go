// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

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

// NewGetPipelineVersionTemplateParams creates a new GetPipelineVersionTemplateParams object
// with the default values initialized.
func NewGetPipelineVersionTemplateParams() *GetPipelineVersionTemplateParams {
	var ()
	return &GetPipelineVersionTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineVersionTemplateParamsWithTimeout creates a new GetPipelineVersionTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPipelineVersionTemplateParamsWithTimeout(timeout time.Duration) *GetPipelineVersionTemplateParams {
	var ()
	return &GetPipelineVersionTemplateParams{

		timeout: timeout,
	}
}

// NewGetPipelineVersionTemplateParamsWithContext creates a new GetPipelineVersionTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPipelineVersionTemplateParamsWithContext(ctx context.Context) *GetPipelineVersionTemplateParams {
	var ()
	return &GetPipelineVersionTemplateParams{

		Context: ctx,
	}
}

// NewGetPipelineVersionTemplateParamsWithHTTPClient creates a new GetPipelineVersionTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPipelineVersionTemplateParamsWithHTTPClient(client *http.Client) *GetPipelineVersionTemplateParams {
	var ()
	return &GetPipelineVersionTemplateParams{
		HTTPClient: client,
	}
}

/*GetPipelineVersionTemplateParams contains all the parameters to send to the API endpoint
for the get pipeline version template operation typically these are written to a http.Request
*/
type GetPipelineVersionTemplateParams struct {

	/*VersionID
	  The ID of the pipeline version whose template is to be retrieved.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) WithTimeout(timeout time.Duration) *GetPipelineVersionTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) WithContext(ctx context.Context) *GetPipelineVersionTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) WithHTTPClient(client *http.Client) *GetPipelineVersionTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersionID adds the versionID to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) WithVersionID(versionID string) *GetPipelineVersionTemplateParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get pipeline version template params
func (o *GetPipelineVersionTemplateParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineVersionTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param version_id
	if err := r.SetPathParam("version_id", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
