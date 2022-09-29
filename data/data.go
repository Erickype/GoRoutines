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
	{Id: 6, Title: "F", Finished: false},
	{Id: 7, Title: "G", Finished: false},
	{Id: 8, Title: "H", Finished: false},
	{Id: 9, Title: "I", Finished: false},
	{Id: 10, Title: "J", Finished: false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.Id == (id + 1) {
			index = i
			book = b
		}
	}
	m.RUnlock()

	return index, book
}

func finishBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Finished book: %s.\n", book.Title)
}

func ReadBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	finishBook(id, m)

	delay := rand.Intn(500)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	wg.Done()
}
