package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type bookResource struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {
	// GET
	fmt.Println("Making GET Call.")
	resp, err := http.Get("http://localhost:8080/api/books")
	if err != nil {
		fmt.Println("Error while making GET call.", err)
		return
	}

	fmt.Printf("%+v\n\n", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var books []bookResource
	json.Unmarshal(body, &books)
	fmt.Println(books)
	fmt.Println("\n")

	// POST
	payload, _ := json.Marshal(bookResource{
		Title: "New Book",
		Link:  "http://new-book.com",
	})

	fmt.Println("Making POST call.")
	resp, err = http.Post(
		"http://localhost:8080/api/books/",
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n\n", resp)

	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var book bookResource
	json.Unmarshal(body, &book)

	fmt.Println(book)

	fmt.Println("\n")
}
