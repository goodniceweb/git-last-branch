package main

// TODO: [must] automated tests
// TODO: [must] default limit for amount of branches
// TODO: [feature] highlight current branch
// TODO: [feature] settings - limit for tons of branches
// TODO: [feature] cli arguments for overriding settings
// TODO: [refactor] move magic literals to config
// TODO: [nice-to-have] verbose error vs regular mode
// TODO: [final] test and build for linux/mac/windows

func main() {
	gitExec := "git"
	currentPathArgs := []string{
		"-C",
		getCurrentPath(),
	}
	lastBranchArgs := []string{
	  "for-each-ref",
	  "--sort",
	  "-authordate",
	  "--format",
	  "%(refname)\t%(objectname:short)\t%(authordate)",
	  "refs/heads",
	}
	mainGitArgs := append(currentPathArgs, lastBranchArgs...)
	rawBranches := getRawBranches(gitExec, mainGitArgs)
	branches := BuildBranches(rawBranches)
	selectedBranch := getSelectedBranch(branches)
	switchBranchGitArgs := append(currentPathArgs, "checkout", selectedBranch)
	switchToBranch(gitExec, switchBranchGitArgs)
}