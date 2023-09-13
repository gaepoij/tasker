package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

var tasks []string
var doneTasks []string

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
	optionSelect()

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
	optionSelect()

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
	optionSelect()

}

func printTasks() {
	if len(tasks) > 0 {
		fmt.Println("------------Current tasks-------------")
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("%d. %s", i+1, tasks[i])
		}
		fmt.Println("--------------------------------------")
	}

	if len(doneTasks) > 0 {
		fmt.Println("------------Done tasks----------------")
		for i := 0; i < len(doneTasks); i++ {
			fmt.Printf("%d. %s", i+1, doneTasks[i])
		}
		fmt.Println("--------------------------------------")
	}

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
	fmt.Println("[0] Exit 🚪")

	value := readUserInput()

	input, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil || input < 0 || input > 4 {
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
	case 0:
		return
	default:
		fmt.Println("Invalid option, please try again.")
		optionSelect()
	}
}

func run() int {
	optionSelect()

	fmt.Println("Goodbye 👋")
	return 0
}

func main() {
	os.Exit(run())
}
