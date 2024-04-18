// Code generated by go-swagger; DO NOT EDIT.

package service

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
	"github.com/go-openapi/swag"
)

// NewPDF2TextParams creates a new PDF2TextParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPDF2TextParams() *PDF2TextParams {
	return &PDF2TextParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPDF2TextParamsWithTimeout creates a new PDF2TextParams object
// with the ability to set a timeout on a request.
func NewPDF2TextParamsWithTimeout(timeout time.Duration) *PDF2TextParams {
	return &PDF2TextParams{
		timeout: timeout,
	}
}

// NewPDF2TextParamsWithContext creates a new PDF2TextParams object
// with the ability to set a context for a request.
func NewPDF2TextParamsWithContext(ctx context.Context) *PDF2TextParams {
	return &PDF2TextParams{
		Context: ctx,
	}
}

// NewPDF2TextParamsWithHTTPClient creates a new PDF2TextParams object
// with the ability to set a custom HTTPClient for a request.
func NewPDF2TextParamsWithHTTPClient(client *http.Client) *PDF2TextParams {
	return &PDF2TextParams{
		HTTPClient: client,
	}
}

/*
PDF2TextParams contains all the parameters to send to the API endpoint

	for the p d f2 text operation.

	Typically these are written to a http.Request.
*/
type PDF2TextParams struct {

	/* Bbox.

	   output bounding box for each word and page size to html. Sets `htmlmeta`
	*/
	Bbox *bool

	/* BboxLayout.

	   like `bbox` but with extra layout bounding box data.  Sets `htmlmeta`
	*/
	BboxLayout *bool

	/* Colspacing.

	   how much spacing we allow after a word before considering adjacent text to be a new column, as a fraction of the font size (default is 0.7, old releases had a 0.3 default)
	*/
	Colspacing *float64

	/* Cropbox.

	   use the crop box rather than media box
	*/
	Cropbox *bool

	/* Enc.

	   output text encoding name

	   Default: "UTF-8"
	*/
	Enc *string

	/* Eol.

	   output end-of-line convention (unix, dos, or mac)
	*/
	Eol *string

	/* F.

	   first page to convert
	*/
	F *int64

	/* Fixed.

	   assume fixed-pitch (or tabular) text, in fp
	*/
	Fixed *int64

	/* H.

	   height of crop area in pixels (default is 0)
	*/
	H *int64

	/* Htmlmeta.

	   generate a simple HTML file, including the meta information
	*/
	Htmlmeta *bool

	/* L.

	   last page to convert
	*/
	L *int64

	/* Layout.

	   maintain original physical layout
	*/
	Layout *bool

	/* Nodiag.

	   discard diagonal text
	*/
	Nodiag *bool

	/* Opw.

	   owner password (for encrypted files)
	*/
	Opw *string

	/* Pgbrk.

	   insert page breaks between pages
	*/
	Pgbrk *bool

	/* R.

	   resolution, in DPI (default is 72)
	*/
	R *int64

	/* Raw.

	   keep strings in content stream order
	*/
	Raw *bool

	/* Tsv.

	   generate a simple TSV file, including the meta information for bounding boxes
	*/
	Tsv *bool

	/* Upw.

	   user password (for encrypted files)
	*/
	Upw *string

	/* W.

	   width of crop area in pixels (default is 0)
	*/
	W *int64

	/* X.

	   x-coordinate of the crop area top left corner
	*/
	X *int64

	/* Y.

	   y-coordinate of the crop area top left corner
	*/
	Y *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the p d f2 text params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PDF2TextParams) WithDefaults() *PDF2TextParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the p d f2 text params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PDF2TextParams) SetDefaults() {
	var (
		encDefault = string("UTF-8")
	)

	val := PDF2TextParams{
		Enc: &encDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the p d f2 text params
func (o *PDF2TextParams) WithTimeout(timeout time.Duration) *PDF2TextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p d f2 text params
func (o *PDF2TextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p d f2 text params
func (o *PDF2TextParams) WithContext(ctx context.Context) *PDF2TextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p d f2 text params
func (o *PDF2TextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p d f2 text params
func (o *PDF2TextParams) WithHTTPClient(client *http.Client) *PDF2TextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p d f2 text params
func (o *PDF2TextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBbox adds the bbox to the p d f2 text params
func (o *PDF2TextParams) WithBbox(bbox *bool) *PDF2TextParams {
	o.SetBbox(bbox)
	return o
}

// SetBbox adds the bbox to the p d f2 text params
func (o *PDF2TextParams) SetBbox(bbox *bool) {
	o.Bbox = bbox
}

// WithBboxLayout adds the bboxLayout to the p d f2 text params
func (o *PDF2TextParams) WithBboxLayout(bboxLayout *bool) *PDF2TextParams {
	o.SetBboxLayout(bboxLayout)
	return o
}

// SetBboxLayout adds the bboxLayout to the p d f2 text params
func (o *PDF2TextParams) SetBboxLayout(bboxLayout *bool) {
	o.BboxLayout = bboxLayout
}

// WithColspacing adds the colspacing to the p d f2 text params
func (o *PDF2TextParams) WithColspacing(colspacing *float64) *PDF2TextParams {
	o.SetColspacing(colspacing)
	return o
}

// SetColspacing adds the colspacing to the p d f2 text params
func (o *PDF2TextParams) SetColspacing(colspacing *float64) {
	o.Colspacing = colspacing
}

// WithCropbox adds the cropbox to the p d f2 text params
func (o *PDF2TextParams) WithCropbox(cropbox *bool) *PDF2TextParams {
	o.SetCropbox(cropbox)
	return o
}

// SetCropbox adds the cropbox to the p d f2 text params
func (o *PDF2TextParams) SetCropbox(cropbox *bool) {
	o.Cropbox = cropbox
}

// WithEnc adds the enc to the p d f2 text params
func (o *PDF2TextParams) WithEnc(enc *string) *PDF2TextParams {
	o.SetEnc(enc)
	return o
}

// SetEnc adds the enc to the p d f2 text params
func (o *PDF2TextParams) SetEnc(enc *string) {
	o.Enc = enc
}

// WithEol adds the eol to the p d f2 text params
func (o *PDF2TextParams) WithEol(eol *string) *PDF2TextParams {
	o.SetEol(eol)
	return o
}

// SetEol adds the eol to the p d f2 text params
func (o *PDF2TextParams) SetEol(eol *string) {
	o.Eol = eol
}

// WithF adds the f to the p d f2 text params
func (o *PDF2TextParams) WithF(f *int64) *PDF2TextParams {
	o.SetF(f)
	return o
}

// SetF adds the f to the p d f2 text params
func (o *PDF2TextParams) SetF(f *int64) {
	o.F = f
}

// WithFixed adds the fixed to the p d f2 text params
func (o *PDF2TextParams) WithFixed(fixed *int64) *PDF2TextParams {
	o.SetFixed(fixed)
	return o
}

// SetFixed adds the fixed to the p d f2 text params
func (o *PDF2TextParams) SetFixed(fixed *int64) {
	o.Fixed = fixed
}

// WithH adds the h to the p d f2 text params
func (o *PDF2TextParams) WithH(h *int64) *PDF2TextParams {
	o.SetH(h)
	return o
}

// SetH adds the h to the p d f2 text params
func (o *PDF2TextParams) SetH(h *int64) {
	o.H = h
}

// WithHtmlmeta adds the htmlmeta to the p d f2 text params
func (o *PDF2TextParams) WithHtmlmeta(htmlmeta *bool) *PDF2TextParams {
	o.SetHtmlmeta(htmlmeta)
	return o
}

// SetHtmlmeta adds the htmlmeta to the p d f2 text params
func (o *PDF2TextParams) SetHtmlmeta(htmlmeta *bool) {
	o.Htmlmeta = htmlmeta
}

// WithL adds the l to the p d f2 text params
func (o *PDF2TextParams) WithL(l *int64) *PDF2TextParams {
	o.SetL(l)
	return o
}

// SetL adds the l to the p d f2 text params
func (o *PDF2TextParams) SetL(l *int64) {
	o.L = l
}

// WithLayout adds the layout to the p d f2 text params
func (o *PDF2TextParams) WithLayout(layout *bool) *PDF2TextParams {
	o.SetLayout(layout)
	return o
}

// SetLayout adds the layout to the p d f2 text params
func (o *PDF2TextParams) SetLayout(layout *bool) {
	o.Layout = layout
}

// WithNodiag adds the nodiag to the p d f2 text params
func (o *PDF2TextParams) WithNodiag(nodiag *bool) *PDF2TextParams {
	o.SetNodiag(nodiag)
	return o
}

// SetNodiag adds the nodiag to the p d f2 text params
func (o *PDF2TextParams) SetNodiag(nodiag *bool) {
	o.Nodiag = nodiag
}

// WithOpw adds the opw to the p d f2 text params
func (o *PDF2TextParams) WithOpw(opw *string) *PDF2TextParams {
	o.SetOpw(opw)
	return o
}

// SetOpw adds the opw to the p d f2 text params
func (o *PDF2TextParams) SetOpw(opw *string) {
	o.Opw = opw
}

// WithPgbrk adds the pgbrk to the p d f2 text params
func (o *PDF2TextParams) WithPgbrk(pgbrk *bool) *PDF2TextParams {
	o.SetPgbrk(pgbrk)
	return o
}

// SetPgbrk adds the pgbrk to the p d f2 text params
func (o *PDF2TextParams) SetPgbrk(pgbrk *bool) {
	o.Pgbrk = pgbrk
}

// WithR adds the r to the p d f2 text params
func (o *PDF2TextParams) WithR(r *int64) *PDF2TextParams {
	o.SetR(r)
	return o
}

// SetR adds the r to the p d f2 text params
func (o *PDF2TextParams) SetR(r *int64) {
	o.R = r
}

// WithRaw adds the raw to the p d f2 text params
func (o *PDF2TextParams) WithRaw(raw *bool) *PDF2TextParams {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the p d f2 text params
func (o *PDF2TextParams) SetRaw(raw *bool) {
	o.Raw = raw
}

// WithTsv adds the tsv to the p d f2 text params
func (o *PDF2TextParams) WithTsv(tsv *bool) *PDF2TextParams {
	o.SetTsv(tsv)
	return o
}

// SetTsv adds the tsv to the p d f2 text params
func (o *PDF2TextParams) SetTsv(tsv *bool) {
	o.Tsv = tsv
}

// WithUpw adds the upw to the p d f2 text params
func (o *PDF2TextParams) WithUpw(upw *string) *PDF2TextParams {
	o.SetUpw(upw)
	return o
}

// SetUpw adds the upw to the p d f2 text params
func (o *PDF2TextParams) SetUpw(upw *string) {
	o.Upw = upw
}

// WithW adds the w to the p d f2 text params
func (o *PDF2TextParams) WithW(w *int64) *PDF2TextParams {
	o.SetW(w)
	return o
}

// SetW adds the w to the p d f2 text params
func (o *PDF2TextParams) SetW(w *int64) {
	o.W = w
}

// WithX adds the x to the p d f2 text params
func (o *PDF2TextParams) WithX(x *int64) *PDF2TextParams {
	o.SetX(x)
	return o
}

// SetX adds the x to the p d f2 text params
func (o *PDF2TextParams) SetX(x *int64) {
	o.X = x
}

// WithY adds the y to the p d f2 text params
func (o *PDF2TextParams) WithY(y *int64) *PDF2TextParams {
	o.SetY(y)
	return o
}

// SetY adds the y to the p d f2 text params
func (o *PDF2TextParams) SetY(y *int64) {
	o.Y = y
}

// WriteToRequest writes these params to a swagger request
func (o *PDF2TextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bbox != nil {

		// query param bbox
		var qrBbox bool

		if o.Bbox != nil {
			qrBbox = *o.Bbox
		}
		qBbox := swag.FormatBool(qrBbox)
		if qBbox != "" {

			if err := r.SetQueryParam("bbox", qBbox); err != nil {
				return err
			}
		}
	}

	if o.BboxLayout != nil {

		// query param bbox-layout
		var qrBboxLayout bool

		if o.BboxLayout != nil {
			qrBboxLayout = *o.BboxLayout
		}
		qBboxLayout := swag.FormatBool(qrBboxLayout)
		if qBboxLayout != "" {

			if err := r.SetQueryParam("bbox-layout", qBboxLayout); err != nil {
				return err
			}
		}
	}

	if o.Colspacing != nil {

		// query param colspacing
		var qrColspacing float64

		if o.Colspacing != nil {
			qrColspacing = *o.Colspacing
		}
		qColspacing := swag.FormatFloat64(qrColspacing)
		if qColspacing != "" {

			if err := r.SetQueryParam("colspacing", qColspacing); err != nil {
				return err
			}
		}
	}

	if o.Cropbox != nil {

		// query param cropbox
		var qrCropbox bool

		if o.Cropbox != nil {
			qrCropbox = *o.Cropbox
		}
		qCropbox := swag.FormatBool(qrCropbox)
		if qCropbox != "" {

			if err := r.SetQueryParam("cropbox", qCropbox); err != nil {
				return err
			}
		}
	}

	if o.Enc != nil {

		// query param enc
		var qrEnc string

		if o.Enc != nil {
			qrEnc = *o.Enc
		}
		qEnc := qrEnc
		if qEnc != "" {

			if err := r.SetQueryParam("enc", qEnc); err != nil {
				return err
			}
		}
	}

	if o.Eol != nil {

		// query param eol
		var qrEol string

		if o.Eol != nil {
			qrEol = *o.Eol
		}
		qEol := qrEol
		if qEol != "" {

			if err := r.SetQueryParam("eol", qEol); err != nil {
				return err
			}
		}
	}

	if o.F != nil {

		// query param f
		var qrF int64

		if o.F != nil {
			qrF = *o.F
		}
		qF := swag.FormatInt64(qrF)
		if qF != "" {

			if err := r.SetQueryParam("f", qF); err != nil {
				return err
			}
		}
	}

	if o.Fixed != nil {

		// query param fixed
		var qrFixed int64

		if o.Fixed != nil {
			qrFixed = *o.Fixed
		}
		qFixed := swag.FormatInt64(qrFixed)
		if qFixed != "" {

			if err := r.SetQueryParam("fixed", qFixed); err != nil {
				return err
			}
		}
	}

	if o.H != nil {

		// query param h
		var qrH int64

		if o.H != nil {
			qrH = *o.H
		}
		qH := swag.FormatInt64(qrH)
		if qH != "" {

			if err := r.SetQueryParam("h", qH); err != nil {
				return err
			}
		}
	}

	if o.Htmlmeta != nil {

		// query param htmlmeta
		var qrHtmlmeta bool

		if o.Htmlmeta != nil {
			qrHtmlmeta = *o.Htmlmeta
		}
		qHtmlmeta := swag.FormatBool(qrHtmlmeta)
		if qHtmlmeta != "" {

			if err := r.SetQueryParam("htmlmeta", qHtmlmeta); err != nil {
				return err
			}
		}
	}

	if o.L != nil {

		// query param l
		var qrL int64

		if o.L != nil {
			qrL = *o.L
		}
		qL := swag.FormatInt64(qrL)
		if qL != "" {

			if err := r.SetQueryParam("l", qL); err != nil {
				return err
			}
		}
	}

	if o.Layout != nil {

		// query param layout
		var qrLayout bool

		if o.Layout != nil {
			qrLayout = *o.Layout
		}
		qLayout := swag.FormatBool(qrLayout)
		if qLayout != "" {

			if err := r.SetQueryParam("layout", qLayout); err != nil {
				return err
			}
		}
	}

	if o.Nodiag != nil {

		// query param nodiag
		var qrNodiag bool

		if o.Nodiag != nil {
			qrNodiag = *o.Nodiag
		}
		qNodiag := swag.FormatBool(qrNodiag)
		if qNodiag != "" {

			if err := r.SetQueryParam("nodiag", qNodiag); err != nil {
				return err
			}
		}
	}

	if o.Opw != nil {

		// query param opw
		var qrOpw string

		if o.Opw != nil {
			qrOpw = *o.Opw
		}
		qOpw := qrOpw
		if qOpw != "" {

			if err := r.SetQueryParam("opw", qOpw); err != nil {
				return err
			}
		}
	}

	if o.Pgbrk != nil {

		// query param pgbrk
		var qrPgbrk bool

		if o.Pgbrk != nil {
			qrPgbrk = *o.Pgbrk
		}
		qPgbrk := swag.FormatBool(qrPgbrk)
		if qPgbrk != "" {

			if err := r.SetQueryParam("pgbrk", qPgbrk); err != nil {
				return err
			}
		}
	}

	if o.R != nil {

		// query param r
		var qrR int64

		if o.R != nil {
			qrR = *o.R
		}
		qR := swag.FormatInt64(qrR)
		if qR != "" {

			if err := r.SetQueryParam("r", qR); err != nil {
				return err
			}
		}
	}

	if o.Raw != nil {

		// query param raw
		var qrRaw bool

		if o.Raw != nil {
			qrRaw = *o.Raw
		}
		qRaw := swag.FormatBool(qrRaw)
		if qRaw != "" {

			if err := r.SetQueryParam("raw", qRaw); err != nil {
				return err
			}
		}
	}

	if o.Tsv != nil {

		// query param tsv
		var qrTsv bool

		if o.Tsv != nil {
			qrTsv = *o.Tsv
		}
		qTsv := swag.FormatBool(qrTsv)
		if qTsv != "" {

			if err := r.SetQueryParam("tsv", qTsv); err != nil {
				return err
			}
		}
	}

	if o.Upw != nil {

		// query param upw
		var qrUpw string

		if o.Upw != nil {
			qrUpw = *o.Upw
		}
		qUpw := qrUpw
		if qUpw != "" {

			if err := r.SetQueryParam("upw", qUpw); err != nil {
				return err
			}
		}
	}

	if o.W != nil {

		// query param w
		var qrW int64

		if o.W != nil {
			qrW = *o.W
		}
		qW := swag.FormatInt64(qrW)
		if qW != "" {

			if err := r.SetQueryParam("w", qW); err != nil {
				return err
			}
		}
	}

	if o.X != nil {

		// query param x
		var qrX int64

		if o.X != nil {
			qrX = *o.X
		}
		qX := swag.FormatInt64(qrX)
		if qX != "" {

			if err := r.SetQueryParam("x", qX); err != nil {
				return err
			}
		}
	}

	if o.Y != nil {

		// query param y
		var qrY int64

		if o.Y != nil {
			qrY = *o.Y
		}
		qY := swag.FormatInt64(qrY)
		if qY != "" {

			if err := r.SetQueryParam("y", qY); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}