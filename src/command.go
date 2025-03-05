package main

import (
	"fmt"
	"os/exec"
)

func runCommandWithHandledOutput(command string, args []string) []byte {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("command got failed %v\n", err)
		fmt.Println(cmd.String())
		panic(err.Error())
	}

	return stdout
}

func runCommandAndReturnAllOutput(command string, args []string) []byte {
	cmd := exec.Command(command, args...)
	stdouterr, _ := cmd.CombinedOutput()
	return stdouterr
}
