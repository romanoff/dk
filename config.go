package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateConfig(options map[string]string) {
	perm := os.FileMode(0664)
	ioutil.WriteFile(".dk", []byte{}, perm)
	fmt.Println("dk config has been created")
}
