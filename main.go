package main

import (
	"fmt"

	"github.com/go-git/go-git"
)

func main() {

	// path of the repository
	// remember the need to check if the path provided is a git repository

	repo, _ := git.PlainOpen(".")

	ref, _ := repo.Head()

	commit, _ := repo.CommitObject(ref.Hash())

	fmt.Println(commit.)

}
