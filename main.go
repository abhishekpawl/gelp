package main

import (
	"gelp/goHelp"
	"gelp/npm"
	"os"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed, color.Bold)

var green = color.New(color.FgGreen, color.Bold)

var cyan = color.New(color.FgCyan, color.Bold)

func info() {
	cyan.Println("[COMMANDS]")
	cyan.Println("1. -V --version - current version of GELP")
}

func validTech(tech string) bool {
	if tech == "react" || tech == "npm" || tech == "go" {
		return true
	}
	return false
}

func main() {
	/* link := cyan.Add(color.Underline) */
	args := os.Args

	if len(args) == 1 {
		green.Println("Welcome to GELP!")
		red.Printf("\n")
		info()
	} else {
		for idx, com := range args {
			if idx >= 1 {
				switch com {
				case "--version", "-V":
					cyan.Println("GELP V1.0.0")
				case "--tech", "-T":
					if idx+1 >= len(args) {
						red.Println("VALUE MISSING...")
					} else {
						value := args[idx+1]
						if value[0] == '-' {
							red.Println("BAD SYNTAX...")
						} else if validTech(value) {
						} else {
							red.Println("TECH NOT RECOGNIZED...")
							info()
						}
					}
				case "react", "npm", "go":
					check := false
					for _, flag := range args {
						if flag == "-T" || flag == "--tech" {
							check = !check
						}
					}
					if !check {
						red.Println("FLAG MISSING...")
						info()
					} else {
						switch com {
						case "npm":
							npm.SpinUpNpm()
						case "go":
							goHelp.SpinupGo()
						}
					}
				default:
					red.Println("BAD COMMAND...")
				}
			}
		}
	}
}
