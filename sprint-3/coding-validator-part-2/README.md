# validate-demo
#Amajla Kajtazovic | amajlakajtazovic@lewisu.edu | #validator-project

Credits:
a lot of the format for the whole project came from Mr. Biryukov's examples, so I owe a lot to him.

Instructions for Validations:
- all of them have their own separate file in the same directory. They will be called together in main() using the Validations Interface.

Lic() - Checks if there is a license file and also checks if there is content on the inside, if not, prints error.

ReadMe() - checks if there is a readme file, if not, prints error.

LineCheck() - this checks if each line is under 100 characters. it does this by scanning the line and taking the length of the line. I use a function called lineError to say that if the line is over 100: there is an error.

CommentCheck() - This checks if there is a comment. It does this by using the strings.ContainsAny() function.

AUTOMATED TESTING:
1. Enter the command "go test" on the terminal to start the tests.
2. If you would like all outputs to be clearly shown. Type "go test -v"
3. To see the code coverage, type "go test -v --cover"

MANUAL TESTING:
1. Manual testing involves going through the program as a user.
2. Go through the program and make sure each function is doing what needs to be done. 
for example, to check CommentCheck() delete all comments, and see if it catches the error. If it fails, it is doing it's job.

