package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Функція для копіювання файлу
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	err = out.Sync()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Створення директорії Test2
	err := os.Mkdir("Test2", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Копіювання файлу test.txt до Test2 та перейменування його у test2.txt
	src := "Test/test.txt"
	dst := "Test2/test2.txt"
	err = copyFile(src, dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Створення порожніх файлів test3.txt та test4.txt
	err = ioutil.WriteFile("Test2/test3.txt", []byte{}, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("Test2/test4.txt", []byte{}, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Отримання списку файлів у директорії Test2 та виведення їх назв на екран
	files, err := ioutil.ReadDir("Test2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Файли у директорії Test2:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
