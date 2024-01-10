package pkgregister

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func packageRegHander(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // 提交过来的包数据
        p := pkgData{}
        // 包注册响应
        d := pkgRegisterResult{}

        data, err := io.ReadAll(r.Body) 
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        err = json.Unmarshal(data, &p)
        if err != nil || len(p.Name) == 0 || len(p.Version) == 0 {
            return 
        }
        d.Id = p.Name + "-" + p.Version
        jsonData, err := json.Marshal(d)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(jsonData))
    } else {
        http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
        return
    }
}

func startTestPackageServer() *httptest.Server {
    ts := httptest.NewServer(http.HandlerFunc(packageRegHander)) 
    return ts
}

func TestRegisterPackageData(t *testing.T) {
    ts := startTestPackageServer()
    defer ts.Close()
    p := pkgData{
        Name: "mypackage",
        Version: "1.0.1",
    }
    resp, err := registerPackageData(ts.URL, p)
    if err != nil {
        t.Fatal(err)
    }
    if resp.Id != "mypackage-1.0.1" {
        t.Errorf("Expected package id to be mypackage-1.0.1 Got: %s", resp.Id)
    }
}

func TestRegisterEmptyPackageData(t *testing.T) {
    ts := startTestPackageServer()
    defer ts.Close()

    p := pkgData{}

    resp, err := registerPackageData(ts.URL, p)
    if err == nil {
        t.Fatal("Expected error to be non-nil, got nil")
    }
    if len(resp.Id) != 0 {
        t.Errorf("Expected package ID to be empty, got: %s", resp.Id)
    }
}
