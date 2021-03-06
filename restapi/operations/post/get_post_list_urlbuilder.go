// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetPostListURL generates an URL for the get post list operation
type GetPostListURL struct {
	ByTime    *string
	DateFrom  *string
	DateTo    *string
	EventType *string
	Limit     *int64
	Offset    *int64
	Order     *string
	SortBy    *string
	Text      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPostListURL) WithBasePath(bp string) *GetPostListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPostListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPostListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/post/list"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var byTimeQ string
	if o.ByTime != nil {
		byTimeQ = *o.ByTime
	}
	if byTimeQ != "" {
		qs.Set("byTime", byTimeQ)
	}

	var dateFromQ string
	if o.DateFrom != nil {
		dateFromQ = *o.DateFrom
	}
	if dateFromQ != "" {
		qs.Set("dateFrom", dateFromQ)
	}

	var dateToQ string
	if o.DateTo != nil {
		dateToQ = *o.DateTo
	}
	if dateToQ != "" {
		qs.Set("dateTo", dateToQ)
	}

	var eventTypeQ string
	if o.EventType != nil {
		eventTypeQ = *o.EventType
	}
	if eventTypeQ != "" {
		qs.Set("eventType", eventTypeQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var orderQ string
	if o.Order != nil {
		orderQ = *o.Order
	}
	if orderQ != "" {
		qs.Set("order", orderQ)
	}

	var sortByQ string
	if o.SortBy != nil {
		sortByQ = *o.SortBy
	}
	if sortByQ != "" {
		qs.Set("sortBy", sortByQ)
	}

	var textQ string
	if o.Text != nil {
		textQ = *o.Text
	}
	if textQ != "" {
		qs.Set("text", textQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPostListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPostListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPostListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPostListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPostListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetPostListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
