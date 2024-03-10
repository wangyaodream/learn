package main

import (
	//"io"
	// "encoding/json"
	"io"
	"net/http"
	"os"

	//"os"
	"time"
)

func main() {
    go http.ListenAndServe(":5000", nil)
    time.Sleep(time.Second)

    formData := map[string][]string {
        "name": {"Kayak"},
        "category": {"watersports"},
        "price": {"279"},
    }

    // response, err := http.Get("http://localhost:5000/json")
    response, err := http.PostForm("http://localhost:5000/echo", formData)
    if (err == nil && response.StatusCode == http.StatusOK) {
        // defer response.Body.Close()
        // data := []Product {}
        // err = json.NewDecoder(response.Body).Decode(&data)
        // if (err == nil) {
        //     for _, p := range data {
        //         Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
        //     }
        // } else {
        //     Printfln("Decode error: %v", err.Error())
        // }
        io.Copy(os.Stdout, response.Body)
        defer response.Body.Close()
    } else {
        Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
    }
}
