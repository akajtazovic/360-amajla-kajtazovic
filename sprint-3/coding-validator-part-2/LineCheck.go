package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var LineErrorCheckL string

func LineNumber(l int) (error, bool) {
	if l > 100 {
		return errors.New("line number too long"), false
	}
	return nil, true
}

var LineCheckError int = 0

func (d Dir) LineCheck() bool {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err != nil {
		LineCheckError = 1
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		LineCheckError = 1
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("total lines: ", lineCount)
		lc, _ := LineNumber(lineCount)
		fmt.Println("LINE CHECK:")
		fmt.Println("")
		if lc == nil {
			fmt.Println("PASS")
		}
		if lc != nil {
			LineCheckError++
			pc, filename, line, _ := runtime.Caller(1)
			LineErrorCheckL = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("FAIL")
			return false
		}
	}
	return true
}

func (sd StrDir) LineCheck() bool {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		LineCheckError = 1
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		LineCheckError = 1
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("total lines: ", lineCount)
		lc, _ := LineNumber(lineCount)
		fmt.Println("LINE CHECK STARTS HERE")
		fmt.Println("")
		if lc == nil {
			fmt.Println("")
			fmt.Println("PASS")
		}
		if lc != nil {
			LineCheckError++
			pc, filename, line, _ := runtime.Caller(1)
			LineErrorCheckL = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
			return false
		}
	}
	return true
}
