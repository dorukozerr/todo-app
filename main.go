package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Todo struct {
	Text    string `json:"text"`
	Done    bool   `json:"done"`
	Urgency int    `json:"urgency"`
}

type UrgencyLevel struct {
	Name  string
	Value int
	Color string
}

var urgencyLevels = []UrgencyLevel{
	{Name: "Critical", Value: 1, Color: "\033[35m"},  // Magenta
	{Name: "High", Value: 2, Color: "\033[31m"},      // Red
	{Name: "Medium", Value: 3, Color: "\033[33m"},    // Yellow
	{Name: "Low", Value: 4, Color: "\033[36m"},       // Cyan
	{Name: "Very Low", Value: 5, Color: "\033[32m"},  // Green
}

var todoDir string
var todoFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	todoDir = filepath.Join(homeDir, ".todo")
	todoFile = filepath.Join(todoDir, "todos.json")

	if err := os.MkdirAll(todoDir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}
}

func selectUrgency() int {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\nSelect Urgency Level")
		for i, level := range urgencyLevels {
			fmt.Printf("%d %s%s\033[0m\n", i+1, level.Color, level.Name)
		}

		fmt.Print("\nEnter urgency (1-5): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		urgency, err := strconv.Atoi(input)
		if err != nil || urgency < 1 || urgency > 5 {
			fmt.Println("Error: Please enter a number between 1 and 5")
			continue
		}
		
		return urgency
	}
}

func saveTodos(todos []Todo) error {
	file, err := os.Create(todoFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(todos)
}

func loadTodos() []Todo {
	var todos []Todo
	file, err := os.Open(todoFile)
	if err != nil {
		return []Todo{}
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)
	return todos
}

func getColorByUrgency(urgency int) string {
	for _, level := range urgencyLevels {
		if level.Value == urgency {
			return level.Color
		}
	}
	return "\033[0m" // Reset
}

func displayTodos(todos []Todo, showAll bool) {
	reset := "\033[0m"
	for i, todo := range todos {
		if !todo.Done || showAll {
			color := getColorByUrgency(todo.Urgency)
			doneStr := " "
			if todo.Done {
				doneStr = "✓"
			}
			fmt.Printf("%d. [%s] %s%s%s\n", i+1, doneStr, color, todo.Text, reset)
		}
	}
}

func main() {
	todos := loadTodos()
	reader := bufio.NewReader(os.Stdin)

	// Get the base command name from the executable path
	command := filepath.Base(os.Args[0])

	// Get all arguments after the command
	args := os.Args[1:]

	// Special handling for 'td -a' case
	if command == "td" && len(args) > 0 && args[0] == "-a" {
		command = "-a"
		args = args[1:]
	}

	switch command {
	case "td", "main":
		displayTodos(todos, false)

	case "-a":
		displayTodos(todos, true)

	case "tda":
		fmt.Print("Enter todo text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		urgency := selectUrgency()

		todos = append(todos, Todo{
			Text:    text,
			Done:    false,
			Urgency: urgency,
		})
		if err := saveTodos(todos); err != nil {
			fmt.Println("Error saving todo:", err)
		}

	case "tdd":
		if len(args) < 1 {
			fmt.Println("Please provide todo index")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("Invalid index")
			return
		}
		todos[index-1].Done = true
		if err := saveTodos(todos); err != nil {
			fmt.Println("Error saving todo:", err)
		}

	case "tdr":
		if len(args) < 1 {
			fmt.Println("Please provide todo index")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("Invalid index")
			return
		}
		todos = append(todos[:index-1], todos[index:]...)
		if err := saveTodos(todos); err != nil {
			fmt.Println("Error saving todo:", err)
		}

	default:
		fmt.Println("Usage:")
		fmt.Println("  td      - List active todos")
		fmt.Println("  td -a   - List all todos")
		fmt.Println("  tda     - Add new todo")
		fmt.Println("  tdd N   - Mark todo N as done")
		fmt.Println("  tdr N   - Remove todo N")
	}
}
