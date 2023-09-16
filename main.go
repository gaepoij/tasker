package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

var tasks []string
var doneTasks []string

type TaskData struct {
	CurrentTasks []string `json:"currentTasks"`
	DoneTasks    []string `json:"doneTasks"`
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func addTask() {
	fmt.Print("Enter a new task:")
	task := readUserInput()
	tasks = append(tasks, task)

	fmt.Println("Task added successfully!")

}

func setDone() {
	fmt.Println("Enter the task number you want to set done:")
	value := readUserInput()

	input, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil || input <= 0 || input > len(tasks) {
		fmt.Println("Invalid task number, please try again.")
		setDone()
	}

	taskToSetDone := input - 1

	var task = tasks[taskToSetDone]
	task = strings.Join(strings.Split(fmt.Sprintf("✅ %s", task), ""), "\u0336")

	doneTasks = append(doneTasks, task)
	tasks = append(tasks[:taskToSetDone], tasks[taskToSetDone+1:]...)

	fmt.Println("Task set done successfully!")

}

func deleteTask() {
	fmt.Println("Enter the task number you want to delete:")
	value := readUserInput()

	input, err := strconv.Atoi(strings.TrimSpace(value))

	fmt.Printf("Input: %d, Error: %v\n", input, err)

	if err != nil || input <= 0 || input > len(tasks) {
		fmt.Println("Invalid task number, please try again.")
		deleteTask()
	}

	taskToDelete := input - 1
	tasks = append(tasks[:taskToDelete], tasks[taskToDelete+1:]...)

	fmt.Println("Task deleted successfully!")

}

func printTasks() {
	readFromFile()

	if len(tasks) > 0 {
		fmt.Println("------------Current tasks-------------")
		// range loop

		for i, task := range tasks {
			fmt.Printf("%d. %s", i+1, task)
		}

	}
	if len(doneTasks) > 0 {
		fmt.Println("------------Done tasks----------------")
		for i, task := range doneTasks {
			fmt.Printf("%d. %s", i+1, task)
		}
	}
}

func writeToFile(tasks []string, doneTasks []string) {
	TaskData := TaskData{
		CurrentTasks: tasks,
		DoneTasks:    doneTasks,
	}

	jsonString, _ := json.Marshal(TaskData)
	os.WriteFile("tasks.json", jsonString, 0644)
}

func readFromFile() {
	file, _ := os.ReadFile("tasks.json")

	var taskData TaskData
	json.Unmarshal(file, &taskData)

	tasks = taskData.CurrentTasks
	doneTasks = taskData.DoneTasks

}

func optionSelect() {
	screen.Clear()
	printTasks()

	fmt.Println("---------------------------------")
	fmt.Println("Task manager 📝")
	fmt.Println("Select an option to get started:")
	fmt.Println("---------------------------------")

	fmt.Println("Select an option to get started:")

	// Options
	fmt.Println("[1] Add task 📝 ")
	fmt.Println("[2] Set task status done 🖍")
	fmt.Println("[3] Delete task ❌")
	fmt.Println("[4] Clear current tasks 🧹")
	fmt.Println("[5] Clear done tasks 🧹")
	fmt.Println("[0] Exit 🚪")

	value := readUserInput()

	input, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil || input < 0 || input > 6 {
		fmt.Println("Invalid option, please try again.")
		optionSelect()
	}

	fmt.Println("Selected", input)

	switch input {
	case 1:
		addTask()
	case 2:
		setDone()
	case 3:
		if len(tasks) > 0 {
			deleteTask()
		} else {
			fmt.Println("No tasks to delete, please add a task first.")
			addTask()
		}
	case 4:
		// confirmation
		fmt.Println("Are you sure you want to clear all current tasks? (y/n)")
		if strings.TrimSpace(readUserInput()) != "y" {
			break
		} else {
			fmt.Println("Clearing all current tasks...")
			tasks = []string{}
		}

	case 5:
		fmt.Println("Are you sure you want to clear all done tasks? (y/n)")
		if strings.TrimSpace(readUserInput()) != "y" {
			break
		} else {
			fmt.Println("Clearing all done tasks...")
			doneTasks = []string{}
		}
	case 0:
		return
	default:
		fmt.Println("Invalid option, please try again.")
	}

	writeToFile(tasks, doneTasks)
	optionSelect()
}

func run() int {
	optionSelect()

	fmt.Println("Goodbye 👋")
	return 0
}

func main() {
	os.Exit(run())
}
