package main

import (
  "fmt"
  "net/http"
  "github.com/codegangsta/martini"
)

func main() {
  m := martini.Classic()

  // middleware
  // m.Use(Auth)

  m.Get("/", func() string {
    return "route: /"
  })

  // define per-route middleware
  m.Get("/auth", Auth, func() string {
    return "Valid authentication"
  })


  m.Get("/:user", func(params martini.Params) string {
    return "user:" + params["user"]
  })


  m.Run()
  fmt.Println("Starting server")
}


func Auth(res http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.RawQuery)

  if req.URL.RawQuery != "securecat" {
    http.Error(res, "Invalid authentication", 401)
  }
}
