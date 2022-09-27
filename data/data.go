package data

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{}

func FindBook(id int) (int, *Book) {
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
