package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestLineCheck(t *testing.T) {
	var v Validation
	dir := Dir(rootdir())
	v = dir
	lines := v.LineCheck()

	if lines != true {
		t.Error("Line is over 100 mark")
	} else {
		t.Log("The line is under 100.")
	}
}

func TestReadMe(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "README.md"
	var v Validation
	v = Dir(strPath)
	file := v.ReadMe()

	if file != true {
		t.Error("No ReadMe File")
	} else {
		t.Log("There is a README.md file")
	}
}

func TestCommentCheck(t *testing.T) {
	var v Validation
	dir := Dir(rootdir())
	v = dir
	Comment := v.CommentCheck()

	if Comment != true {
		t.Error("No Comments")
	} else {
		t.Log("There is at least one comment")
	}

}
func TestLicense(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "LICENSE"
	var v Validation
	v = Dir(strPath)
	file := v.Lic()

	if file != true {
		t.Error("No License File")
	} else {
		t.Log("There is a License File")
	}
}

func TestRootdir(t *testing.T) {
	result := rootdir()
	currentdir, err := os.Getwd()
	if result != currentdir {
		fmt.Println(err)
		t.Error("Needs current directory")
	} else {
		t.Log("It is in current directory")
	}
}

func TestReadMeStrDirect(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "README.md"
	var v Validation
	v = StrDir{strPath}
	file := v.ReadMe()

	if file != true {
		t.Error("No ReadMe File")
	} else {
		t.Log("There is a readme file")
	}
}
func TestLicStrDirect(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "LICENSE"
	var v Validation
	v = StrDir{strPath}
	file := v.Lic()

	if file != true {
		t.Error("There is no License File")
	} else {
		t.Log("There is a License File")
	}
}

func TestLineCheckStrDir(t *testing.T) {
	var v Validation
	str := StrDir{rootdir()}
	v = str
	lines := v.LineCheck()

	if lines != true {
		t.Error("Lines are too long")
	} else {
		t.Log("Lines are all within range")
	}
}

func TestCommentCheckStrDir(t *testing.T) {
	var v Validation
	str := StrDir{rootdir()}
	v = str
	comment := v.CommentCheck()

	if comment != true {
		t.Error("NO COMMENT")
	} else {
		t.Log("THERE IS AT LEAST ONE COMMENT")
	}
}

func TestLineNumber(t *testing.T) {
	fil, err := os.Open("val.go")
	if err != nil {
		t.Error("CANNOT OPEN VAL.GO FILE")
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
		line := len(lines)
		l, _ := LineNumber(line)

		if l != nil {
			t.Error("LINE IS TOO LONG")
		} else {
			t.Log("LINE IS WITHIN RANGE")
		}
	}
}
