package main

import (
	"fmt"
	"math"
	"strings"
)

func getRawBranches(gitExec string, mainGitArgs []string) []string {
	stdout := runCommandWithHandledOutput(gitExec, mainGitArgs)
	branches := strings.Split(string(stdout), "\n")
	min := math.Min(float64(len(branches)), float64(globalConfig.AmountOfBranches))
	branches = branches[0:int(min)]
	return cutLast(branches)
}

func switchToBranch(gitExec string, gitArgs []string) {
	fmt.Println(string(runCommandAndReturnAllOutput(gitExec, gitArgs)))
}
