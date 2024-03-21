package pipeline

import (
	"net/http"
)

type RequestPipline func(*ComponentContext)

var emptyPipline RequestPipline = func(*ComponentContext) {}

func CreatePipeline(components ...MiddlewareComponent) RequestPipline {
    f := emptyPipline
    for i := len(components) -1; i >= 0; i-- {
        currentComponent := components[i]
        nextFunc := f
        f = func(context *ComponentContext) {
            if (context.error == nil) {
                currentComponent.ProcessRequest(context, nextFunc)
            }
        }
        currentComponent.Init()
    }
    return f
}

func (pl RequestPipline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
    ctx := ComponentContext {
        Request: req,
        ResponseWriter: resp,
    }
    pl(&ctx)
    return ctx.error
}
