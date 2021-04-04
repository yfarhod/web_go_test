package main

import ("fmt"
"net/http"
"html/template")

type User struct {
  name string
  age uint16
  money int16
  avg_grades, happiness float64
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is %s. He is %d years old and he has %d money.",
  u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
  user1 := User{"Bob", 25, -10, 4.2, 0.8}
  user1.setNewName("Alex")
  // fmt.Fprintf(w, user1.getAllInfo())
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, user1)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleRequest()

}
