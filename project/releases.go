package project

import "github.com/monax/relic"

var History relic.ImmutableHistory = relic.NewHistory("Relic").
	MustDeclareReleases(
		"1.1.0",
		`Add ImmutableHistory and tweak suggested usage docs`,
		"1.0.1",
		`Documentation fixes and typos`,
		"1.0.0",
		`Minor improvements:
- Rename DeclareReleases to DeclareReleases (breaking API change)
- Add sample snippet to readme
- Sign version tags
`,
		"0.0.1",
		`First release of Relic extracted from various initial projects, it can:
- Generate changelogs
- Print the current version
- Ensure valid semantic version numbers
`,
	)
