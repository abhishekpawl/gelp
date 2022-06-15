package react

import (
	"fmt"
	"gelp/helper"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen, color.Bold)

var cyan = color.New(color.FgCyan, color.Bold)

func SpinupReact() {
	cyan.Print("Enter project name: ")
	var dir string
	fmt.Scanln(&dir)

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		/* go func() { */
		cyan.Println("Initialization may take some time...")
		cmd1 := exec.Command("npx", "create-react-app", dir)
		err1 := cmd1.Run()
		if err1 != nil {
			log.Fatal(err1)
		}
		wg.Done()
		/* }() */
	}()

	wg.Wait()

	green.Println("React project initialized successfully.")

	curr, err2 := os.Getwd()
	if err2 != nil {
		log.Fatal(err2)
	}
	green.Print("Project instantiated at path: ")
	cyan.Add(color.Underline).Printf("%+v\n", curr)

	err3 := os.Chdir(dir)
	if err3 != nil {
		log.Fatal(err3)
	}

	helper.Code()
}
