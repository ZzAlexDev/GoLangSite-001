package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string  //User Name
	Age                  uint16  //User Age
	Money                int16   //User Balance
	Avg_grades, Happynes float64 //User raiting, user happynes
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User Name is: %s. He is %d and he has money egual: %d $", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	//u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -25, 4.2, 0.8}
	//bob.setNewName("Alllllexxx")
	// bob.name = "Alexxx"
	//fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, "<b>Main Text</b>")
	//fmt.Fprintf(w, `<h1>Main Text</h1>`)
	//fmt.Fprintf(w, "<b>Main Text</b>")

	tmp, _ := template.ParseFiles("templates/home_page.html")
	tmp.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page!")
}

func HandleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8800", nil)
}

func main() {
	//	var bob user ...
	//  bob := User{name: "Bob", age: "25", money: -50, avg_grades: 4.2, happynes: 0.8}
	HandleRequest()
}
