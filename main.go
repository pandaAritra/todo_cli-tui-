package main

import (
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
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

	todos.print()

	// fmt.Printf("%T \n", strconv.Itoa(56))
}
