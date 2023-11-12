package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diazadriel0/go-course/exercises"
)

var filePath string = "./files/txt/table.txt"

func SaveTable() {
	var table = exercises.MultiplicationTable()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file " + err.Error())
		return
	}
	fmt.Fprintln(file, table)
	file.Close()
}

func SumTable() {
	var table = exercises.MultiplicationTable()
	
	if !Append(table) {
		fmt.Println("Error concatenating the content")
	}

}

func Append(table string) bool {
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error Appending " + err.Error())
		file.Close()
		return false
	}

	_, err = file.WriteString(table)
	if err != nil {
		fmt.Println("Error writting string " + err.Error())
		file.Close()
		return false
	} 

	file.Close()
	return true
}

func ReadFile() {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file " + err.Error())
		return
	}

	fmt.Println(string(file))
}

func ReadLines() {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file " + err.Error())
		file.Close()
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">> " + line)
	}

	file.Close()
}