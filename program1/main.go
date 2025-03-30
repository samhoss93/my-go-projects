package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Define a Task struct
type Task struct {
	ID        int
	Name      string
	Status    string
	CreatedAt time.Time
}

// TaskManager manages a list of tasks
type TaskManager struct {
	Tasks map[int]*Task
	mu    sync.Mutex // For thread-safe access
}

// NewTaskManager initializes a new TaskManager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[int]*Task),
	}
}

// AddTask adds a new task (pointer receiver)
func (tm *TaskManager) AddTask(name string) (int, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if name == "" {
		return 0, errors.New("task name cannot be empty")
	}

	id := len(tm.Tasks) + 1
	task := &Task{
		ID:        id,
		Name:      name,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}
	tm.Tasks[id] = task
	return id, nil
}

// CompleteTask marks a task as done
func (tm *TaskManager) CompleteTask(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.Tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %d not found", id)
	}
	task.Status = "Completed"
	return nil
}

// Worker processes tasks concurrently (goroutine)
func Worker(tm *TaskManager, taskID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker: Processing task %d\n", taskID)
	time.Sleep(2 * time.Second) // Simulate work

	err := tm.CompleteTask(taskID)
	if err != nil {
		fmt.Printf("Error completing task %d: %v\n", taskID, err)
		return
	}
	fmt.Printf("Worker: Task %d completed!\n", taskID)
}

func main() {
	tm := NewTaskManager()

	// Add tasks
	task1, _ := tm.AddTask("Write Go code")
	task2, _ := tm.AddTask("Review PRs")
	task3, _ := tm.AddTask("Deploy to production")

	// Process tasks concurrently
	var wg sync.WaitGroup
	wg.Add(3)
	go Worker(tm, task1, &wg)
	go Worker(tm, task2, &wg)
	go Worker(tm, task3, &wg)
	wg.Wait() // Wait for all goroutines to finish

	// Print final task statuses
	fmt.Println("\nFinal Task Statuses:")
	for id, task := range tm.Tasks {
		fmt.Printf("Task %d: %s (%s)\n", id, task.Name, task.Status)
	}
	foorbar() // Call foobar
}
