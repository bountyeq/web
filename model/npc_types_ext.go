package model

import (
	"github.com/bountyeq/web/sanitize"
)

// Description returns a description
func (n *NpcTypes) Description() string {
	return "Test"
}

// CleanName returns a clean version of a NPC's name
func (n *NpcTypes) CleanName() string {
	return sanitize.CleanName(n.Name)
}

// ClassName returns a class's name
func (n *NpcTypes) ClassName() string {
	return sanitize.ClassName(n.Class)
}

// ClassIcon returns a class's icon
func (n *NpcTypes) ClassIcon() string {
	return sanitize.ClassIcon(n.Class)
}

// RaceName returns a race name
func (n *NpcTypes) RaceName() string {
	return sanitize.RaceName(n.Race)
}

// RaceIcon returns a race's icon
func (n *NpcTypes) RaceIcon() string {
	return sanitize.RaceIcon(n.Race)
}
