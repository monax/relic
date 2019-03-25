package main

import (
	"fmt"

	"github.com/monax/relic/v2/project"
)

func main() {
	fmt.Println(project.History.CurrentNotes())
}
