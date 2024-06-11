package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Importance int    `json:"importance"`
}

func loadTasks(filename string) (Tasks, error) {
	var tasks Tasks
	jsonData, jsonErr := os.Open(filename)
	if jsonErr != nil {
		return tasks, jsonErr
	}
	defer jsonData.Close()

	byteValue, _ := io.ReadAll(jsonData)
	err := json.Unmarshal(byteValue, &tasks)
	if err != nil {
		return tasks, fmt.Errorf("error while decoding the data: %v", err)
	}

	return tasks, nil
}

func saveTasks(filename string, tasks Tasks) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	return err
}

func printTasks(tasks Tasks) {
	for _, task := range tasks.Tasks {
		fmt.Printf("Task Title: %s, Description: %s, Importance: %d\n", task.Title, task.Desc, task.Importance)
	}
}

func addTask(tasks *Tasks, title, desc string, importance int) {
	tasks.Tasks = append(tasks.Tasks, Task{Title: title, Desc: desc, Importance: importance})
}

func removeTask(tasks *Tasks, title string) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			break
		}
	}
}

func changeTask(tasks *Tasks, title, newDesc string, newImportance int) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			tasks.Tasks[i].Desc = newDesc
			tasks.Tasks[i].Importance = newImportance
			break
		}
	}
}

func main() {
	action := flag.String("action", "", "Action to perform: add, remove, change")
	title := flag.String("title", "", "Title of the task")
	desc := flag.String("desc", "", "Description of the task")
	importance := flag.Int("importance", 0, "Importance of the task")
	flag.Parse()

	filename := "/Users/dylandibeneditto/Desktop/new/go-todo/items.json"

	tasks, err := loadTasks(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch *action {
	case "add":
		if *title == "" || *desc == "" {
			fmt.Println("Title and description must be provided to add a task.")
			return
		}
		addTask(&tasks, *title, *desc, *importance)
		err = saveTasks(filename, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case "remove":
		if *title == "" {
			fmt.Println("Title must be provided to remove a task.")
			return
		}
		removeTask(&tasks, *title)
		err = saveTasks(filename, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case "change":
		if *title == "" || *desc == "" {
			fmt.Println("Title and new description must be provided to change a task.")
			return
		}
		changeTask(&tasks, *title, *desc, *importance)
		err = saveTasks(filename, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	default:
		fmt.Println("Tasks:")
		printTasks(tasks)
	}
}
