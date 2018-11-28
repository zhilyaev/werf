package build

import (
	"fmt"

	"github.com/flant/dapp/pkg/build/stage"
)

type LegacyStage interface {
	GetPrevStage() LegacyStage
	GetImage() stage.Image
	LayerCommit(gitArtifact *GitArtifact) (string, error)
}

type StubStage struct {
	LayerCommitMap map[string]string
	PrevStage      *StubStage
	Image          *stage.StubImage
}

func (stage *StubStage) GetPrevStage() LegacyStage {
	return stage.PrevStage
}

func (stage *StubStage) GetImage() stage.Image {
	return stage.Image
}

func (stage *StubStage) LayerCommit(gitArtifact *GitArtifact) (string, error) {
	if commit, hasKey := stage.LayerCommitMap[gitArtifact.GetParamshash()]; hasKey {
		return commit, nil
	}

	panic(fmt.Errorf("assertion failed: StubStage layer commit should be present for git `%s`", gitArtifact.GetParamshash()))
}