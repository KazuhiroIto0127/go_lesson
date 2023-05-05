package main

import "fmt"

type Attr struct {
	Name string
	Age int
}

type Teacher struct{
	Attr
	Subject string
}

type Student struct{
	Attr
	Score int
}

func main(){
	teacher := Teacher {
		Attr: Attr{
			Name: "John Taro",
			Age: 20,
		},
		Subject: "Math",
	}

	student := Student {
		Attr: Attr{
			Name: "Tanaka Dayo",
			Age: 15,
		},
		Score: 45,
	}

	fmt.Println(teacher.Name, teacher.Subject)
	fmt.Println(student.Name, student.Score)
}
