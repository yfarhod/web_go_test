package main

import ("fmt"
"net/http")

type User struct {
  name string
  age uint16
  money int16
  avg_grades, happiness float64
}

func home_page(page http.ResponseWriter, r *http.Request) {
  user1 := User{"Bob", 25, -10, 4.2, 0.8}
  user1.name = "Alex"
  fmt.Fprintf(page, "User name is: " + user1.name)
}

func contacts_page(page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(page, "Contacts page!")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleRequest()

}
