package main

import "fmt"

type database struct {
	host string
	port int
}
type Db interface {
	Get(id string) (string, error)
}

func (r *database) Get(id string) (string, error) {
	fmt.Println("get method called")
	return "1", nil

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
	db := &database{"satyam", 5505}
	repo := NewRepo(db)
	data, err := repo.Get("1")
	if err != nil {
		fmt.Println("there is some error")
	} else {
		fmt.Println(data)
	}

}

// package main

// import "fmt"

// type Database struct {
// 	Host string
// 	Port int
// }

// type DB interface {
// 	Get(id string) (string, error)
// }

// func (r *Database) Get(id string) (string, error) {
// 	fmt.Println("get method called")
// 	return "1", nil
// }

// type Repository struct {
// 	db DB
// }

// func NewRepo(db DB) *Repository {
// 	return &Repository{db}
// }

// func (r *Repository) Get(id string) (string, error) {
// 	return r.db.Get(id)
// }

// func main() {
// 	db := &Database{"satyam", 5050}
// 	repo := NewRepo(db)
// 	data, err := repo.Get("1")
// 	if err != nil {
// 		fmt.Println("there is some error")
// 	} else {
// 		fmt.Println(data)
// 	}

// }
