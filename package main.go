package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Book struct with required fields
type Book struct {
	Title    string
	Author   string
	Year     int
	Borrowed bool
}

// Slice to store all books
var books []Book

var reader = bufio.NewReader(os.Stdin)

// Helper: Read input from user
func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// 4. Function to add a book
func addBook() {
	title := readInput("Enter book title: ")
	author := readInput("Enter author name: ")
	yearStr := readInput("Enter publication year: ")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("Invalid year. Book not added.")
		return
	}

	book := Book{
		Title:    title,
		Author:   author,
		Year:     year,
		Borrowed: false,
	}
	books = append(books, book)
	fmt.Println("Book added successfully!\n")
}

// 5. Function to list all books
func listBooks() {
	if len(books) == 0 {
		fmt.Println("No books in the library.\n")
		return
	}
	fmt.Println("Library Books:")
	for i, b := range books {
		status := "Available"
		if b.Borrowed {
			status = "Borrowed"
		}
		fmt.Printf("%d. \"%s\" by %s (%d) - %s\n", i+1, b.Title, b.Author, b.Year, status)
	}
	fmt.Println()
}

// 6. Function to borrow a book
func borrowBook() {
	if len(books) == 0 {
		fmt.Println("No books to borrow.\n")
		return
	}
	listBooks()
	bookNumStr := readInput("Enter the number of the book to borrow: ")
	bookNum, err := strconv.Atoi(bookNumStr)
	if err != nil || bookNum < 1 || bookNum > len(books) {
		fmt.Println("Invalid book number.\n")
		return
	}
	if books[bookNum-1].Borrowed {
		fmt.Println("Sorry, that book is already borrowed.\n")
		return
	}
	books[bookNum-1].Borrowed = true
	fmt.Printf("You have borrowed \"%s\".\n\n", books[bookNum-1].Title)
}

// 7. Function to return a book
func returnBook() {
	if len(books) == 0 {
		fmt.Println("No books to return.\n")
		return
	}
	listBooks()
	bookNumStr := readInput("Enter the number of the book to return: ")
	bookNum, err := strconv.Atoi(bookNumStr)
	if err != nil || bookNum < 1 || bookNum > len(books) {
		fmt.Println("Invalid book number.\n")
		return
	}
	if !books[bookNum-1].Borrowed {
		fmt.Println("That book was not borrowed.\n")
		return
	}
	books[bookNum-1].Borrowed = false
	fmt.Printf("You have returned \"%s\".\n\n", books[bookNum-1].Title)
}

// 8. Main menu loop
func main() {
	for {
		fmt.Println("Mini Library Manager")
		fmt.Println("1. Add Book")
		fmt.Println("2. List Books")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Exit")
		choice := readInput("Choose an option (1-5): ")

		switch choice {
		case "1":
			addBook()
		case "2":
			listBooks()
		case "3":
			borrowBook()
		case "4":
			returnBook()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.\n")
		}
	}
}