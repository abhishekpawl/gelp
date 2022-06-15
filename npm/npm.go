package npm

import (
	"fmt"
	"gelp/helper"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen, color.Bold)

var cyan = color.New(color.FgCyan, color.Bold)

func SpinupNpm() {
	cyan.Print("Enter project name: ")
	var dir string
	fmt.Scanln(&dir)

	err1 := os.Mkdir(dir, os.ModePerm)
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := os.Chdir(dir)
	if err2 != nil {
		log.Fatal(err2)
	}

	cmd3 := exec.Command("npm", "init", "-y")
	err3 := cmd3.Run()
	if err3 != nil {
		log.Fatal(err3)
	}
	green.Println("NPM project initialized.")

	curr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	green.Print("Project instantiated at path: ")
	cyan.Add(color.Underline).Printf("%+v\n", curr)

	helper.Code()
}
