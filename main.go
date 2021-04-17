package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name   string
	email  string
	age    uint16
	wallet float64
}

func (u User) get_user() string {
	return fmt.Sprintf("User name: %s Age: %d Email: %s Wallet: $%f", u.name, u.age, u.email, u.wallet)
}

func (u *User) set_name(name string) {
	u.name = name
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go serve is runing")
}

func about_page(w http.ResponseWriter, r *http.Request) {
	john := User{name: "John", email: "john@example.com", age: 32, wallet: 354.22}
	bob := User{"Bob", "bob@example.com", 28, 533.12}

	bob.set_name("Alex")

	fmt.Fprintf(w, "About page")
	fmt.Fprintf(w, john.get_user())
	fmt.Fprintf(w, bob.get_user())
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":4443", nil)
}

func main() {
	handleRequest()
}
