package data

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{
	{Id: 1, Title: "A", Finished: false},
	{Id: 2, Title: "B", Finished: false},
	{Id: 3, Title: "C", Finished: false},
	{Id: 4, Title: "D", Finished: false},
	{Id: 5, Title: "E", Finished: false},
}

func findBook(id int) (int, *Book) {
	index := -1
	var book *Book

	for i, b := range books {
		if b.Id == id {
			index = i
			book = b
		}
	}
	return index, book
}

func finishBook(id int) {
	i, book := findBook(id)
	if i > 0 {
		return
	}
	book.Finished = true
	books[i] = book
	fmt.Printf("Finished book: %s\n.", book.Title)
}

func ReadBook(id int, wg *sync.WaitGroup) {
	finishBook(id)

	delay := rand.Intn(500)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	wg.Done()
}
