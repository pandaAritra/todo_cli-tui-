package main

import (
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	completedAt *time.Time
}

type Todos []Todo

func main() {
	var todos Todos = make([]Todo, 0) // var todos = Todos{}

	todos.add("make a todo app")
	todos.add("new function added")
	todos.add("drink tea")
	todos.markeAsDone("drink tea")
	todos.delete("new function added")
	todos.listAll()

	// fmt.Printf("%T \n", strconv.Itoa(56))
}
