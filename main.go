package main

import (
	"bufio"
	"fmt"
	"notes/fileops"
	"notes/n"
	"os"
	"strings"
)

func main() {
	var note string = collectData("Enter note title")
	var noteContent string = collectData("Enter note content")
	var newNote, err = n.New(note, noteContent)
	if err != nil {
		fmt.Print(err)
		return
	}

	fileops.WriteToFile(newNote)
	fmt.Println("Note saved to", fileops.StorageFile)

}

func collectData(promptText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(promptText)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	return value
}
