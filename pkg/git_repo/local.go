package git_repo

import (
	"fmt"
)

type Local struct {
	Base
	WorkTree string
}

func (repo *Local) HeadCommit() (string, error) {
	commit, err := repo.getHeadCommitForRepo(repo.WorkTree)

	if err == nil {
		fmt.Printf("Using commit `%s` of repo `%s`\n", commit, repo.String())
	}

	return commit, err
}

func (repo *Local) CreatePatch(opts PatchOptions) (Patch, error) {
	return repo.createPatch(repo.WorkTree, opts)
}

func (repo *Local) CreateArchive(opts ArchiveOptions) (Archive, error) {
	return repo.createArchive(repo.WorkTree, opts)
}
