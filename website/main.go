package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     uint8
	Money   float64
	Hobbies []string
}

func (u *User) getInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: $%.2f", u.Name, u.Age, u.Money)
}

func (u *User) setName(name string) {
	u.Name = name
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:    "Bob",
		Age:     23,
		Money:   2100,
		Hobbies: []string{"Skate", "Dance", "Football"},
	}
	user.setName("Alex")
	//_, err := fmt.Fprintf(w, bob.getInfo())
	//_, err := fmt.Fprintf(w, "<b>%s</b>", "Bold text")
	//_, err := fmt.Fprintf(w, `<b>%s</b>
	//					<h1>Header</h1>`, "Bold text")

	tmpl, _ := template.ParseFiles("templates/home.html")

	_ = tmpl.Execute(w, &user)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)

	_ = http.ListenAndServe(":8083", nil)
}

func main() {
	handleRequest()
}
