package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u *User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user1 := User{Name: "dasom", Email: "dasom@gmail.com", Age: 33}
	user2 := User{Name: "kelly", Email: "kelly@gmail.com", Age: 11}
	users := []User{user1, user2}
	// struct key값이 대문자로 시작해야 템플릿에 값 주입 가능
	// tmpl, err := template.New("tmpl").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}\n")
	// if err != nil {
	// 	panic(err)
	// }
	// tmpl.Execute(os.Stdout, user1)
	// tmpl.Execute(os.Stdout, user2)

	// 템플릿 파일로 관리
	tmplFile, err := template.New("tmpl").ParseFiles("file/tmpl1.tmpl", "file/index.html")
	if err != nil {
		panic(err)
	}
	tmplFile.ExecuteTemplate(os.Stdout, "index.html", users)
	// tmplFile.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
}
