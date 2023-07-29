package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	hyper := []string{"left_control", "left_option"}
    complexModification := ComplexModification{
		Title: "App launcher",
		Rules: []Rule {
				appRule("Visual Studio Code", "v", hyper),
				appRule("Terminal", "t", hyper),
				appRule("Firefox", "m", hyper),
				appRule("Finder", "f", hyper),
				appRule("Microsoft Teams", "c", hyper),
			},
	}
	
	file, _ := json.MarshalIndent(complexModification, "", "  ")
	_ = os.WriteFile("karabiner.json", file, 0644)
}

func appRule(name string, key string, modifiers []string) Rule {
	manipulator := Manipulator {
		Type: "basic",
		From: FromEventDefinition{
			KeyCode: key,
			Modifiers: Modifiers{Mandatory: modifiers},
		},
		To: []ToEventDefinition{app(name)},
	}

	cmd := strings.Join(modifiers[:], "+") + fmt.Sprintf("+%s", key)

	return Rule{
		Description: fmt.Sprintf("Launch %s [%s]", name, cmd),
		Manipulators: []Manipulator{manipulator},
	}
}
