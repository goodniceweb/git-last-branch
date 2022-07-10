package main

// TODO: [feature] highlight current branch
// TODO: [feature] settings - limit for tons of branches
// TODO: [feature] cli arguments for overriding settings
// TODO: [refactor] move magic literals to config
// TODO: [nice-to-have] verbose error vs regular mode
// TODO: [final] test and build for linux/mac/windows

var globalConfig config = config{
	AmountOfBranches: 5,
	SpacesBetweenColumns: 2,
}

func main() {
	gitExec := "git"
	currentPathArgs := []string{
		"-C",
		getCurrentPath(),
	}
	lastBranchArgs := []string{
	  "for-each-ref",
	  "--sort",
	  "-committerdate",
	  "--format",
	  "%(refname)\t%(objectname:short)\t%(committerdate)",
	  "refs/heads",
	}
	mainGitArgs := append(currentPathArgs, lastBranchArgs...)
	rawBranches := getRawBranches(gitExec, mainGitArgs)
	branches := BuildBranches(rawBranches)
	selectedBranch := getSelectedBranch(branches)
	switchBranchGitArgs := append(currentPathArgs, "checkout", selectedBranch)
	switchToBranch(gitExec, switchBranchGitArgs)
}