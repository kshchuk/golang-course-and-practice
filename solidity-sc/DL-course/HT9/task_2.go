package main

import (
	"fmt"
	"io/fs"
	"os"
)

func GetMeta(path string) (fs.FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}

func main() {
	path := "Test/test.txt"
	fileInfo, err := GetMeta(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Назва файлу:", fileInfo.Name())
	fmt.Println("Розмір файлу в байтах:", fileInfo.Size())
	fmt.Println("Дозволи на доступ:", fileInfo.Mode())
	fmt.Println("Дата та час останньої модифікації:", fileInfo.ModTime())
	fmt.Println("Чи є файл директорією:", fileInfo.IsDir())
}
