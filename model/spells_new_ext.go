package model

import "fmt"

// Summary returns a one line summary of a spell
func (s *SpellsNew) Summary() string {
	summary := ""
	if s.Effectid1 == 0 && s.EffectBaseValue1 < 0 {
		summary += fmt.Sprintf("%d DD", -s.EffectBaseValue1)
	}
	if s.Effectid1 == 83 {
		summary += "Teleport/Banish"
	}
	if s.Effectid1 == 0 && s.EffectBaseValue1 > 0 {
		summary += fmt.Sprintf("%d Heal", s.EffectBaseValue1)
	}
	return summary
}
