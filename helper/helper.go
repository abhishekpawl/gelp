package helper

import (
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed, color.Bold)

var green = color.New(color.FgGreen, color.Bold)

var cyan = color.New(color.FgCyan, color.Bold)

func Code() {
	bin, _ := exec.LookPath("code")
	if strings.Contains(bin, "code") {
		green.Println("VSCode found on Device.")
		cyan.Println("Opening project in VSCode...")
		cmd4 := exec.Command("code", ".")
		err4 := cmd4.Run()
		if err4 != nil {
			log.Fatal(err4)
		}
	} else {
		red.Println("VSCode not found on Device.")
	}
}
