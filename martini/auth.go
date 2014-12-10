package main

import (
  "fmt"
  "net/http"
)

func Auth(res http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.RawQuery)

  if req.URL.RawQuery != "securecat" {
    http.Error(res, "Invalid authentication", 401)
  }
}
