// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetKasseReader is a Reader for the SetKasse structure.
type SetKasseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetKasseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSetKasseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSetKasseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetKasseNoContent creates a SetKasseNoContent with default headers values
func NewSetKasseNoContent() *SetKasseNoContent {
	return &SetKasseNoContent{}
}

/*SetKasseNoContent handles this case with default header values.

OK
*/
type SetKasseNoContent struct {
}

func (o *SetKasseNoContent) Error() string {
	return fmt.Sprintf("[POST /set-kasse][%d] setKasseNoContent ", 204)
}

func (o *SetKasseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetKasseInternalServerError creates a SetKasseInternalServerError with default headers values
func NewSetKasseInternalServerError() *SetKasseInternalServerError {
	return &SetKasseInternalServerError{}
}

/*SetKasseInternalServerError handles this case with default header values.

internal error
*/
type SetKasseInternalServerError struct {
}

func (o *SetKasseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /set-kasse][%d] setKasseInternalServerError ", 500)
}

func (o *SetKasseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
