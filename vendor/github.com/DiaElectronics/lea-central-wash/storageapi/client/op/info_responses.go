// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// InfoReader is a Reader for the Info structure.
type InfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInfoOK creates a InfoOK with default headers values
func NewInfoOK() *InfoOK {
	return &InfoOK{}
}

/*InfoOK handles this case with default header values.

OK
*/
type InfoOK struct {
	Payload string
}

func (o *InfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] infoOK  %+v", 200, o.Payload)
}

func (o *InfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
