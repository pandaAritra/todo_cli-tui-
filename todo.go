package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt    time.Time
	CompleatedAt *time.Time
}

type Todos []Todo

func (todos *todos) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}





// func Todolist() {
// 	list := make([]Todo, 5)

// 	list[1] = Todo{Title: "get food", Completed: false}

// 	fmt.Println(list[1])

// }
