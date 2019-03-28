package repo

import (
	"github.com/drone/go-scm/scm"
	"github.com/laidingqing/stackbuild/core"
)

// convertRepository convert remote repository info to local repo.
func convertRepository(src *scm.Repository) *core.Repository {
	return &core.Repository{
		UID:       src.ID,
		Namespace: src.Namespace,
		Name:      src.Name,
		Slug:      scm.Join(src.Namespace, src.Name),
		HTTPURL:   src.Clone,
		SSHURL:    src.CloneSSH,
		Link:      src.Link,
		Private:   src.Private,
		Branch:    src.Branch,
	}
}
