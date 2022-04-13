package main

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

type gitBranches []gitBranch
type gitBranch struct {
	Branch string
	LastCommit string
	AuthorDate string
	Text string
}

func (branches gitBranches) Map(f func(gitBranch) string) []string {
	vsm := make([]string, len(branches))
	for i, v := range branches {
			vsm[i] = f(v)
	}
	return vsm
}

func GetFormattedOutput(branch gitBranch, branches gitBranches) string {
	allBranches := branches.Map(func (b gitBranch) string {
		return b.Branch
	})
	return text.AlignLeft.Apply(branch.Branch, len(longestStringIn(allBranches)) + 2) + branch.AuthorDate
}

func BuildBranches(rawBranches []string) []gitBranch {
	branchesStruct := make([]gitBranch, len(rawBranches))
	for i, b := range rawBranches {
		parts := strings.Split(b, "\t")
		branchesStruct[i].Branch = strings.Replace(parts[0], "refs/heads/", "", 1)
		branchesStruct[i].LastCommit = parts[1]
		branchesStruct[i].AuthorDate = parts[2]
	}
	for i, b := range branchesStruct {
		branchesStruct[i].Text = GetFormattedOutput(b, branchesStruct)
	}
	return branchesStruct
}
