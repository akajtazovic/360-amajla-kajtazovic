package main

import (
	"bufio"
	"fmt"
	"os"
)

//commentTTT

type Validation interface {
	Lic() bool
	ReadMe() bool
	LineCheck() bool
	CommentCheck() bool
}

//comment

type Dir string
type StrDir struct {
	strDirect string
}

func rootdir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var userinput = ""

	for userinput != "val" {
		scanner.Scan()
		userinput = scanner.Text()
		fmt.Println("In: ", userinput)
		if userinput == "val detail" {
			fmt.Println("DETAILS: ")
			totalerrors := theError + LineCheckError + LicenseError + CommentCheckError
			fmt.Println("Total Errors: ", totalerrors)
			fmt.Println("If any errors popped up, the location will be below.")
			fmt.Println("If the area is blank, no error was detected.")
			fmt.Println(LineErrorCheckL)
			fmt.Println(LicenseErrorL)

		}
		if userinput == "val help" {
			fmt.Println("HELP:")
			fmt.Println("")
			fmt.Println("Lic() checks if you have a license file.")
			fmt.Println("ReadMe() checks if you have a readme file.")
			fmt.Println("LineCheck() checks if your line characters are under 100.")
			fmt.Println("CommentCheck() checks if you have comments")
			fmt.Println("")
			fmt.Println("You can find a summary of each validation after it is completed,")
			fmt.Println("or you can use val detail for a summary of the total errors,")
			fmt.Println("and where they are located.")
		}
		if userinput == "val" {
			var v Validation
			dir := Dir(rootdir())
			v = dir
			v.Lic()
			v.LineCheck()
			v.CommentCheck()
			v.ReadMe()
			str := StrDir{rootdir()}
			v = str
			v.Lic()
			v.LineCheck()
			v.CommentCheck()
			v.ReadMe()
		}
	}
}
