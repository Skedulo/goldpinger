// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bloomberg/goldpinger/pkg/models"
)

// CheckServicePodsReader is a Reader for the CheckServicePods structure.
type CheckServicePodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckServicePodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckServicePodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckServicePodsOK creates a CheckServicePodsOK with default headers values
func NewCheckServicePodsOK() *CheckServicePodsOK {
	return &CheckServicePodsOK{}
}

/*CheckServicePodsOK handles this case with default header values.

Success, return response
*/
type CheckServicePodsOK struct {
	Payload *models.CheckResults
}

func (o *CheckServicePodsOK) Error() string {
	return fmt.Sprintf("[GET /check][%d] checkServicePodsOK  %+v", 200, o.Payload)
}

func (o *CheckServicePodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CheckResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
