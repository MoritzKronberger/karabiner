package main

type Modifiers struct {
	Mandatory []string `json:"mandatory,omitempty"`
	Optional []string `json:"optional,omitempty"`
}

type FromEventDefinition struct {
	KeyCode string `json:"key_code,omitempty"`
	Modifiers Modifiers `json:"modifiers,omitempty"`
}

type ToEventDefinition struct {
	KeyCode string `json:"key_code,omitempty"`
	Modifiers []string `json:"modifiers,omitempty"`
	ShellCommand string `json:"shell_command,omitempty"`
}

type Manipulator struct {
	Type string `json:"type,omitempty"`
	From FromEventDefinition `json:"from,omitempty"`
	To []ToEventDefinition `json:"to,omitempty"`
}

type Rule struct {
	Description string `json:"description,omitempty"`
	Manipulators []Manipulator `json:"manipulators,omitempty"`
}

type ComplexModification struct {
	Title string `json:"title,omitempty"`
	Rules []Rule `json:"rules,omitempty"`
}
