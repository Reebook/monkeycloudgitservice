package main

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

	//	Return all branch name
	branchesName, err := getBranches(".")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Branches")
		for i, s := range branchesName {
			fmt.Println(i, s)
		}
	}

	//Return if exist a specific branch
	branchName := "devnotes2"
	result, _ := getBranch(".", branchName)
	if result {
		fmt.Println("Exist " + branchName)
	}
	InitRepository("/app/Repos/newRepo", false)

	//Initalize a Object Git.CloneOptions
	option := git.CloneOptions{
		URL:               "https://github.com/Colibri-code/monkeycloudgitservice",
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}
	//Initalize a repoCloneOptions
	repoCloneOptions := RepoCloneOptions{"/app/Repos/test2/",
		false,
		"master",
		false,
		false,
		option,
	}
	//Clonning a Repository
	CloneRepository(repoCloneOptions)
}

//Return all branches specific repository
// If only have one branch "master" dont return it
func getBranches(repoPath string) ([]string, error) {
	repo, err := git.PlainOpen(repoPath)
	const BranchPrefix = "refs/heads/"
	var branchNames []string
	if err != nil {
		return nil, err
	}

	if repo != nil {
		branches, _ := repo.Branches()

		if branches != nil {

			_ = branches.ForEach(func(branch *plumbing.Reference) error {
				branchNames = append(branchNames, strings.TrimPrefix(branch.Name().String(), BranchPrefix))
				return nil
			})

			return branchNames, err
		}
	}

	return branchNames, nil
}

//Verify if a branch exist
func getBranch(repoPath string, branchName string) (bool, error) {
	const BranchPrefix = "refs/heads/"
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return false, err
	}

	_, err = repo.Reference(plumbing.ReferenceName(BranchPrefix+branchName), true)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// initializes a new Git repository.
func InitRepository(repoPath string, isBare bool) (*Repository, error) {
	r, err := git.PlainInit(repoPath, isBare)

	if err != nil {
		return nil, err
	}

	return &Repository{Path: repoPath,
		goRepo: r}, nil
}

//CloneRepository clone a repository
func CloneRepository(repoCloneOptions RepoCloneOptions) (*Repository, error) {
	repository, err := git.PlainClone(repoCloneOptions.dir, false, &repoCloneOptions.cloneOptions)

	if err != nil {
		return nil, err
	}

	return &Repository{Path: repoCloneOptions.dir,
		goRepo: repository}, nil
}

//Objects
// Git Repository
type Repository struct {
	Path   string
	goRepo *git.Repository
}

type RepoCloneOptions struct {
	dir          string
	Bare         bool
	Branch       string
	Shared       bool
	NoCheckout   bool
	cloneOptions git.CloneOptions
}
