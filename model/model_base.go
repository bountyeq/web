package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["aa_ability"] = aa_abilityTableInfo
	tables["aa_rank_effects"] = aa_rank_effectsTableInfo
	tables["aa_rank_prereqs"] = aa_rank_prereqsTableInfo
	tables["aa_ranks"] = aa_ranksTableInfo
	tables["account"] = accountTableInfo
	tables["account_flags"] = account_flagsTableInfo
	tables["account_ip"] = account_ipTableInfo
	tables["account_rewards"] = account_rewardsTableInfo
	tables["adventure_details"] = adventure_detailsTableInfo
	tables["adventure_members"] = adventure_membersTableInfo
	tables["adventure_stats"] = adventure_statsTableInfo
	tables["adventure_template"] = adventure_templateTableInfo
	tables["adventure_template_entry"] = adventure_template_entryTableInfo
	tables["adventure_template_entry_flavor"] = adventure_template_entry_flavorTableInfo
	tables["alternate_currency"] = alternate_currencyTableInfo
	tables["auras"] = aurasTableInfo
	tables["banned_ips"] = banned_ipsTableInfo
	tables["base_data"] = base_dataTableInfo
	tables["blocked_spells"] = blocked_spellsTableInfo
	tables["books"] = booksTableInfo
	tables["bug_reports"] = bug_reportsTableInfo
	tables["bugs"] = bugsTableInfo
	tables["buyer"] = buyerTableInfo
	tables["char_create_combinations"] = char_create_combinationsTableInfo
	tables["char_create_point_allocations"] = char_create_point_allocationsTableInfo
	tables["char_recipe_list"] = char_recipe_listTableInfo
	tables["character_activities"] = character_activitiesTableInfo
	tables["character_alt_currency"] = character_alt_currencyTableInfo
	tables["character_alternate_abilities"] = character_alternate_abilitiesTableInfo
	tables["character_auras"] = character_aurasTableInfo
	tables["character_bandolier"] = character_bandolierTableInfo
	tables["character_bind"] = character_bindTableInfo
	tables["character_buffs"] = character_buffsTableInfo
	tables["character_corpse_items"] = character_corpse_itemsTableInfo
	tables["character_corpses"] = character_corpsesTableInfo
	tables["character_currency"] = character_currencyTableInfo
	tables["character_data"] = character_dataTableInfo
	tables["character_disciplines"] = character_disciplinesTableInfo
	tables["character_enabledtasks"] = character_enabledtasksTableInfo
	tables["character_inspect_messages"] = character_inspect_messagesTableInfo
	tables["character_item_recast"] = character_item_recastTableInfo
	tables["character_languages"] = character_languagesTableInfo
	tables["character_leadership_abilities"] = character_leadership_abilitiesTableInfo
	tables["character_material"] = character_materialTableInfo
	tables["character_memmed_spells"] = character_memmed_spellsTableInfo
	tables["character_pet_buffs"] = character_pet_buffsTableInfo
	tables["character_pet_info"] = character_pet_infoTableInfo
	tables["character_pet_inventory"] = character_pet_inventoryTableInfo
	tables["character_potionbelt"] = character_potionbeltTableInfo
	tables["character_skills"] = character_skillsTableInfo
	tables["character_spells"] = character_spellsTableInfo
	tables["character_tasks"] = character_tasksTableInfo
	tables["character_tribute"] = character_tributeTableInfo
	tables["chatchannels"] = chatchannelsTableInfo
	tables["command_settings"] = command_settingsTableInfo
	tables["completed_tasks"] = completed_tasksTableInfo
	tables["content_flags"] = content_flagsTableInfo
	tables["damageshieldtypes"] = damageshieldtypesTableInfo
	tables["data_buckets"] = data_bucketsTableInfo
	tables["db_str"] = db_strTableInfo
	tables["db_version"] = db_versionTableInfo
	tables["discovered_items"] = discovered_itemsTableInfo
	tables["doors"] = doorsTableInfo
	tables["eqtime"] = eqtimeTableInfo
	tables["eventlog"] = eventlogTableInfo
	tables["faction_base_data"] = faction_base_dataTableInfo
	tables["faction_list"] = faction_listTableInfo
	tables["faction_list_mod"] = faction_list_modTableInfo
	tables["faction_values"] = faction_valuesTableInfo
	tables["fishing"] = fishingTableInfo
	tables["forage"] = forageTableInfo
	tables["friends"] = friendsTableInfo
	tables["global_loot"] = global_lootTableInfo
	tables["gm_ips"] = gm_ipsTableInfo
	tables["goallists"] = goallistsTableInfo
	tables["graveyard"] = graveyardTableInfo
	tables["grid"] = gridTableInfo
	tables["grid_entries"] = grid_entriesTableInfo
	tables["ground_spawns"] = ground_spawnsTableInfo
	tables["group_id"] = group_idTableInfo
	tables["group_leaders"] = group_leadersTableInfo
	tables["guild_bank"] = guild_bankTableInfo
	tables["guild_members"] = guild_membersTableInfo
	tables["guild_ranks"] = guild_ranksTableInfo
	tables["guild_relations"] = guild_relationsTableInfo
	tables["guilds"] = guildsTableInfo
	tables["hackers"] = hackersTableInfo
	tables["horses"] = horsesTableInfo
	tables["instance_list"] = instance_listTableInfo
	tables["instance_list_player"] = instance_list_playerTableInfo
	tables["inventory"] = inventoryTableInfo
	tables["inventory_snapshots"] = inventory_snapshotsTableInfo
	tables["inventory_versions"] = inventory_versionsTableInfo
	tables["ip_exemptions"] = ip_exemptionsTableInfo
	tables["item_tick"] = item_tickTableInfo
	tables["items"] = itemsTableInfo
	tables["keyring"] = keyringTableInfo
	tables["launcher"] = launcherTableInfo
	tables["launcher_zones"] = launcher_zonesTableInfo
	tables["ldon_trap_entries"] = ldon_trap_entriesTableInfo
	tables["ldon_trap_templates"] = ldon_trap_templatesTableInfo
	tables["level_exp_mods"] = level_exp_modsTableInfo
	tables["lfguild"] = lfguildTableInfo
	tables["login_accounts"] = login_accountsTableInfo
	tables["login_api_tokens"] = login_api_tokensTableInfo
	tables["login_server_admins"] = login_server_adminsTableInfo
	tables["login_server_list_types"] = login_server_list_typesTableInfo
	tables["login_world_servers"] = login_world_serversTableInfo
	tables["logsys_categories"] = logsys_categoriesTableInfo
	tables["lootdrop"] = lootdropTableInfo
	tables["lootdrop_entries"] = lootdrop_entriesTableInfo
	tables["loottable"] = loottableTableInfo
	tables["loottable_entries"] = loottable_entriesTableInfo
	tables["mail"] = mailTableInfo
	tables["merchantlist"] = merchantlistTableInfo
	tables["merchantlist_temp"] = merchantlist_tempTableInfo
	tables["name_filter"] = name_filterTableInfo
	tables["npc_emotes"] = npc_emotesTableInfo
	tables["npc_faction"] = npc_factionTableInfo
	tables["npc_faction_entries"] = npc_faction_entriesTableInfo
	tables["npc_scale_global_base"] = npc_scale_global_baseTableInfo
	tables["npc_spells"] = npc_spellsTableInfo
	tables["npc_spells_effects"] = npc_spells_effectsTableInfo
	tables["npc_spells_effects_entries"] = npc_spells_effects_entriesTableInfo
	tables["npc_spells_entries"] = npc_spells_entriesTableInfo
	tables["npc_types"] = npc_typesTableInfo
	tables["npc_types_tint"] = npc_types_tintTableInfo
	tables["object"] = objectTableInfo
	tables["object_contents"] = object_contentsTableInfo
	tables["perl_event_export_settings"] = perl_event_export_settingsTableInfo
	tables["petitions"] = petitionsTableInfo
	tables["pets"] = petsTableInfo
	tables["pets_equipmentset"] = pets_equipmentsetTableInfo
	tables["pets_equipmentset_entries"] = pets_equipmentset_entriesTableInfo
	tables["player_titlesets"] = player_titlesetsTableInfo
	tables["profanity_list"] = profanity_listTableInfo
	tables["proximities"] = proximitiesTableInfo
	tables["qs_merchant_transaction_record"] = qs_merchant_transaction_recordTableInfo
	tables["qs_merchant_transaction_record_entries"] = qs_merchant_transaction_record_entriesTableInfo
	tables["qs_player_aa_rate_hourly"] = qs_player_aa_rate_hourlyTableInfo
	tables["qs_player_delete_record"] = qs_player_delete_recordTableInfo
	tables["qs_player_delete_record_entries"] = qs_player_delete_record_entriesTableInfo
	tables["qs_player_events"] = qs_player_eventsTableInfo
	tables["qs_player_handin_record"] = qs_player_handin_recordTableInfo
	tables["qs_player_handin_record_entries"] = qs_player_handin_record_entriesTableInfo
	tables["qs_player_move_record"] = qs_player_move_recordTableInfo
	tables["qs_player_move_record_entries"] = qs_player_move_record_entriesTableInfo
	tables["qs_player_npc_kill_record"] = qs_player_npc_kill_recordTableInfo
	tables["qs_player_npc_kill_record_entries"] = qs_player_npc_kill_record_entriesTableInfo
	tables["qs_player_speech"] = qs_player_speechTableInfo
	tables["qs_player_trade_record"] = qs_player_trade_recordTableInfo
	tables["qs_player_trade_record_entries"] = qs_player_trade_record_entriesTableInfo
	tables["quest_globals"] = quest_globalsTableInfo
	tables["raid_details"] = raid_detailsTableInfo
	tables["raid_leaders"] = raid_leadersTableInfo
	tables["raid_members"] = raid_membersTableInfo
	tables["reports"] = reportsTableInfo
	tables["respawn_times"] = respawn_timesTableInfo
	tables["rule_sets"] = rule_setsTableInfo
	tables["rule_values"] = rule_valuesTableInfo
	tables["saylink"] = saylinkTableInfo
	tables["sharedbank"] = sharedbankTableInfo
	tables["skill_caps"] = skill_capsTableInfo
	tables["spawn2"] = spawn2TableInfo
	tables["spawn_condition_values"] = spawn_condition_valuesTableInfo
	tables["spawn_conditions"] = spawn_conditionsTableInfo
	tables["spawn_events"] = spawn_eventsTableInfo
	tables["spawnentry"] = spawnentryTableInfo
	tables["spawngroup"] = spawngroupTableInfo
	tables["spell_buckets"] = spell_bucketsTableInfo
	tables["spell_globals"] = spell_globalsTableInfo
	tables["spells_new"] = spells_newTableInfo
	tables["start_zones"] = start_zonesTableInfo
	tables["starting_items"] = starting_itemsTableInfo
	tables["task_activities"] = task_activitiesTableInfo
	tables["tasks"] = tasksTableInfo
	tables["tasksets"] = tasksetsTableInfo
	tables["timers"] = timersTableInfo
	tables["titles"] = titlesTableInfo
	tables["trader"] = traderTableInfo
	tables["tradeskill_recipe"] = tradeskill_recipeTableInfo
	tables["tradeskill_recipe_entries"] = tradeskill_recipe_entriesTableInfo
	tables["traps"] = trapsTableInfo
	tables["tribute_levels"] = tribute_levelsTableInfo
	tables["tributes"] = tributesTableInfo
	tables["variables"] = variablesTableInfo
	tables["veteran_reward_templates"] = veteran_reward_templatesTableInfo
	tables["zone"] = zoneTableInfo
	tables["zone_flags"] = zone_flagsTableInfo
	tables["zone_points"] = zone_pointsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
