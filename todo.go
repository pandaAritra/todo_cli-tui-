package main

import (
	"errors"
	"fmt"
	"time"
)

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		completedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) listAll() {
	for i, todo := range *todos {
		fmt.Printf("%v :is -> %+v ", i, todo.Title)
		if todo.Completed == true {
			fmt.Printf("completed at %v", todo.completedAt.Format(time.RFC822))
		}
		fmt.Println()

	}
}

func (todos *Todos) markeAsDone(title string) {
	t := *todos
	for i := range t {
		if t[i].Title == title {
			t[i].Completed = true
			crrTime := time.Now()
			t[i].completedAt = &crrTime
		}
	}
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		return err
	}
	return nil
}

func (todos *Todos) delete(title string) {
	deleted := false
	for i, todo := range *todos {
		if todo.Title == title {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		fmt.Printf("%v was not found \n", title)
	}
}
