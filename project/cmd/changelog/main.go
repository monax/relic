package main

import (
	"fmt"

	"github.com/monax/relic/project"
)

func main() {
	fmt.Println(project.Project.MustChangelog())
}
