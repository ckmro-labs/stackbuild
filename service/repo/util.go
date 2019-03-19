package repo

import (
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/laidingqing/stackbuild/core"
)

func userToken(user *core.User, provider core.VcsProvider) (string, string, error) {
	if len(user.Authentications) == 0 {
		return "", "", fmt.Errorf("user %v no authentication", user.Login)
	}

	for _, auth := range user.Authentications {
		if auth.AuthName == provider {
			return auth.Token, auth.Refresh, nil
		}
	}
	return "", "", fmt.Errorf("no found token: %s", provider)
}

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
