package bounty

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/bountyeq/web/config"
	"github.com/bountyeq/web/model"
	"github.com/bountyeq/web/sanitize"
	"github.com/bountyeq/web/site"
	"github.com/gin-gonic/gin"
)

// FactionEntry is used for a faction standing
type FactionEntry struct {
	ID    uint32
	Name  string
	Value int32
}

// List returns a list of creatures
func List(c *gin.Context) {
	var err error
	search, ok := c.Get("search")
	fmt.Println("search", search)
	if ok {
		fmt.Println("search", search)
	}
	id := c.Param("id")
	if len(id) > 0 {
		err = getBounty(c, id)
		if err != nil {
			c.Error(err)
			return
		}
		return
	}
	npcs := []model.NpcTypes{}
	err = config.ContentDB.Limit(10).Find(&npcs).Error
	if err != nil {
		c.Error(fmt.Errorf("find NpcTypes: %w", err))
		return
	}
	c.HTML(http.StatusOK, "bounty_list.tmpl", gin.H{
		"Title": "BountyEQ",
		"Npcs":  npcs,
		"Site":  site.Fetch(),
	})
}

func getBounty(c *gin.Context, strID string) error {
	description := ""
	id, err := strconv.Atoi(strID)
	if err != nil {
		return fmt.Errorf("parse id %s: %w", strID, err)
	}
	if id == 0 {
		return fmt.Errorf("invalid id %s", strID)
	}

	npc := &model.NpcTypes{}
	config.ContentDB.Where("id = ?", id).First(npc)
	if err != nil {
		return fmt.Errorf("could not find NPC ID %d: %w", id, err)
	}

	description += fmt.Sprintf("<p>%s is a level %d %s %s that hits for %d to %d damage and has %d hitpoints.</p>", npc.Name, npc.Level, npc.ClassName(), npc.RaceName(), npc.Mindmg, npc.Maxdmg, npc.Hp)

	spells := []*model.SpellsNew{}
	if npc.NpcSpellsID > 0 {
		description += fmt.Sprintf("<p>%s is known to", npc.Name)

		npcSpell := model.NpcSpells{}
		config.ContentDB.Where("id = ?", npc.NpcSpellsID).Find(&npcSpell)
		if npcSpell.AttackProc > 0 {
			spell := &model.SpellsNew{}
			config.ContentDB.Where("id = ?", npcSpell.AttackProc).First(spell)
			description += fmt.Sprintf(" proc <a href='/spell/%d'>%s</a> while attacking, which has a %d%% chance to trigger a %s", spell.ID, spell.Name.String, npcSpell.ProcChance, spell.Summary())
			spell.Name.String = fmt.Sprintf("%s (Attack Proc)", spell.Name.String)

			spells = append(spells, spell)
		}
		if npcSpell.DefensiveProc > 0 {
			spell := &model.SpellsNew{}
			config.ContentDB.Where("id = ?", npcSpell.DefensiveProc).First(spell)
			description += fmt.Sprintf(" proc <a href='/spell/%d'>%s</a> while being attacked, which has a %d%% chance to trigger a %s", spell.ID, spell.Name.String, npcSpell.DprocChance, spell.Summary())
			spell.Name.String = fmt.Sprintf("%s (Defensive Proc)", spell.Name.String)
			spells = append(spells, spell)
		}
		if npcSpell.AttackProc > 0 {
			description += " as well as cast "
		} else {
			description += " cast "
		}
		spellEntries := []model.NpcSpellsEntries{}
		config.ContentDB.Where("npc_spells_id = ?", npc.NpcSpellsID).Find(&spellEntries)
		for i, se := range spellEntries {
			spell := &model.SpellsNew{}
			config.ContentDB.Where("id = ?", se.Spellid).First(spell)
			spells = append(spells, spell)
			if i == len(spellEntries)-1 {
				description += fmt.Sprintf("and <a href='/spell/%d'>%s</a> causing %s", spell.ID, spell.Name.String, spell.Summary())
			} else {
				description += fmt.Sprintf("<a href='/spell/%d'>%s</a> causing %s, ", spell.ID, spell.Name.String, spell.Summary())
			}

		}
		description += ".</p>"
	}
	if npc.NpcSpellsEffectsID > 0 {
		spellEntries := []model.NpcSpellsEffectsEntries{}
		config.ContentDB.Where("npc_spells_effect_id = ?", npc.NpcSpellsEffectsID).Find(&spellEntries)

		for _, se := range spellEntries {
			spell := &model.SpellsNew{}
			config.ContentDB.Where("id = ?", se.SpellEffectID).First(spell)
			spells = append(spells, spell)
		}
	}

	factions := []*FactionEntry{}
	if npc.NpcFactionID > 0 {
		npcFactions := []*model.NpcFactionEntries{}
		config.ContentDB.Where("npc_faction_id = ?", npc.NpcFactionID).Find(&npcFactions)
		for _, faction := range npcFactions {
			npcFaction := model.FactionList{}
			config.ContentDB.Where("id = ?", faction.FactionID).First(&npcFaction)
			f := &FactionEntry{
				ID:    faction.FactionID,
				Name:  npcFaction.Name,
				Value: faction.NpcValue,
			}
			factions = append(factions, f)
		}
	}

	items := []*model.Items{}
	if npc.LoottableID > 0 {
		rare := ""
		lootTables := []*model.LoottableEntries{}
		config.ContentDB.Where("loottable_id = ?", npc.LoottableID).Find(&lootTables)
		for _, lt := range lootTables {
			lootDrops := []*model.LootdropEntries{}
			config.ContentDB.Where("lootdrop_id = ?", lt.LootdropID).Find(&lootDrops)
			for _, ld := range lootDrops {
				item := &model.Items{}
				config.ContentDB.Where("id = ?", ld.ItemID).First(item)
				items = append(items, item)
				if ld.Chance < 0.5 {
					if len(rare) == 0 {
						rare = ", with rare drops including "
					}
					rare += fmt.Sprintf("<a href='/item/%d'>%s</a>, ", item.ID, item.Name)
				}
			}
		}

		description += fmt.Sprintf("<p>%s has a drop pool of %d total items%s.</p>", npc.Name, len(items), rare)
	}

	c.HTML(http.StatusOK, "bounty_view.tmpl", gin.H{
		"Title":       sanitize.CleanName(npc.Name),
		"Npc":         npc,
		"Spells":      spells,
		"Factions":    factions,
		"Items":       items,
		"Description": template.HTML(description),
		"Site":        site.Fetch(),
	})
	fmt.Println("got npc", npc)
	return nil
}
