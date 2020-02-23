package labelmaker

import (
	"strings"
)

// ParseRepository takes a repository and returns the owner and name.
func ParseRepository(repo string) (owner, name string) {
	splitRepo := strings.SplitN(repo, "/", -1)
	if len(splitRepo) == 2 {
		return splitRepo[0], splitRepo[1]

	}
	return "", ""
}
