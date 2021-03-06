// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/DiaElectronics/lea-central-wash/storageapi/model"
)

// LoadRelayReader is a Reader for the LoadRelay structure.
type LoadRelayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoadRelayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoadRelayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewLoadRelayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLoadRelayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoadRelayOK creates a LoadRelayOK with default headers values
func NewLoadRelayOK() *LoadRelayOK {
	return &LoadRelayOK{}
}

/*LoadRelayOK handles this case with default header values.

OK
*/
type LoadRelayOK struct {
	Payload *model.RelayReport
}

func (o *LoadRelayOK) Error() string {
	return fmt.Sprintf("[POST /load-relay][%d] loadRelayOK  %+v", 200, o.Payload)
}

func (o *LoadRelayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RelayReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoadRelayNotFound creates a LoadRelayNotFound with default headers values
func NewLoadRelayNotFound() *LoadRelayNotFound {
	return &LoadRelayNotFound{}
}

/*LoadRelayNotFound handles this case with default header values.

not found
*/
type LoadRelayNotFound struct {
}

func (o *LoadRelayNotFound) Error() string {
	return fmt.Sprintf("[POST /load-relay][%d] loadRelayNotFound ", 404)
}

func (o *LoadRelayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoadRelayInternalServerError creates a LoadRelayInternalServerError with default headers values
func NewLoadRelayInternalServerError() *LoadRelayInternalServerError {
	return &LoadRelayInternalServerError{}
}

/*LoadRelayInternalServerError handles this case with default header values.

internal error
*/
type LoadRelayInternalServerError struct {
}

func (o *LoadRelayInternalServerError) Error() string {
	return fmt.Sprintf("[POST /load-relay][%d] loadRelayInternalServerError ", 500)
}

func (o *LoadRelayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*LoadRelayBody load relay body
swagger:model LoadRelayBody
*/
type LoadRelayBody struct {

	// hash
	// Required: true
	Hash model.Hash `json:"hash"`
}

// Validate validates this load relay body
func (o *LoadRelayBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoadRelayBody) validateHash(formats strfmt.Registry) error {

	if err := o.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "hash")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoadRelayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoadRelayBody) UnmarshalBinary(b []byte) error {
	var res LoadRelayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
