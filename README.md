# Relic
Relic is a library to help with versioning your projects by storing release metadata 
and versions as code.

## Purpose
Relic allows you define your project version history in a declarative style by
defining a `History` object somewhere in your project whose methods allow you to
declare releases defined by a version number and release note. It ensures your releases
have monotonically increasing unique versions. 

Relic can generate the current version and a complete changelog using this information.

By keeping the changelog with the version they are synchronised and you are reminded to produce 
the changelog.

## Usage
```go
// Add file to your project in which to record your projects revision history
package project

import (
	"fmt"
	"text/template"
	"github.com/monax/relic"
)

// Create a global variable in which to store your project history.
// MustDeclareReleases allows you to declare your releases by specifying a version and release note
// for each release. To add a new release just insert it at the top.
var history relic.ImmutableHistory = relic.NewHistory("Relic").
	MustDeclareReleases(
		"1.1.0",
		`Add ImmutableHistory and tweak suggested usage docs`,
		"1.0.1",
		`Documentation fixes and typos`,
		"1.0.0",
		`Minor improvements:
- Rename RecordReleases to DeclareReleases (breaking API change)
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

func PrintReleaseInfo() {
	// Print the current version
	fmt.Printf("%s (Version: %v)\n", history.Project(), history.CurrentVersion().String())
	// Print the complete changelog 
	fmt.Println(history.Changelog())
	// Get specific release
	release, err := history.Release("0.0.1")
	if err != nil {
		panic(err)
	}
	// Print major version of release
	fmt.Printf("Release Major version: %v", release.Version.Major())
}

// You can also define histories with a custom template
var ProjectWithTemplate = relic.NewHistory("Test Project").
		WithChangelogTemplate(template.Must(template.New("tests").
			Parse("{{range .Releases}}{{$.Name}} (v{{.Version}}): {{.Notes}}\n{{end}}"))).
		MustDeclareReleases(
			"0.1.0",
			"Basic functionality",
			"0.0.2",
			"Build scripts",
			"0.0.1",
			"Proof of concept",
		)

```

See Relic's own [`project` package](project/releases.go) and [Makefile](Makefile) for suggested usage within a project.

## Dependencies
Go standard library and tooling plus Make and Bash for builds (but not required).