package services

import (
	"reflect"
	"sync"
)

func AddTransient(factoryFunc interface{}) (err error) {
    return addService(Transient, factoryFunc)
}

func AddScoped(factoryFunc interface{}) (err error) {
    return addService(Scoped, factoryFunc)
}

func AddSingleton(factoryFunc interface{}) (err error) {
    factoryFuncVal := reflect.ValueOf(factoryFunc)
    if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {

    }
}

