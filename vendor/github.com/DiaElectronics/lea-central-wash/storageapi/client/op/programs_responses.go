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
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/DiaElectronics/lea-central-wash/storageapi/model"
)

// ProgramsReader is a Reader for the Programs structure.
type ProgramsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProgramsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProgramsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewProgramsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProgramsOK creates a ProgramsOK with default headers values
func NewProgramsOK() *ProgramsOK {
	return &ProgramsOK{}
}

/*ProgramsOK handles this case with default header values.

OK
*/
type ProgramsOK struct {
	Payload []*model.ProgramInfo
}

func (o *ProgramsOK) Error() string {
	return fmt.Sprintf("[POST /programs][%d] programsOK  %+v", 200, o.Payload)
}

func (o *ProgramsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProgramsInternalServerError creates a ProgramsInternalServerError with default headers values
func NewProgramsInternalServerError() *ProgramsInternalServerError {
	return &ProgramsInternalServerError{}
}

/*ProgramsInternalServerError handles this case with default header values.

internal error
*/
type ProgramsInternalServerError struct {
}

func (o *ProgramsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /programs][%d] programsInternalServerError ", 500)
}

func (o *ProgramsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProgramsBody programs body
swagger:model ProgramsBody
*/
type ProgramsBody struct {

	// station ID
	// Required: true
	// Minimum: 1
	StationID *int64 `json:"stationID"`
}

// Validate validates this programs body
func (o *ProgramsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProgramsBody) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"stationID", "body", o.StationID); err != nil {
		return err
	}

	if err := validate.MinimumInt("args"+"."+"stationID", "body", int64(*o.StationID), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ProgramsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProgramsBody) UnmarshalBinary(b []byte) error {
	var res ProgramsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
