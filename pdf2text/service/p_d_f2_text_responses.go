// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PDF2TextReader is a Reader for the PDF2Text structure.
type PDF2TextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PDF2TextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPDF2TextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPDF2TextInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPDF2TextOK creates a PDF2TextOK with default headers values
func NewPDF2TextOK() *PDF2TextOK {
	return &PDF2TextOK{}
}

/*
PDF2TextOK describes a response with status code 200, with default header values.

pdf2text result
*/
type PDF2TextOK struct {
	Payload string
}

// IsSuccess returns true when this p d f2 text o k response has a 2xx status code
func (o *PDF2TextOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this p d f2 text o k response has a 3xx status code
func (o *PDF2TextOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p d f2 text o k response has a 4xx status code
func (o *PDF2TextOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this p d f2 text o k response has a 5xx status code
func (o *PDF2TextOK) IsServerError() bool {
	return false
}

// IsCode returns true when this p d f2 text o k response a status code equal to that given
func (o *PDF2TextOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the p d f2 text o k response
func (o *PDF2TextOK) Code() int {
	return 200
}

func (o *PDF2TextOK) Error() string {
	return fmt.Sprintf("[POST /][%d] pDF2TextOK  %+v", 200, o.Payload)
}

func (o *PDF2TextOK) String() string {
	return fmt.Sprintf("[POST /][%d] pDF2TextOK  %+v", 200, o.Payload)
}

func (o *PDF2TextOK) GetPayload() string {
	return o.Payload
}

func (o *PDF2TextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPDF2TextInternalServerError creates a PDF2TextInternalServerError with default headers values
func NewPDF2TextInternalServerError() *PDF2TextInternalServerError {
	return &PDF2TextInternalServerError{}
}

/*
PDF2TextInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type PDF2TextInternalServerError struct {
}

// IsSuccess returns true when this p d f2 text internal server error response has a 2xx status code
func (o *PDF2TextInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p d f2 text internal server error response has a 3xx status code
func (o *PDF2TextInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p d f2 text internal server error response has a 4xx status code
func (o *PDF2TextInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this p d f2 text internal server error response has a 5xx status code
func (o *PDF2TextInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this p d f2 text internal server error response a status code equal to that given
func (o *PDF2TextInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the p d f2 text internal server error response
func (o *PDF2TextInternalServerError) Code() int {
	return 500
}

func (o *PDF2TextInternalServerError) Error() string {
	return fmt.Sprintf("[POST /][%d] pDF2TextInternalServerError ", 500)
}

func (o *PDF2TextInternalServerError) String() string {
	return fmt.Sprintf("[POST /][%d] pDF2TextInternalServerError ", 500)
}

func (o *PDF2TextInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
