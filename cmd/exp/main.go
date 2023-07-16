package main

import (
	"html/template"
	"os"
	"path/filepath"
)

type User struct {
	Name  string
	Age   int
	Meta  UserMeta
	Chats []string
	Fl    float64
	Map   map[string]string
	Haha  bool
}

type UserMeta struct {
	Visits struct {
		Name string
		Age  int
	}
}

func main() {
	path := filepath.Join("cmd", "exp", "hello.gohtml")
	t, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	// user := struct{
	// 	Name string
	// }{
	// 	Name: "Hung Ha",
	// }

	user := User{
		Name: "Dinh Dat Dinh Diem",
		Age:  19,
		Meta: UserMeta{
			Visits: struct {
				Name string
				Age  int
			}{Name: "Dinh", Age: 20},
		},
		Chats: []string{"hello", "world"},
		Fl:    3.4,
		Map:   map[string]string{"hana": "koisi", "lee": "sin"},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
