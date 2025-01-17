// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DownloadHostLogsHandlerFunc turns a function with the right signature into a download host logs handler
type DownloadHostLogsHandlerFunc func(DownloadHostLogsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadHostLogsHandlerFunc) Handle(params DownloadHostLogsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DownloadHostLogsHandler interface for that can handle valid download host logs params
type DownloadHostLogsHandler interface {
	Handle(DownloadHostLogsParams, interface{}) middleware.Responder
}

// NewDownloadHostLogs creates a new http.Handler for the download host logs operation
func NewDownloadHostLogs(ctx *middleware.Context, handler DownloadHostLogsHandler) *DownloadHostLogs {
	return &DownloadHostLogs{Context: ctx, Handler: handler}
}

/* DownloadHostLogs swagger:route GET /v1/clusters/{cluster_id}/hosts/{host_id}/logs installer downloadHostLogs

Download host logs.

*/
type DownloadHostLogs struct {
	Context *middleware.Context
	Handler DownloadHostLogsHandler
}

func (o *DownloadHostLogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDownloadHostLogsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
