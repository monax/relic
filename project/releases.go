package project

import "github.com/monax/relic"

var Project = relic.NewHistory("Relic").
	MustRecordReleases(
		"0.0.1",
		`First release of Relic extracted from various initial projects, it can:
- Generate changelogs
- Print the current version
- Ensure valid semantic version numbers
`,
	)
