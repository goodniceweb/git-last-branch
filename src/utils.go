package main

import "os"

func getCurrentPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic("can't get current path")
	}
	return currentPath
}

func cutLast(slice []string) []string {
	return slice[:len(slice) - 1]
}

func longestStringIn(slice []string) string {
	best, length := "", 0
	for _, word := range slice {
			if len(word) > length {
					best, length = word, len(word)
			}
	}
	return best
}
