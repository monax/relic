# Relic
Relic is a small library to help with versioning your projects by storing release metadata 
and versions as code.

## Purpose
Relic allows you define your project version history in a declarative style by
defining a `History` object somewhere in your project whose methods allow you to
add releases defined by a version and release notes. It ensures you releases form
a valid monotonic progression of semantic version compliant tagged releases. 

From this is can generate the current version and a complete changelog using this information.

By keeping the changelog with the version they are syncrhonised and you are reminded to produce 
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
var Project = relic.NewHistory("Relic").
	MustDeclareReleases(
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
	fmt.Printf("%s (Version: %v)\n", Project.Name, Project.CurrentVersion().String())
	// Print the complete changelog 
	fmt.Println(Project.Changelog())
	// Get specific release
	release, err := Project.Release("0.0.1")
	if err != nil {
		panic(err)
	}
	// Print major version of release
	fmt.Printf("Release Major version: %v", release.Version.Major())
}

//
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