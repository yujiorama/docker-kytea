package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	command := "help"
	if len(os.Args) > 0 {
		command = os.Args[1]
	}

	switch command {
	case "kytea", "train-kytea":
		program := strings.Join(
			[]string{os.Getenv("KYTEA_DIR"), "bin", command},
			"/",
		)
		cmd := exec.Command(program, os.Args[2:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
	case "help":
		program := strings.Join(
			[]string{os.Getenv("KYTEA_DIR"), "bin", "kytea"},
			"/",
		)
		cmd := exec.Command(program, "--help")
		cmd.Run()
	}
}
