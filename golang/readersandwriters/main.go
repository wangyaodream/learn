package main

import (
    "io"
//    "strings"
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
    // // strings.NewReader返回值是一个指针
    // r := strings.NewReader("Thisisalongstring")
    // var builder strings.Builder
    // processData(r, &builder)
    // Printfln("String builder contents: %s", builder.String())

    pipReader, pipWriter := io.Pipe()
    go func() {
        GenerateData(pipWriter)
        pipWriter.Close()
    }()
    ConsumeData(pipReader)
}
