package main

import ("fmt"
"net/http")

func home_page(page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(page, "Go mu fu go go go")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleRequest()

}
