package main

import (
	"fmt"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func main () {
	gitExec := "git"
	gitArg0 := "for-each-ref"
	gitArg1 := "--sort='-authordate'"
	gitArg2 := "--format='%(refname)(objectname:short)(authordate)'"
	gitArg3 := "refs/heads"
	cmd := exec.Command(gitExec, gitArg0, gitArg1, gitArg2, gitArg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("fatal: not a git repository")
		return
	}
	fmt.Println(string(stdout))

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}