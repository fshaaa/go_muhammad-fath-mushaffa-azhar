package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type Book struct {
	id       string
	title    string
	price    int
	category string
}

var books []Book

func MenuBook() {
	fmt.Println("==BOOK MANAGEMENT==")
	fmt.Println("1. Get All Book")
	fmt.Println("2. Create a Book")
	fmt.Println("3. Update a Book")
	fmt.Println("4. Delete a Book")
	fmt.Println("5. Exit")
	fmt.Print("Choose your menu: ")
}

func GetAllBooks() {
	fmt.Println("All books")
	for _, book := range books {
		fmt.Println("===")
		fmt.Println("ID: " + book.id)
		fmt.Println("Title: " + book.title)
		fmt.Println("Price:", book.price)
		fmt.Println("Category: " + book.category)
		fmt.Println("===")
	}
}

func CreateBook() {
	var book Book
	fmt.Println("Enter Title: ")
	fmt.Scanln(&book.title)
	fmt.Println("Enter Price: ")
	fmt.Scanln(&book.price)
	fmt.Println("Enter Category: ")
	fmt.Scanln(&book.category)

	uuid := uuid.New()
	book.id = uuid.String()

	books = append(books, book)
	fmt.Println("Book Added!")
}

func UpdateBook() {
	var book Book
	fmt.Println("Enter ID: ")
	fmt.Scanln(&book.id)
	fmt.Println("Enter Title: ")
	fmt.Scanln(&book.title)
	fmt.Println("Enter Price: ")
	fmt.Scanln(&book.price)
	fmt.Println("Enter Category: ")
	fmt.Scanln(&book.category)

	for i, b := range books {
		if book.id == b.id {
			books = append(books[:i], books[i+1:]...)
			books = append(books, book)

			fmt.Println("The ID: " + b.id)
			fmt.Println("Book Updated!")
			break
		}
	}
}

func DeleteBook() {
	var id string
	fmt.Println("Enter ID: ")
	fmt.Scanln(&id)
	for i, book := range books {
		if book.id == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Println("Book Deleted!")
		}
	}
}

func main() {
	for {
		MenuBook()
		var input int
		fmt.Scanln(&input)
		if input == 1 {
			GetAllBooks()
		} else if input == 2 {
			CreateBook()
		} else if input == 3 {
			UpdateBook()
		} else if input == 4 {
			DeleteBook()
		} else if input == 5 {
			fmt.Println("Bye...")
			os.Exit(1)
		} else {
			fmt.Println("Your input is invalid")
		}
	}
}
