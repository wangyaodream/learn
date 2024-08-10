package pkgregister

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func packageRegHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        p := pkgData{}

        d := pkgRegisterResult{}

        defer r.Body.Close()
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
        }
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(jsonData))
    } else {
        http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
        return
    }
}

func startTestPackageServer() *httptest.Server {
    ts := httptest.NewServer(http.HandlerFunc(packageRegHandler))
    return ts
}

func TestRegisterPackageData(t *testing.T) {
    ts := startTestPackageServer()
    defer ts.Close()
    p := pkgData {
        Name: "mypackage",
        Version: "0.1",
    }
    resp, err := registerPackageData(ts.URL, p)
    if err != nil {
        t.Fatal(err)
    }
    if resp.Id != "mypackage-0.1" {
        t.Errorf("Excepted package id to be mypackage-0.1, Got: %s", resp.Id)
    }

}


func TestRegisterEmptyPackageData(t *testing.T) {
    ts := startTestPackageServer()
    defer ts.Close()
    p := pkgData{}

    resp, err := registerPackageData(ts.URL, p)
    if err == nil {
        t.Fatal("Excepted error to be non-nil, got nil")
    }

    if len(resp.Id) != 0 {
        t.Errorf("Excepted package ID to be empty, got: %s", resp.Id)
    }
}
