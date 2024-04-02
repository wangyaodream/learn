package pipeline

import (
	"net/http"
	"platform/services"
	"reflect"
)

type RequestPipline func(*ComponentContext)

var emptyPipline RequestPipline = func(*ComponentContext) {}

// func CreatePipeline(components ...MiddlewareComponent) RequestPipline {
//     f := emptyPipline
//     for i := len(components) -1; i >= 0; i-- {
//         currentComponent := components[i]
//         nextFunc := f
//         f = func(context *ComponentContext) {
//             if (context.error == nil) {
//                 currentComponent.ProcessRequest(context, nextFunc)
//             }
//         }
//         currentComponent.Init()
//     }
//     return f
// }

func CreatePipline(components ...interface{}) RequestPipline {
    f := emptyPipline
    for i := len(components) - 1; i >= 0; i-- {
        currentComponent := components[i]
        services.Populate(currentComponent)
        nextFunc := f
        if servComp, ok := currentComponent.(ServicesMiddlewareComponent); ok {
            f = createServiceDependentFunction(currentComponent, nextFunc) 
            servComp.Init()
        } else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
            f = func(context *ComponentContext) {
                if (context.error == nil) {
                    stdComp.ProcessRequest(context, nextFunc)
                }
            }
            stdComp.Init()
        } else {
            panic("Value is not a middleware component")
        }
    } 
    return f
}

func createServiceDependentFunction(component interface{},
        nextFunc RequestPipline) RequestPipline {
    method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
    if (method.IsValid()) {
        return func(context *ComponentContext) {
            if (context.error == nil) {
                _, err := services.CallForContext(context.Request.Context(),
                    method.Interface(), context, nextFunc)
                if (err != nil) {
                    context.Error(err)
                }
            }
        }
    } else {
        panic("No ProcessRequestWithServices method defined")
    }
}

func (pl RequestPipline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
    ctx := ComponentContext {
        Request: req,
        ResponseWriter: resp,
    }
    pl(&ctx)
    return ctx.error
}
