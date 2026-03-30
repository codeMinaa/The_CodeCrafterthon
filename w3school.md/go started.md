# Go Get Started
To start using Go, you need two things:

* A text editor, like VS Code, to write Go code
* A compiler, like GCC, to translate the Go code into a language that the computer will understand
There are many text editors and compilers to choose from. In this tutorial, we will use an IDE (see below).


# Go Install
You can find the relevant installation files at https://golang.org/dl/.

Follow the instructions related to your operating system. To check if Go was installed successfully, you can run the following command in a terminal window:

go version
Which should show the version of your Go installation.

# Go Install IDE
An IDE (Integrated Development Environment) is used to edit AND compile the code.

Popular IDE's include Visual Studio Code (VS Code), Vim, Eclipse, and Notepad. These are all free, and they can be used to both edit and debug Go code.

Note: Web-based IDE's can work as well, but functionality is limited.

We will use VS Code in our tutorial, which we believe is a good place to start.

You can find the latest version of VS Code at https://code.visualstudio.com/.

# Go Quickstart
Let's create our first Go program.

Launch the VS Code editor
Open the extension manager or alternatively, press Ctrl + Shift + x
In the search box, type "go" and hit enter
Find the Go extension by the GO team at Google and install the extension
After the installation is complete, open the command palette by pressing Ctrl + Shift + p
Run the Go: Install/Update Tools command
Select all the provided tools and click OK
VS Code is now configured to use Go.

Open up a terminal window and type:

go mod init example.com/hello
Do not worry if you do not understand why we type the above command. Just think of it as something that you always do, and that you will learn more about in a later chapter.

Create a new file (File > New File). Copy and paste the following code and save the file as helloworld.go (File > Save As):
* Example:
```go 
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
.

