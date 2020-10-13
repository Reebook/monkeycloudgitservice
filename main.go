package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {

	repo, _ := git.PlainOpen(".") // path remember u need to find those repositories

	objects, _ := repo.Objects()

	objects.ForEach(func(db object.Object) error {

		fmt.Println(db)

		return nil
	})
}
