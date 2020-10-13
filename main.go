package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {

	repo, _ := git.PlainOpen(".") // path remember u need to find those repositories

	blobs, _ := repo.BlobObjects() // a git object type check the git documentation

	blobs.ForEach(func(b *object.Blob) error {

		fmt.Println(b.Hash)

		return nil
	})
}
