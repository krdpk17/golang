package main

import (
        "fmt"
        "log"
        "net/http"
        )

type helloHandler int
func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "you have hit %s\n", r.URL.Path)
}

func main(){
        new_mux := http.NewServeMux()
        new_mux.Handle("/one", helloHandler(1))
        err_closure := http.ListenAndServe(":9999", new_mux)
        log.Fatal(err_closure)

}
