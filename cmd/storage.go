package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadTasks() ([]string, error) {
	file, err := os.Open(todoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to open %s: %w", todoPath, err)
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- ") {
			task := strings.TrimPrefix(line, "- ")
			tasks = append(tasks, task)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", todoPath, err)
	}

	return tasks, nil
}

func saveTasks(tasks []string) error {
	// Ensure parent directory exists
	dir := strings.TrimSuffix(todoPath, "/"+getFileName(todoPath))
	if dir != "" && dir != todoPath {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	file, err := os.Create(todoPath)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", todoPath, err)
	}
	defer file.Close()

	// Write a simple markdown header
	fmt.Fprintf(file, "# Todos\n\n")
	for _, task := range tasks {
		fmt.Fprintf(file, "- %s\n", task)
	}

	return nil
}

func addTask(task string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}
	tasks = append(tasks, task)
	return saveTasks(tasks)
}

func getFileName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
