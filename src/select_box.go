package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func getSelectedBranch(branches []gitBranch) string {
	prompt := promptui.Select{
		Label: "Select branch",
		Items: branches,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   "=> {{ .Text | cyan }}",
			Inactive: "   {{ .Text | cyan }}",
			Selected: " ",
		},
	}
	resultIndex, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Select failed %v\n", err)
		panic("unexpected error while selecting git branch")
	}

	return branches[resultIndex].Branch
} 
