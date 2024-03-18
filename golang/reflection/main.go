package main

import (
	"reflect"
	"strings"
	// "fmt"
)


func inspectFuncType(f interface{}) {
    funcType := reflect.TypeOf(f)
    if (funcType.Kind() == reflect.Func) {
        Printfln("Function parameters: %v", funcType.NumIn())
        for i := 0; i < funcType.NumIn(); i++ {
            paramType := funcType.In(i)
            if (i < funcType.NumIn() - 1) {
                Printfln("Parameter #%v, Type: %v", i, paramType)
            } else {
                Printfln("Parameter #%v, Type: %v, Variadic: %v")
            }
        }
        Printfln("Function results: %v", funcType.NumOut())
        for i := 0; i < funcType.NumOut(); i++ {
            resultType := funcType.Out(i)
            Printfln("Result #%v, Type: %v", i, resultType)
        }
    }
}

func inspectStructs(structs ...interface{}) {
    for _, s := range structs {
        structType := reflect.TypeOf(s)
        if (structType.Kind() == reflect.Struct) {
            inspectStructType(structType)
        }
    }
}

func inspectStructType(structType reflect.Type) {
    Printfln("--- Struct Type: %v", structType)
    for i := 0;i < structType.NumField(); i++ {
        field := structType.Field(i)
        Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
            field.Index, field.Name, field.Type, field.PkgPath == "")
    }
    Printfln("--- End struct Type: %v", structType)
}

func setMap(m interface{}, key interface{}, val interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    valValue := reflect.ValueOf(val)
    if (mapValue.Kind() == reflect.Map &&
        mapValue.Type().Key() == keyValue.Type() &&
        mapValue.Type().Elem() == valValue.Type()) {
        mapValue.SetMapIndex(keyValue, valValue)
    } else {
        Printfln("Not a map or mismatched types")
    }
}

func removeFromMap(m interface{}, key interface{}) {
    mapValue := reflect.ValueOf(m)
    keyValue := reflect.ValueOf(key)
    if (mapValue.Kind() == reflect.Map &&
        mapValue.Type().Key() == keyValue.Type()) {
        mapValue.SetMapIndex(keyValue, reflect.Value{})
    }
}

func printMapContents(m interface{}) {
    mapValue := reflect.ValueOf(m)
    if (mapValue.Kind() == reflect.Map) {
        for _, keyVal := range mapValue.MapKeys() {
            reflectVal := mapValue.MapIndex(keyVal)
            Printfln("Map Key: %v, Value: %v", keyVal, reflectVal)
        }
    } else {
        Printfln("Not a map")
    }
}

func describeMap(m interface{}) {
    mapType := reflect.TypeOf(m)
    if (mapType.Kind() == reflect.Map) {
        Printfln("Key type: %v, val type: %v,", mapType.Key(), mapType.Elem())
    } else {
        Printfln("Not a map")
    }
}

func findAndSplit(slice interface{}, target interface{}) interface{} {
    sliceVal := reflect.ValueOf(slice)
    targetType := reflect.TypeOf(target)
    if (sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType) {
        for i := 0;i < sliceVal.Len(); i++ {
            if sliceVal.Index(i).Interface() == target {
                return sliceVal.Slice(0, i + 1)
            }
        }
    }
    return slice
}

func enumerateStrings(arrayOrSlice interface{}) {
    arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
    if (arrayOrSliceVal.Kind() == reflect.Array || 
            arrayOrSliceVal.Kind() == reflect.Slice) && 
            arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
        for i := 0; i < arrayOrSliceVal.Len(); i++ {
            Printfln("Element: %v, Value: %v", i, arrayOrSliceVal.Index(i))
        }
    }

}

func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
    arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
    replacementVal := reflect.ValueOf(replacement)
    if (arrayOrSliceVal.Kind() == reflect.Slice) {
        elemVal := arrayOrSliceVal.Index(index)
        if (elemVal.CanSet()) {
            elemVal.Set(replacementVal)
        }
    } else if (arrayOrSliceVal.Kind() == reflect.Ptr &&
        arrayOrSliceVal.Elem().Kind() == reflect.Array &&
        arrayOrSliceVal.Elem().CanSet()) {
            arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
    }
}

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
    elemVal := reflect.ValueOf(val)
    if (elemVal.Type() == stringPtrType) {
        upperStr := strings.ToUpper(elemVal.Elem().String())
        if (elemVal.Elem().CanSet()) {
            elemVal.Elem().SetString(upperStr)
        }
    }
}

func checkElemType(val interface{}, arrOrSlice interface{}) bool {
    elemType := reflect.TypeOf(val)
    arrOrSliceType := reflect.TypeOf(arrOrSlice)
    return (arrOrSliceType.Kind() == reflect.Array || arrOrSliceType.Kind() == reflect.Slice) && arrOrSliceType.Elem() == elemType

}

func createPointerType(t reflect.Type) reflect.Type {
    return reflect.PointerTo(t)
}

func followPointerType(t reflect.Type) reflect.Type {
    if t.Kind() == reflect.Ptr {
        return t.Elem()
    }
    return t
}

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
    city := "London"
    hobby := "Running"

    slice := []string {name, city, hobby}
    array := [3]string {name, city, hobby}
    // price := 279
    //
    // newVal, ok := convert(price, 100.00)
    // Printfln("Converted %v: %v, %T", ok, newVal, newVal)
    // newVal, ok = convert(name, 100.00)
    // Printfln("Converted %v: %v, %T", ok, newVal, newVal)
    //
    // newVal, ok = convert(5000, int8(100))
    // Printfln("Converted %v: %v, %T", ok, newVal, newVal)
    t := reflect.TypeOf(name)
    Printfln("Original Type: %v", t)
    pt := createPointerType(t)
    Printfln("Pointer Type: %v", pt)
    Printfln("Follow pointer type: %v", followPointerType(pt))
    transformString(&name)
    Printfln("Follow pointer value: %v", name)

    Printfln("Slice (string): %v", checkElemType("testString", slice))
    Printfln("Array (string): %v", checkElemType("testString", array))
    Printfln("Array (int): %v", checkElemType(10, array))

    Printfln("Original slice: %v", slice)
    newCity := "Paris"
    setValue(slice, 1, newCity)
    Printfln("Modified slice: %v", slice)

    Printfln("Original slice: %v", array)
    newCity = "Rome"
    setValue(&array, 1, newCity)
    Printfln("Modified slice: %v", array)

    enumerateStrings(slice)
    enumerateStrings(array)

    pricesMap := map[string]float64 {
        "Kayak": 279, "Lifejacket": 100, "Soccer Ball": 20.50,
    }
    describeMap(pricesMap)

    printMapContents(pricesMap)

    inspectStructs( Purchase{} )

    inspectFuncType(Find)
}

