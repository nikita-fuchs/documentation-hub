// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetOracleQueryByPubkeyAndQueryIDReader is a Reader for the GetOracleQueryByPubkeyAndQueryID structure.
type GetOracleQueryByPubkeyAndQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleQueryByPubkeyAndQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOracleQueryByPubkeyAndQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOracleQueryByPubkeyAndQueryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOracleQueryByPubkeyAndQueryIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOracleQueryByPubkeyAndQueryIDOK creates a GetOracleQueryByPubkeyAndQueryIDOK with default headers values
func NewGetOracleQueryByPubkeyAndQueryIDOK() *GetOracleQueryByPubkeyAndQueryIDOK {
	return &GetOracleQueryByPubkeyAndQueryIDOK{}
}

/*GetOracleQueryByPubkeyAndQueryIDOK handles this case with default header values.

Successful operation
*/
type GetOracleQueryByPubkeyAndQueryIDOK struct {
	Payload *models.OracleQuery
}

func (o *GetOracleQueryByPubkeyAndQueryIDOK) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}/queries/{query-id}][%d] getOracleQueryByPubkeyAndQueryIdOK  %+v", 200, o.Payload)
}

func (o *GetOracleQueryByPubkeyAndQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OracleQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleQueryByPubkeyAndQueryIDBadRequest creates a GetOracleQueryByPubkeyAndQueryIDBadRequest with default headers values
func NewGetOracleQueryByPubkeyAndQueryIDBadRequest() *GetOracleQueryByPubkeyAndQueryIDBadRequest {
	return &GetOracleQueryByPubkeyAndQueryIDBadRequest{}
}

/*GetOracleQueryByPubkeyAndQueryIDBadRequest handles this case with default header values.

Invalid public key or query ID
*/
type GetOracleQueryByPubkeyAndQueryIDBadRequest struct {
	Payload *models.Error
}

func (o *GetOracleQueryByPubkeyAndQueryIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}/queries/{query-id}][%d] getOracleQueryByPubkeyAndQueryIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetOracleQueryByPubkeyAndQueryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleQueryByPubkeyAndQueryIDNotFound creates a GetOracleQueryByPubkeyAndQueryIDNotFound with default headers values
func NewGetOracleQueryByPubkeyAndQueryIDNotFound() *GetOracleQueryByPubkeyAndQueryIDNotFound {
	return &GetOracleQueryByPubkeyAndQueryIDNotFound{}
}

/*GetOracleQueryByPubkeyAndQueryIDNotFound handles this case with default header values.

Oracle query not found
*/
type GetOracleQueryByPubkeyAndQueryIDNotFound struct {
	Payload *models.Error
}

func (o *GetOracleQueryByPubkeyAndQueryIDNotFound) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}/queries/{query-id}][%d] getOracleQueryByPubkeyAndQueryIdNotFound  %+v", 404, o.Payload)
}

func (o *GetOracleQueryByPubkeyAndQueryIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
