// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/kubeflow/pipelines/backend/api/go_http_client/job_model"
)

// DisableJobReader is a Reader for the DisableJob structure.
type DisableJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDisableJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDisableJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableJobOK creates a DisableJobOK with default headers values
func NewDisableJobOK() *DisableJobOK {
	return &DisableJobOK{}
}

/*DisableJobOK handles this case with default header values.

A successful response.
*/
type DisableJobOK struct {
	Payload interface{}
}

func (o *DisableJobOK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/jobs/{id}/disable][%d] disableJobOK  %+v", 200, o.Payload)
}

func (o *DisableJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableJobDefault creates a DisableJobDefault with default headers values
func NewDisableJobDefault(code int) *DisableJobDefault {
	return &DisableJobDefault{
		_statusCode: code,
	}
}

/*DisableJobDefault handles this case with default header values.

DisableJobDefault disable job default
*/
type DisableJobDefault struct {
	_statusCode int

	Payload *job_model.APIStatus
}

// Code gets the status code for the disable job default response
func (o *DisableJobDefault) Code() int {
	return o._statusCode
}

func (o *DisableJobDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/jobs/{id}/disable][%d] DisableJob default  %+v", o._statusCode, o.Payload)
}

func (o *DisableJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
