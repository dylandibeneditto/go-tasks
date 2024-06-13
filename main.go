package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title   string   `json:"title"`
	Commits []string `json:"commits"`
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
		color.Set(color.FgHiRed)
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
	color.Cyan("Your tasks:")
	fmt.Println("---")
	for _, task := range tasks.Tasks {
		color.Set(color.FgHiYellow)
		fmt.Print(" " + strconv.Itoa(len(task.Commits)))
		color.Unset()
		fmt.Print(" - ")
		color.Set(color.FgHiMagenta)
		fmt.Println(task.Title)
		if len(task.Commits) == 0 {
			color.HiBlack("%s", " > N/A")
		} else {
			color.HiBlack(" > last: %s", task.Commits[0])
		}
		color.Unset()
		fmt.Println("---")
	}
}

func addTask(tasks *Tasks, title string) {
	tasks.Tasks = append(tasks.Tasks, Task{Title: title, Commits: []string{}})
	color.Green("Task '%s' added.", title)
}

func removeTask(tasks *Tasks, title string) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			color.Red("Task '%s' removed.", title)
			return
		}
	}
	color.Red("Task '%s' not found.", title)
}

func renameTask(tasks *Tasks, reader *bufio.Reader, title string) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			color.Set(color.FgHiBlack)
			fmt.Print("New task title: ")
			color.Unset()
			newTitleRaw, _ := reader.ReadString('\n')
			newTitle := strings.TrimSpace(newTitleRaw)
			tasks.Tasks[i].Title = newTitle
			color.Yellow("Task '%s' changed to '%d'.", title, newTitle)
			return
		}
	}
	color.Red("Task '%s' not found.", title)
}

func commitTask(tasks *Tasks, reader *bufio.Reader, title string) {
	for i, task := range tasks.Tasks {
		if task.Title == title {
			color.Set(color.FgHiBlack)
			fmt.Printf("(%s) - commit: ", title)
			color.Unset()
			newCommitRaw, _ := reader.ReadString('\n')
			newCommit := strings.TrimSpace(newCommitRaw)
			tasks.Tasks[i].Commits = append([]string{newCommit}, tasks.Tasks[i].Commits...)
			color.Green("Commit '%s' added to task '%d'", newCommit, title)
		}
	}
	color.Red("Task '%s' not found.", title)
}

func getTitle(reader *bufio.Reader) string {
	color.Set(color.FgHiBlack)
	fmt.Print("Enter task title: ")
	color.Unset()
	title, _ := reader.ReadString('\n')
	return strings.TrimSpace(title)
}

func main() {
	filename := "/Users/dylandibeneditto/Desktop/new/go-todo/items.json"

	tasks, err := loadTasks(filename)
	if err != nil {
		color.Red("Error loading tasks: %v", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	printTasks(tasks)
	color.Set(color.FgHiBlack)
	fmt.Print("\n(add, remove, rename, commit): ")
	color.Unset()
	actionString, _ := reader.ReadString('\n')
	action := strings.TrimSpace(actionString)

	switch action {
	case "add":
		title := getTitle(reader)
		if title == "" {
			color.Red("Title must be provided to add a task.")
			return
		}
		addTask(&tasks, title)
	case "remove":
		title := getTitle(reader)
		if title == "" {
			color.Red("Title must be provided to remove a task.")
			return
		}
		removeTask(&tasks, title)
	case "rename":
		title := getTitle(reader)
		if title == "" {
			color.Red("Title must be provided to change a task.")
			return
		}
		renameTask(&tasks, reader, title)
	case "commit":
		title := getTitle(reader)
		if title == "" {
			color.Red("Title must be provided to commit to a task.")
		}
		commitTask(&tasks, reader, title)
	default:
		color.Red("Invalid action: %s", action)
		return
	}

	err = saveTasks(filename, tasks)
	if err != nil {
		color.Red("Error saving tasks: %v", err)
	}
}
