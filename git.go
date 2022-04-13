package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getRawBranches(gitExec string, mainGitArgs []string) []string {
	cmd := exec.Command(gitExec, mainGitArgs...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("git command failed %v\n", err)
		fmt.Println(cmd.String())
		fmt.Println("fatal: not found repository")
		panic("fatal: not found repository")
	}

	branches := strings.Split(string(stdout), "\n")
	return cutLast(branches)
}

func switchToBranch(gitExec string, gitArgs []string) {
	cmd := exec.Command(gitExec, gitArgs...)
	stdouterr, _ := cmd.CombinedOutput()
	fmt.Println(string(stdouterr))
}
