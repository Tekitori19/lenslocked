package main

import (
	"html/template"
	"os"
)

type User struct{
	Name string
	Age int
	Meta UserMeta
}

type UserMeta struct{
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
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
		Age: 19,
		Meta:  UserMeta{
			Visits: 11,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}