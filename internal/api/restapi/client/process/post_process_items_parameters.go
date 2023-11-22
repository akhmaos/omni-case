// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostProcessItemsParams creates a new PostProcessItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProcessItemsParams() *PostProcessItemsParams {
	return &PostProcessItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProcessItemsParamsWithTimeout creates a new PostProcessItemsParams object
// with the ability to set a timeout on a request.
func NewPostProcessItemsParamsWithTimeout(timeout time.Duration) *PostProcessItemsParams {
	return &PostProcessItemsParams{
		timeout: timeout,
	}
}

// NewPostProcessItemsParamsWithContext creates a new PostProcessItemsParams object
// with the ability to set a context for a request.
func NewPostProcessItemsParamsWithContext(ctx context.Context) *PostProcessItemsParams {
	return &PostProcessItemsParams{
		Context: ctx,
	}
}

// NewPostProcessItemsParamsWithHTTPClient creates a new PostProcessItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProcessItemsParamsWithHTTPClient(client *http.Client) *PostProcessItemsParams {
	return &PostProcessItemsParams{
		HTTPClient: client,
	}
}

/*
PostProcessItemsParams contains all the parameters to send to the API endpoint

	for the post process items operation.

	Typically these are written to a http.Request.
*/
type PostProcessItemsParams struct {

	// Body.
	Body PostProcessItemsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post process items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessItemsParams) WithDefaults() *PostProcessItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post process items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post process items params
func (o *PostProcessItemsParams) WithTimeout(timeout time.Duration) *PostProcessItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post process items params
func (o *PostProcessItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post process items params
func (o *PostProcessItemsParams) WithContext(ctx context.Context) *PostProcessItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post process items params
func (o *PostProcessItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post process items params
func (o *PostProcessItemsParams) WithHTTPClient(client *http.Client) *PostProcessItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post process items params
func (o *PostProcessItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post process items params
func (o *PostProcessItemsParams) WithBody(body PostProcessItemsBody) *PostProcessItemsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post process items params
func (o *PostProcessItemsParams) SetBody(body PostProcessItemsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostProcessItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}