package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Name       string `json:"name"`
	Done       bool   `json:"done"`
	CreateDate string `json:"CreateDate"`
	DoneDate   string `json:"doneDate"`
}

var todosList Todos

func importTodos(filePath string) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	value, _ := io.ReadAll((jsonFile))

	json.Unmarshal(value, &todosList)
	jsonFile.Close()
}

func addTodo() {
	var name string

	fmt.Println("What task do you want to do? ")
	fmt.Scanln(&name)

	currentTime := time.Now()
	date := fmt.Sprintf("%d-%d-%d %d:%d", currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())

	todo := Todo{
		Name:       name,
		Done:       false,
		CreateDate: date,
		DoneDate:   "",
	}

	todosList.Todos = append(todosList.Todos, todo)
	saveTodo()
	fmt.Println("-----------------")
	showTodos()
}

func saveTodo() {
	file, _ := json.MarshalIndent(todosList, "", " ")

	_ = os.WriteFile("todos.json", file, 0644)
}

func doneTodos() {
	var index int

	fmt.Println("Which todo you did? ")
	fmt.Println("-----------------")
	showTodos()
	fmt.Println("-----------------")
	fmt.Scanln(&index)

	if index > len(todosList.Todos) {
		fmt.Println("Something is wrong")
	} else {
		todosList.Todos[index-1].Done = true
		saveTodo()
	}

}

func showTodos() {
	var done string
	fmt.Println("Yours todos:")
	for i := 0; i < len(todosList.Todos); i++ {
		if todosList.Todos[i].Done {
			done = "✔️"
		} else {
			done = "✖️"
		}
		fmt.Printf("%d: %s %s\n", i+1, todosList.Todos[i].Name, done)
	}
}

func showInfoTodos() {
	var done string
	fmt.Println("Yours todos:")
	for i := 0; i < len(todosList.Todos); i++ {
		if todosList.Todos[i].Done {
			done = "✔️"
		} else {
			done = "✖️"
		}

		fmt.Printf(`-----------------
Name: %s
Done: %s
Created date: %s
Ended date: %s
`, todosList.Todos[i].Name, done, todosList.Todos[i].CreateDate, todosList.Todos[i].DoneDate)
	}
}

func deleteTodo() {
	var index int

	fmt.Println("Which todo you want to delete?")
	fmt.Println("-----------------")
	showTodos()
	fmt.Println("-----------------")
	fmt.Scanln(&index)

	if index > len(todosList.Todos) {
		fmt.Println("Something is wrong")
	} else {
		index--
		todosList.Todos = append(todosList.Todos[:index], todosList.Todos[index+1:]...)
		saveTodo()
	}
}

func main() {
	importTodos("todos.json")
	// addTodo()
	// showTodos()
	// saveTodo()
	app := &cli.App{
		Name:  "todoCLI",
		Usage: "Making todo list",
		Commands: []*cli.Command{
			{
				Name:    "addTodo",
				Aliases: []string{"a"},
				Usage:   "Add task to todo",
				Action: func(ctx *cli.Context) error {
					addTodo()
					return nil
				},
			},
			{
				Name:    "doneTodo",
				Aliases: []string{"d"},
				Usage:   "Mark a task as done",
				Action: func(ctx *cli.Context) error {
					doneTodos()
					return nil
				},
			},
			{
				Name:    "infoTodo",
				Aliases: []string{"i"},
				Usage:   "Shows more informations about todos",
				Action: func(ctx *cli.Context) error {
					showInfoTodos()
					return nil
				},
			},
			{
				Name:    "deleteTodo",
				Aliases: []string{"D"},
				Usage:   "Delete selected todo",
				Action: func(ctx *cli.Context) error {
					deleteTodo()
					return nil
				},
			},
		},
		Action: func(*cli.Context) error {
			showTodos()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
