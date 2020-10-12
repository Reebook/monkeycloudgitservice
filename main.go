package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {

	// path of the repository
	// remember the need to check if the path provided is a git repository

	repo, _ := git.PlainOpen(".")

	ref, _ := repo.Head()

	commit, _ := repo.CommitObject(ref.Hash())

	files, _ := commit.Files()

	files.ForEach(func(f *object.File) error {
		fmt.Println(f.Hash, f.Name)
		return nil
	})

}
