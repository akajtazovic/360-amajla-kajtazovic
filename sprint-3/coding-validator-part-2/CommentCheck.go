package main

import (
	"fmt"
	"os"
	"strings"
)

var CommentCheckError int = 0

func (d Dir) CommentCheck() bool {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	theFile, _ := os.ReadFile(gofile)
	lines := string(theFile)
	comment := "//"
	if strings.ContainsAny(comment, lines) {
		fmt.Println("There's a comment.")
		return true
	} else {
		fmt.Println("There is no comment")
		CommentCheckError = 1
		return false
	}
}

func (sd StrDir) CommentCheck() bool {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	theFile, _ := os.ReadFile(gofile)
	lines := string(theFile)
	comment := "//"
	if strings.ContainsAny(comment, lines) {
		fmt.Println("There's a comment.")
	} else {
		fmt.Println("There is no comment.")
		CommentCheckError = 1
		return false
	}
	return true
}
