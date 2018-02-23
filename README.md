# Relic
Relic is a small library to help with versioning your projects by storing release metadata 
and versions as code.

## Purpose
Relic allows you define your project version history in a declarative style by
defining a `History` object somewhere in your project whose methods allow you to
add releases defined by a version and release notes. It ensures you releases form
a valid monotonic progression of semantic version compliant tagged releases. 

From this is can generate the current version and a complete changelog using this information.

## Usage
```go

```
See Relic's own [`project` package](project/releases.go) and [Makefile](Makefile) for example
usage within a project.

## Dependencies
Go standard libary and tooling plus Make and Bash for builds (but not required)