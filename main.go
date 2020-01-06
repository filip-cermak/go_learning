package main

import ( 
	"fmt"
	"time"
)

type Entry interface {
	Title() string
}

type Book struct {
	Name	string
	Author 	string
	Published time.Time
}

func (b Book) Title() string {
	return fmt.Sprintf("%s by (%s)", b.Name, b.Author, b.Published.Format("Ja\n 2006"))
}

type Movie struct {
	Name		string
	Director	string
	Year		int
}

func (m Movie) Title() string {
	return fmt.Sprintf("%s (%d)", m.Name, m.Year)
}

func Display(e Entry) string {
	return e.Title()
}

func main() {
	langs := []string{"a", "b", "c"}
	b:= Book{Name: "Filip", Author: "David", Published: time.Date(2001, time.May, 22, 0, 0, 0, 0, time.UTC)}
	m:= Movie{Name: "The Godfather", Director: "Franceis", Year: 1972}

	fmt.Println(Display(b))
	fmt.Println(Display(m))
}


