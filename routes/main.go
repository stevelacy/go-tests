
package main

import "log"
import "net/http"
import "github.com/gorilla/mux"

import "routes"

func main () {
  port := "3000"
  router := mux.NewRouter()

  router.HandleFunc("/", routes.Index).Methods("GET")
  router.HandleFunc("/home", routes.Home).Methods("GET")

  http.Handle("/", router)

  log.Print("starting on port ", port)
  http.ListenAndServe(":" + port, nil)

  fmt.Print(routes.Config)
}
