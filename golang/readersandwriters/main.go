package main

import (
	// "io"
	"encoding/json"
	"fmt"
	"strings"
)

// func processData(reader io.Reader, writer io.Writer) {
//     b := make([]byte, 2)
//     for {
//         count, err := reader.Read(b)
//         if (count > 0) {
//             writer.Write(b[0:count])
//             fmt.Println("count:", count)
//             Printfln("Read %v bytes: %v", count, string(b[0:count]))
//         }
//
//         if err == io.EOF {
//             break
//         }
//     }
// }

func main() {

	names := []string{"Kayak", "Lifejacket", "Soccer ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	fmt.Print(writer.String())
	// // strings.NewReader返回值是一个指针
	// r := strings.NewReader("Thisisalongstring")
	// var builder strings.Builder
	// processData(r, &builder)
	// Printfln("String builder contents: %s", builder.String())

	// pipReader, pipWriter := io.Pipe()
	// go func() {
	//     GenerateData(pipWriter)
	//     pipWriter.Close()
	// }()
	// ConsumeData(pipReader)
	// var b bool = true
	// var str string = "Hello"
	// var fval float64 = 99.99
	// var ival int = 200
	// var pointer *int = &ival
	//
	// var writer strings.Builder
	// encoder := json.NewEncoder(&writer)
	//
	// for _, val := range []interface{} {b, str, fval, ival, pointer} {
	//     encoder.Encode(val)
	// }
	//
	// fmt.Print(writer.String())

}
