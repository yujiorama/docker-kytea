package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	command := "help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "chasenize":
		kyteaArgs := []string{
			"kytea",
			"-wordbound",
			"\n",
			"-tagbound",
			",",
			"-tagmax",
			"0",
		}
		kytea := exec.Command("/docker-entrypoint", kyteaArgs...)
		kytea.Stdin = os.Stdin

		program := strings.Join(
			[]string{os.Getenv("KYTEA_DIR"), "bin", command},
			"/",
		)
		chasenize := exec.Command(program, os.Args[2:]...)
		var err error
		chasenize.Stdin, err = kytea.StdoutPipe()
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
		chasenize.Stderr = os.Stderr
		chasenize.Stdout = os.Stdout

		if err := chasenize.Start(); err != nil {
			log.Fatalf("Error: %v", err.Error())
		}

		if err := kytea.Run(); err != nil {
			log.Fatalf("Error: %v", err.Error())
		}

		if err := chasenize.Wait(); err != nil {
			log.Fatalf("Error: %v", err.Error())
		}

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
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run()
	}
}
