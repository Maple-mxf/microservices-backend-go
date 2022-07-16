package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStatFile(t *testing.T) {
	fileInfo, err := os.Stat("./backend.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.IsDir())
}
