package sanitize

import "fmt"

// ClassIcon returns a class's icon
func ClassIcon(classID uint32) string {
	switch classID {
	case 1:
		return "xa-shield" //warrior
	case 2:
		return "xa-ankh" //cleric
	case 3:
		return "xa-fireball-sword" //paladin
	case 4:
		return "xa-arrow-cluster" //ranger
	case 5:
		return "xa-bat-sword" //shd
	case 6:
		return "xa-leaf" //druid
	case 7:
		return "xa-hand-emblem" //Monk
	case 8:
		return "xa-ocarina" //Bard
	case 9:
		return "xa-hood" //rogue
	case 10:
		return "xa-incense" //shaman
	case 11:
		return "xa-skull" //necro
	case 12:
		return "xa-fire" //wiz
	case 13:
		return "xa-burning-book" //magician
	case 14:
		return "xa-crystal-ball" //enchanter
	case 15:
		return "xa-pawprint" //beastlord
	case 16:
		return "xa-axe" //ber
	case 41:
		return "xa-capitol" //shopkeeper
	}
	return "xa-help"
}

// ClassName returns a class's name
func ClassName(classID uint32) string {

	switch classID {
	case 1:
		return "Warrior"
	case 2:
		return "Cleric"
	case 3:
		return "Paladin"
	case 4:
		return "Ranger"
	case 5:
		return "Shadowknight"
	case 6:
		return "Druid"
	case 7:
		return "Monk"
	case 8:
		return "Bard"
	case 9:
		return "Rogue"
	case 10:
		return "Shaman"
	case 11:
		return "Necromancer"
	case 12:
		return "Wizard"
	case 13:
		return "Magician"
	case 14:
		return "Enchanter"
	case 15:
		return "Beastlord"
	case 16:
		return "Berserker"
	case 20:
		return "GM Warrior"
	case 21:
		return "GM Cleric"
	case 22:
		return "GM Paladin"
	case 23:
		return "GM Ranger"
	case 24:
		return "GM Shadow Knight"
	case 25:
		return "GM Druid"
	case 26:
		return "GM Monk"
	case 27:
		return "GM Bard"
	case 28:
		return "GM Rogue"
	case 29:
		return "GM Shaman"
	case 30:
		return "GM Necromancer"
	case 31:
		return "GM Wizard"
	case 32:
		return "GM Magician"
	case 33:
		return "GM Enchanter"
	case 34:
		return "GM Beastlord"
	case 35:
		return "GM Berserker"
	case 40:
		return "Banker"
	case 41:
		return "Shopkeeper"
	case 59:
		return "Discord Merchant"
	case 60:
		return "Adventure Recruiter"
	case 61:
		return "Adventure Merchant"
	case 63:
		return "Tribute Master"
	case 64:
		return "Guild Tribute Master?"
	case 66:
		return "Guild Bank"
	case 67:
		return "Radiant Crystal Merchant"
	case 68:
		return "Ebon Crystal Merchant"
	case 69:
		return "Fellowships"
	case 70:
		return "Alternate Currency Merchant"
	case 71:
		return "Mercenary Merchant"
	}
	return fmt.Sprintf("Unknown (%d)", classID)
}
