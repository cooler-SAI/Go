package main

import (
	"fmt"
	"os"
	"testing"
)

// Тест для записи данных в файл
func TestWriteToFile(t *testing.T) {
	// Подготовка тестовых данных
	posts := []Post{
		{UserID: 10, ID: 91, Title: "test title 1", Body: "test body 1"},
		{UserID: 10, ID: 92, Title: "test title 2", Body: "test body 2"},
	}

	// Создание временного файла
	file, err := os.CreateTemp("", "posts_test_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(file.Name()) // Удаляем файл после теста
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Запись данных в файл
	for _, post := range posts {
		line := fmt.Sprintf("ID: %d, Title: %s\n", post.ID, post.Title)
		_, err := file.WriteString(line)
		if err != nil {
			t.Fatalf("Error writing to file: %v", err)
		}
	}

	// Проверка содержимого файла
	_, _ = file.Seek(0, 0) // Перемещаем указатель в начало файла
	content, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	expectedContent := "ID: 91, Title: test title 1\nID: 92, Title: test title 2\n"
	if string(content) != expectedContent {
		t.Errorf("Expected file content:\n%s\nGot:\n%s", expectedContent, string(content))
	}
}
