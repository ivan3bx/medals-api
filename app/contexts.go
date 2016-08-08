//************************************************************************//
// API "cellar": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/ivan3bx/medal-service/design
// --out=$(GOPATH)/src/github.com/ivan3bx/medal-service
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// IndexMedalsContext provides the medals index action context.
type IndexMedalsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Sort *string
}

// NewIndexMedalsContext parses the incoming request URL and body, performs validations and creates the
// context used by the medals controller index action.
func NewIndexMedalsContext(ctx context.Context, service *goa.Service) (*IndexMedalsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := IndexMedalsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramSort := req.Params["sort"]
	if len(paramSort) > 0 {
		rawSort := paramSort[0]
		rctx.Sort = &rawSort
		if rctx.Sort != nil {
			if !(*rctx.Sort == "country" || *rctx.Sort == "total" || *rctx.Sort == "gold" || *rctx.Sort == "silver" || *rctx.Sort == "bronze") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`sort`, *rctx.Sort, []interface{}{"country", "total", "gold", "silver", "bronze"}))
			}
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *IndexMedalsContext) OK(r MedalsCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.medals+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *IndexMedalsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
