package main

import (
	"fmt"
	"sort"
)

type database struct {
	host string
	port int
}
type Db interface {
	Get(id string) (string, error)
}

func (r *database) Get(id string) (string, error) {
	fmt.Println("get method called")
	return id, nil

}

type Repository struct {
	db Db
}

func NewRepo(db Db) *Repository {
	return &Repository{db}
}
func (r *Repository) Get(id string) (string, error) {
	return r.db.Get(id)
}

func main() {
	num1 := "112"
	digit1 = int(num1[0] - '0')
	a := []int{1, 2}
	sort.Ints(a)
	db := &database{"satyam", 5505}
	repo := NewRepo(db)
	data, err := repo.Get("11")
	if err != nil {
		fmt.Println("there is some error")
	} else {
		fmt.Println(data)
	}

}
