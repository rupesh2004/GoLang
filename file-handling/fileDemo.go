package main

import (
	"fmt"
	"os"
)

func main() {
	files := []struct {
		name        string
		permissions os.FileMode
	}{
		{"file_000.txt", 0000},
		{"file_001.txt", 0001},
		{"file_010.txt", 0010},
		{"file_011.txt", 0011},
		{"file_100.txt", 0100},
		{"file_101.txt", 0101},
		{"file_110.txt", 0110},
		{"file_111.txt", 0111},
	}

	for _, f := range files {
		err := os.WriteFile(f.name, []byte("This is a test file.\n"), f.permissions)
		if err != nil {
			fmt.Printf("Failed to create %s: %v\n", f.name, err)
			return
		}
		fmt.Printf("Created %s with permissions %03o\n", f.name, f.permissions)
	}
}
