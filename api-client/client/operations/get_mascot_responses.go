// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/meetup-terraform-provider/api-client/models"
)

// GetMascotReader is a Reader for the GetMascot structure.
type GetMascotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMascotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMascotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetMascotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMascotOK creates a GetMascotOK with default headers values
func NewGetMascotOK() *GetMascotOK {
	return &GetMascotOK{}
}

/*GetMascotOK handles this case with default header values.

Success
*/
type GetMascotOK struct {
	Payload *models.MascotResponse
}

func (o *GetMascotOK) Error() string {
	return fmt.Sprintf("[GET /{id}][%d] getMascotOK  %+v", 200, o.Payload)
}

func (o *GetMascotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MascotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMascotNotFound creates a GetMascotNotFound with default headers values
func NewGetMascotNotFound() *GetMascotNotFound {
	return &GetMascotNotFound{}
}

/*GetMascotNotFound handles this case with default header values.

Not Found
*/
type GetMascotNotFound struct {
}

func (o *GetMascotNotFound) Error() string {
	return fmt.Sprintf("[GET /{id}][%d] getMascotNotFound ", 404)
}

func (o *GetMascotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}