package main

import (
    "reflect"
    // "strings"
    // "fmt"
)

var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func getTypePath(t reflect.Type) (path string) {
    path = t.PkgPath()
    if (path == "") {
        path = "(built-in)"
    }
    return
}

// func printDetails(values ...interface{}) {
//     for _, elem := range values {
//         switch val := elem.(type) {
//         case Product:
//             Printfln("Product: Name: %v, Category: %v, Price: %v",
//                 val.Name, val.Category, val.Price)
//         case Customer:
//             Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
//         }
//     }
// }

// // 利用反射来支持多种不同类型的参数传入并处理
// func printDetails(values ...interface{}) {
//     for _, elem := range values {
//         fieldDetails := []string {}
//         elemType := reflect.TypeOf(elem)
//         elemValue := reflect.ValueOf(elem)
//         if elemType.Kind() == reflect.Struct {
//             for i := 0; i < elemType.NumField(); i++ {
//                 fieldName := elemType.Field(i).Name
//                 fieldVal := elemValue.Field(i)
//                 fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName,fieldVal))
//             }
//             Printfln("%v(%v)", elemType.Name(), strings.Join(fieldDetails, ", "))
//         } else {
//             Printfln("%v: %v", elemType.Name(), elemValue)
//         }
//     }
// }

// // 使用Type的一些特性
// func printDetails(values ...interface{}) {
//     for _, elem := range values {
//         elemType := reflect.TypeOf(elem)
//         Printfln("Name: %v, PkgPath: %v, Kind: %v",
//             elemType.Name(), getTypePath(elemType), elemType.Kind())
//     }
// }

// 使用Value的一些特性
func printDetails(values ...interface{}) {
    for _, elem := range values {
        elemValue := reflect.ValueOf(elem)
        elemType := reflect.TypeOf(elem)
        if (elemType == intPtrType) {
            Printfln("Pointer to Int: %v", elemValue.Elem().Int())
        } else if (elemType == byteSliceType) {
            Printfln("Byte slice: %v", elemValue.Bytes())
        } else {
            switch elemValue.Kind() {
                case reflect.Bool:
                    var val bool = elemValue.Bool()
                    Printfln("Bool: %v", val)
                case reflect.Int:
                    var val int64 = elemValue.Int()
                    Printfln("Int: %v", val)
                case reflect.Float32, reflect.Float64:
                    var val float64 = elemValue.Float()
                    Printfln("Float: %v", val)
                case reflect.String:
                    var val string = elemValue.String()
                    Printfln("String: %v", val)
                // case reflect.Ptr:
                //     var val reflect.Value = elemValue.Elem()
                //     if (val.Kind() == reflect.Int) {
                //         Printfln("Pointer to Int: %v", val.Int())
                //     }
                default:
                    Printfln("Other: %v", elemValue.String())
            }
        }
    }
}

type Payment struct {
    Currency string
    Amount float64
}

func main() {
    product := Product {
        Name: "Kayak", Category: "Watersports", Price: 279,
    }
    customer := Customer { Name: "Alice", City: "New York"}
    // 可以使用不同的类型
    // printDetails(product, customer)
    payment := Payment { Currency: "USD", Amount: 100.6}
    number := 100
    slice := []byte("Alice")
    printDetails(product, customer, payment, 11, true, &number, "Alice", slice)
}

