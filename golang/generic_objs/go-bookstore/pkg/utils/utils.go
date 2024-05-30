package utils

import (
	"encoding/json"
	// "go-bookstore/pkg/models"
	"io"
	"net/http"
	// "reflect"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// func UpdateBook(updateBook *models.Book, details *models.Book) {
// 	updateBookType := reflect.TypeOf(updateBook)
// 	updateBookValue := reflect.TypeOf(updateBook)

// 	detailsType := reflect.TypeOf(details)
// 	detailsValue := reflect.ValueOf(details)

// 	for i := 0; i < updateBookType.NumField(); i++ {
// 		updateBookValue := updateBookValue.Field(i)
// TODO 需要类型检测才能用 不容易实现

// 	}
// }
