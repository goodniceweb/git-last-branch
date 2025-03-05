package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getBranches() gitBranches {
	return gitBranches{
		gitBranch{
			Branch: "super-duper-long-branch-name",
			LastCommit: "f51306e",
			AuthorDate: "Wed Apr 13 14:59:39 2022 +0000",
			Text: "super-duper-long-branch-name  Wed Apr 13 14:59:39 2022 +0000",
		},
		gitBranch{
			Branch: "main",
			LastCommit: "f51316e",
			AuthorDate: "Wed Apr 13 14:59:09 2022 +0000",
			Text: "main                          Wed Apr 13 14:59:09 2022 +0000",
		},
		gitBranch{
			Branch: "very-long-branch-name",
			LastCommit: "f51326e",
			AuthorDate: "Wed Apr 13 14:58:09 2022 +0000",
			Text: "very-long-branch-name         Wed Apr 13 14:58:09 2022 +0000",
		},
	}
}

func TestMapGitBranches(t *testing.T) {
	bs := getBranches()
	actual := bs.Map(func (b gitBranch) string {
		return b.Branch
	})
	expected := []string{"super-duper-long-branch-name", "main", "very-long-branch-name"}
	assert.Equal(t, expected, actual, "Map function should iterate over items and return function result")
}

func TestBuildBranches(t *testing.T) {
	rawBranches := []string{
		"refs/heads/super-duper-long-branch-name\tf51306e\tWed Apr 13 14:59:39 2022 +0000",
		"refs/heads/main\tf51316e\tWed Apr 13 14:59:09 2022 +0000",
		"refs/heads/very-long-branch-name\tf51326e\tWed Apr 13 14:58:09 2022 +0000",
	}
	actual := BuildBranches(rawBranches)
	expected := getBranches()
	assert.Equal(t, expected, actual, "BuildBranches should create a correct array of struct")
}