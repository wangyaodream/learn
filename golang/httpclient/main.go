package main

import (
	//"io"
	// "encoding/json"
	"io"
	"net/http"
	"os"
	"time"
    // "strings"
    // "net/http/cookiejar"
)

func main() {
    go http.ListenAndServe(":5000", nil)
    time.Sleep(time.Second)

    // formData := map[string][]string {
    //     "name": {"Kayak"},
    //     "category": {"watersports"},
    //     "price": {"279"},
    // }

    // response, err := http.Get("http://localhost:5000/json")
    // response, err := http.PostForm("http://localhost:5000/echo", formData)
    // if (err == nil && response.StatusCode == http.StatusOK) {
    //     // defer response.Body.Close()
    //     // data := []Product {}
    //     // err = json.NewDecoder(response.Body).Decode(&data)
    //     // if (err == nil) {
    //     //     for _, p := range data {
    //     //         Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
    //     //     }
    //     // } else {
    //     //     Printfln("Decode error: %v", err.Error())
    //     // }
    //     io.Copy(os.Stdout, response.Body)
    //     defer response.Body.Close()
    // } else {
    //     Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
    // }


    // var builder strings.Builder
    // err := json.NewEncoder(&builder).Encode(Products[0])
    // if (err == nil) {
    //     // response, err := http.Post("http://localhost:5000/echo", "application/json",            strings.NewReader(builder.String()))
    //     reqURL, err := url.Parse("http://localhost:5000/echo")
    //     if (err == nil) {
    //         req := http.Request {
    //             Method: http.MethodPost,
    //             URL: reqURL,
    //             Header: map[string][]string {
    //                 "Content-Type": {"application/json"},
    //             },
    //             Body: io.NopCloser(strings.NewReader(builder.String())),
    //         }
    //         response, err := http.DefaultClient.Do(&req)
    //         if (err == nil && response.StatusCode == http.StatusOK) {
    //             io.Copy(os.Stdout, response.Body)
    //             defer response.Body.Close()
    //         } else {
    //             Printfln("Request Error: %v", err.Error())
    //         }
    //     } else {
    //         Printfln("Parse Error: %v", err.Error())
    //     }
    //     // if (err == nil && response.StatusCode == http.StatusOK) {
    //     //     io.Copy(os.Stdout, response.Body)
    //     //     defer response.Body.Close()
    //     // } else {
    //     //     Printfln("Error: %v", err.Error())
    //     // }
    // } else {
    //     Printfln("Encode Error: %v", err.Error())
    // }

    // Working with cookies
    // jar, err := cookiejar.New(nil)
    // if (err == nil) {
    //     http.DefaultClient.Jar = jar
    // }
    //
    // for i := 0; i < 3; i++ {
    //     req, err := http.NewRequest(http.MethodGet,
    //     "http://localhost:5000/cookie", nil)
    //     if (err == nil) {
    //         response, err := http.DefaultClient.Do(req)
    //         if (err == nil && response.StatusCode == http.StatusOK) {
    //             io.Copy(os.Stdout, response.Body)
    //             defer response.Body.Close()
    //         } else {
    //             Printfln("Request Error: %v", err.Error())
    //         }
    //     } else {
    //         Printfln("Request init error: %v", err.Error())
    //     }
    // }

    req, err := http.NewRequest(http.MethodGet,
    "http://localhost:5000/redirect1", nil)
    if (err == nil) {
        var response *http.Response
        response, err = http.DefaultClient.Do(req)
        if (err == nil) {
            io.Copy(os.Stdout, response.Body)
        } else {
            Printfln("Request Error: %v", err.Error())
        }
    } else {
        Printfln("Error: %v", err.Error())
    }
}
