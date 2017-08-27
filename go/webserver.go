// go build webserver.go
// ./webserver 1337

package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Halo Bali");
}

func main() {
  fmt.Println("serving http at port "+string(os.Args[1]))

  http.HandleFunc("/", handler)
  http.ListenAndServe(":"+string(os.Args[1]),nil)
}
