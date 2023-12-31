// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/akhmaos/omni-case/internal/api/restapi/models"
)

// PostProcessItemsOKCode is the HTTP code returned for type PostProcessItemsOK
const PostProcessItemsOKCode int = 200

/*
PostProcessItemsOK ok

swagger:response postProcessItemsOK
*/
type PostProcessItemsOK struct {

	/*
	  In: Body
	*/
	Payload *PostProcessItemsOKBody `json:"body,omitempty"`
}

// NewPostProcessItemsOK creates PostProcessItemsOK with default headers values
func NewPostProcessItemsOK() *PostProcessItemsOK {

	return &PostProcessItemsOK{}
}

// WithPayload adds the payload to the post process items o k response
func (o *PostProcessItemsOK) WithPayload(payload *PostProcessItemsOKBody) *PostProcessItemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post process items o k response
func (o *PostProcessItemsOK) SetPayload(payload *PostProcessItemsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProcessItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProcessItemsInternalServerErrorCode is the HTTP code returned for type PostProcessItemsInternalServerError
const PostProcessItemsInternalServerErrorCode int = 500

/*
PostProcessItemsInternalServerError internal server error

swagger:response postProcessItemsInternalServerError
*/
type PostProcessItemsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.InternalServerError `json:"body,omitempty"`
}

// NewPostProcessItemsInternalServerError creates PostProcessItemsInternalServerError with default headers values
func NewPostProcessItemsInternalServerError() *PostProcessItemsInternalServerError {

	return &PostProcessItemsInternalServerError{}
}

// WithPayload adds the payload to the post process items internal server error response
func (o *PostProcessItemsInternalServerError) WithPayload(payload models.InternalServerError) *PostProcessItemsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post process items internal server error response
func (o *PostProcessItemsInternalServerError) SetPayload(payload models.InternalServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProcessItemsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PostProcessItemsDefault error

swagger:response postProcessItemsDefault
*/
type PostProcessItemsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProcessItemsDefault creates PostProcessItemsDefault with default headers values
func NewPostProcessItemsDefault(code int) *PostProcessItemsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProcessItemsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post process items default response
func (o *PostProcessItemsDefault) WithStatusCode(code int) *PostProcessItemsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post process items default response
func (o *PostProcessItemsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post process items default response
func (o *PostProcessItemsDefault) WithPayload(payload *models.Error) *PostProcessItemsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post process items default response
func (o *PostProcessItemsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProcessItemsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
