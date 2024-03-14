package main

import (
	"reflect"
	"strings"
	// "fmt"
)

func IsInt(v reflect.Value) bool {
    switch v.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            return true
    }
    return false
}

func IsFloat(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.Float32, reflect.Float64:
        return true
    }
    return false
}

func convert(src, target interface{}) (result interface{}, assigned bool) {
    srcVal := reflect.ValueOf(src)
    targetVal := reflect.ValueOf(target)
    if (srcVal.Type().ConvertibleTo(targetVal.Type())) {
        if (IsInt(targetVal) && IsInt(srcVal)) && targetVal.OverflowInt(srcVal.Int()) {
            Printfln("Int overflow")
            return src, false
        } else if (IsFloat(targetVal) && IsFloat(srcVal) && targetVal.OverflowFloat(srcVal.Float())) {

            Printfln("Float overflow")
            return src, false
        }
        result = srcVal.Convert(targetVal.Type()).Interface()
        assigned = true
    } else {
        result = src
    }
    return
}

func contains(slice interface{}, target interface{}) (found bool) {
    sliceVal := reflect.ValueOf(slice)
    targetType := reflect.TypeOf(target)
    if (sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem().Comparable() &&
        targetType.Comparable()) {
        for i := 0; i < sliceVal.Len(); i++ {
            if sliceVal.Index(i).Interface() == target {
                found = true
            }
        }
    }
    return
}

func setAll(src interface{}, targets ...interface{}) {
    srcVal := reflect.ValueOf(src)
    for _, target := range targets {
        targetVal := reflect.ValueOf(target)
        if (targetVal.Kind() == reflect.Ptr &&
                targetVal.Elem().Type() == srcVal.Type() &&
                targetVal.Elem().CanSet()) {
            targetVal.Elem().Set(srcVal)
        }
    }
}

func incrementOrUpper(values ...interface{}) {
    for _, elem := range values {
        elemValue := reflect.ValueOf(elem)
        if (elemValue.Kind() == reflect.Ptr) {
            elemValue = elemValue.Elem()
        }
        if (elemValue.CanSet()) {
            switch (elemValue.Kind()) {
                case reflect.Int:
                    elemValue.SetInt(elemValue.Int() + 1)
                case reflect.String:
                    elemValue.SetString(strings.ToUpper(elemValue.String()))
            }
            Printfln("Modified Value: %v", elemValue)
        } else {
            Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
        }
    }
}

func selectValue(data interface{}, index int) (result interface{}) {
    dataVal := reflect.ValueOf(data)
    if (dataVal.Kind() == reflect.Slice) {
        result = dataVal.Index(index).Interface()
    }
    return
}
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
    // product := Product {
    //     Name: "Kayak", Category: "Watersports", Price: 279,
    // }
    // customer := Customer { Name: "Alice", City: "New York"}
    // // 可以使用不同的类型
    // // printDetails(product, customer)
    // payment := Payment { Currency: "USD", Amount: 100.6}
    // number := 100
    // printDetails(product, customer, payment, 11, true, &number, "Alice")
    // names := []string {"Alice", "Bob", "Charlie"}
    // scores := []interface{} {99, 88, 77}
    // val := selectValue(scores, 1).(int)
    // Printfln("Selected: %v", val)
    // name := "Alice"
    // price := 279
    // city := "London"

    // incrementOrUpper(&name, &price, &city)
    // for _, val := range []interface{} { name, price, city } {
    //     Printfln("Value: %v", val)
    // }
    //
    // setAll("New String", &name, &price, &city)
    // setAll(10, &name, &price, &city)
    // for _, val := range []interface{} { name, price, city } {
    //     Printfln("Value: %v", val)
    // }

    // Converting Values

    name := "Alice"
    price := 279

    newVal, ok := convert(price, 100.00)
    Printfln("Converted %v: %v, %T", ok, newVal, newVal)
    newVal, ok = convert(name, 100.00)
    Printfln("Converted %v: %v, %T", ok, newVal, newVal)

    newVal, ok = convert(5000, int8(100))
    Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}

