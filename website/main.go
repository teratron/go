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

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, &user)
	if err != nil {
		fmt.Println(err)
	}
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Contacts")
	if err != nil {
		fmt.Println(err)
	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handleRequest()
}
