package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

type GitRepo struct {
	*git.Repository
}

func Repo() (GitRepo, error) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		return GitRepo{}, err
	}

	return GitRepo{r}, nil
}

func (r GitRepo) Branch() (string, error) {
	ref, err := r.Head()
	if err != nil {
		return "", err
	}

	return ref.Name().Short(), nil
}

func (r GitRepo) TrackingBranch() (string, error) {
	ref, err := r.Head()
	if err != nil {
		return "", err
	}

	if !ref.Name().IsBranch() {
		return "", nil
	}

	c, _ := r.Config()
	bc, ok := c.Branches[ref.Name().Short()]
	if !ok {
		return "", nil
	}

	return bc.Merge.Short(), nil
}

func (r GitRepo) CommitMessage() []string {
	// TODO
	panic("not implemented")
}

func (r GitRepo) Branches() []string {
	// TODO
	panic("not implemented")
}

func (r GitRepo) RemoteBranches() []string {
	// TODO
	panic("not implemented")
}
