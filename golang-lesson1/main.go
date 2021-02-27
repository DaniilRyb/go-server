package main

import (
	"fmt"
	"net/http"
)

func homePage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Page is open!!!")
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!!!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
