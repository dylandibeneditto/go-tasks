package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title      string `json:"title"`
	Importance int    `json:"importance"`
}

func loadTasks(filename string) (Tasks, error) {
	var tasks Tasks

	jsonData, err := os.Open(filename)
	if err != nil {
		return tasks, err
	}
	defer jsonData.Close()

	byteValue, _ := io.ReadAll(jsonData)

	err = json.Unmarshal(byteValue, &tasks)
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
		fmt.Printf("Task Title: %s, Importance: %d\n", task.Title, task.Importance)
	}
}

func addTask(tasks *Tasks, title string, importance int) {
	tasks.Tasks = append(tasks.Tasks, Task{Title: title, Importance: importance})
}

func removeTask(tasks *Tasks, title string) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			break
		}
	}
}

func changeTask(tasks *Tasks, title string, newImportance int) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			tasks.Tasks[i].Importance = newImportance
			break
		}
	}
}

func getTitle(reader *bufio.Reader) string {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	return strings.TrimSpace(title)
}

func getImportance(reader *bufio.Reader) int {
	for {
		fmt.Print("Enter task importance: ")
		importanceRead, _ := reader.ReadString('\n')
		importanceStr := strings.TrimSpace(importanceRead)
		importance, err := strconv.Atoi(importanceStr)
		if err == nil {
			return importance
		}
		fmt.Println("Invalid importance value. Please enter a valid integer.")
	}
}

func main() {

	filename := "/Users/dylandibeneditto/Desktop/new/go-todo/items.json"

	tasks, err := loadTasks(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Tasks:")
	printTasks(tasks)

	actionString, _ := reader.ReadString('\n')
	action := strings.Trim(actionString, " \n")

	switch action {
	case "add":
		fmt.Println("add")
		title := getTitle(reader)
		importance := getImportance(reader)
		if title == "" {
			fmt.Println("Title must be provided to add a task.")
			return
		}
		addTask(&tasks, title, importance)
	case "remove":
		title := getTitle(reader)
		if title == "" {
			fmt.Println("Title must be provided to remove a task.")
			return
		}
		removeTask(&tasks, title)
	case "change":
		title := getTitle(reader)
		importance := getImportance(reader)
		if title == "" {
			fmt.Println("Title must be provided to change a task.")
			return
		}
		changeTask(&tasks, title, importance)
	default:
		fmt.Println("exiting after command '" + action + "' is not found")
	}

	err = saveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}
