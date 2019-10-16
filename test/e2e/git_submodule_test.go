package e2e

import (
	"testing"

	"github.com/argoproj/argo-cd/test/e2e/fixture"
	v1 "k8s.io/api/core/v1"

	. "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	. "github.com/argoproj/argo-cd/test/e2e/fixture/app"
)

func TestGitSubmoduleSSHSupport(t *testing.T) {
	Given(t).
		RepoURLType(fixture.RepoURLTypeSSHSubmoduleParent).
		Path("submodule").
		Recurse().
		CustomSSHKnownHostsAdded().
		SubmoduleSSHRepoURLAdded().
		When().
		CreateFromFile(func(app *Application){}).
		Sync().
		Then().
		Expect(SyncStatusIs(SyncStatusCodeSynced)).
		Expect(Pod(func(p v1.Pod) bool { return p.Name == "pod-in-submodule" }))
}

func TestGitSubmoduleHTTPSSupport(t *testing.T) {
	Given(t).
		RepoURLType(fixture.RepoURLTypeHTTPSSubmoduleParent).
		Path("submodule").
		Recurse().
		CustomCACertAdded().
		SubmoduleHTTPSRepoURLAdded().
		When().
		CreateFromFile(func(app *Application){}).
		Sync().
		Then().
		Expect(SyncStatusIs(SyncStatusCodeSynced)).
		Expect(Pod(func(p v1.Pod) bool { return p.Name == "pod-in-submodule" }))
}
