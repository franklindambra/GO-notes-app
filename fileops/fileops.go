package fileops

import (
	"encoding/json"
	"fmt"
	"notes/n"
	"os"
)

const StorageFile = "storage.json"

func WriteToFile(note *n.Note) {

	var notes []n.Note

	data, err := os.ReadFile(StorageFile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error reading file:", err)
			return
		}
	} else if len(data) > 0 {
		if err := json.Unmarshal(data, &notes); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
	}

	notes = append(notes, *note)

	formattedData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		fmt.Println("marshal error", err)
		return
	}

	if err := os.WriteFile(StorageFile, formattedData, 0644); err != nil {
		fmt.Println("write error:", err)
	}

}
