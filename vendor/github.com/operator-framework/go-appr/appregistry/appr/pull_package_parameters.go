// Code generated by go-swagger; DO NOT EDIT.

package appr

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPullPackageParams creates a new PullPackageParams object
// with the default values initialized.
func NewPullPackageParams() *PullPackageParams {
	var ()
	return &PullPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPullPackageParamsWithTimeout creates a new PullPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPullPackageParamsWithTimeout(timeout time.Duration) *PullPackageParams {
	var ()
	return &PullPackageParams{

		timeout: timeout,
	}
}

// NewPullPackageParamsWithContext creates a new PullPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewPullPackageParamsWithContext(ctx context.Context) *PullPackageParams {
	var ()
	return &PullPackageParams{

		Context: ctx,
	}
}

// NewPullPackageParamsWithHTTPClient creates a new PullPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPullPackageParamsWithHTTPClient(client *http.Client) *PullPackageParams {
	var ()
	return &PullPackageParams{
		HTTPClient: client,
	}
}

/*PullPackageParams contains all the parameters to send to the API endpoint
for the pull package operation typically these are written to a http.Request
*/
type PullPackageParams struct {

	/*Format
	  reponse format: json or blob

	*/
	Format *string
	/*MediaType
	  content type

	*/
	MediaType string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string
	/*Release
	  release name

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pull package params
func (o *PullPackageParams) WithTimeout(timeout time.Duration) *PullPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pull package params
func (o *PullPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pull package params
func (o *PullPackageParams) WithContext(ctx context.Context) *PullPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pull package params
func (o *PullPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pull package params
func (o *PullPackageParams) WithHTTPClient(client *http.Client) *PullPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pull package params
func (o *PullPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the pull package params
func (o *PullPackageParams) WithFormat(format *string) *PullPackageParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the pull package params
func (o *PullPackageParams) SetFormat(format *string) {
	o.Format = format
}

// WithMediaType adds the mediaType to the pull package params
func (o *PullPackageParams) WithMediaType(mediaType string) *PullPackageParams {
	o.SetMediaType(mediaType)
	return o
}

// SetMediaType adds the mediaType to the pull package params
func (o *PullPackageParams) SetMediaType(mediaType string) {
	o.MediaType = mediaType
}

// WithNamespace adds the namespace to the pull package params
func (o *PullPackageParams) WithNamespace(namespace string) *PullPackageParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the pull package params
func (o *PullPackageParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the pull package params
func (o *PullPackageParams) WithPackage(packageVar string) *PullPackageParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the pull package params
func (o *PullPackageParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRelease adds the release to the pull package params
func (o *PullPackageParams) WithRelease(release string) *PullPackageParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the pull package params
func (o *PullPackageParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *PullPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param media_type
	if err := r.SetPathParam("media_type", o.MediaType); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}