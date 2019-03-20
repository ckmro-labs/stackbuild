package docker

import (
	"strings"

	"github.com/docker/distribution/reference"
)

func parseImage(s string) (canonical, domain string, latest bool, err error) {
	named, err := reference.ParseNormalizedNamed(s)
	if err != nil {
		return
	}
	named = reference.TagNameOnly(named)

	return named.String(),
		reference.Domain(named),
		strings.HasSuffix(named.String(), ":latest"),
		nil
}
