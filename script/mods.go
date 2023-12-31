package script

var adminMods = []string{
	"production_efficiency",
	"cossacks_privilege_slots",
	"nobles_influence_modifier",
	"country_admin_power",
	"mil_tech_cost_modifier",
	"power_projection_from_insults",
	"monarch_diplomatic_power",
	"reelection_cost",
	"legitimacy",
	"overextension_impact_modifier",
	"administrative_efficiency",
	"same_culture_advisor_cost",
	"monthly_gold_inflation_modifier",
	"nobles_loyalty_modifier",
	"promote_culture_cost",
	"inflation_reduction",
	"global_unrest",
	"monthly_splendor",
	"idea_cost",
	"max_revolutionary_zeal",
	"female_advisor_chance",
	"meritocracy",
	"advisor_pool",
	"development_cost_in_primary_culture",
	"dhimmi_loyalty_modifier",
	"free_leader_pool",
	"devotion",
	"migration_cost",
	"burghers_influence_modifier",
	"reform_progress_growth",
	"prestige",
	"embracement_cost",
	"interest",
	"possible_policy",
	"global_allowed_num_of_buildings",
	"expand_administration_cost",
	"governing_capacity_modifier",
	"loyalty_change_on_revoked",
	"monarch_military_power",
	"num_of_parliament_issues",
	"imperial_authority_value",
	"papal_influence",
	"great_project_upgrade_cost",
	"min_autonomy_in_territories",
	"gold_depletion_chance_modifier",
	"development_cost",
	"harsh_treatment_cost",
	"technology_cost",
	"prestige_decay",
	"years_of_nationalism",
	"num_accepted_cultures",
	"yearly_corruption",
	"global_institution_spread",
	"heir_chance",
	"dip_tech_cost_modifier",
	"core_creation",
	"global_tax_modifier",
	"all_estate_loyalty_equilibrium",
	"build_cost",
	"state_maintenance_modifier",
	"adm_tech_cost_modifier",
	"free_adm_policy",
	"global_autonomy",
	"stability_cost_modifier",
	"yearly_absolutism",
	"innovativeness_gain",
	"imperial_mandate",
	"burghers_loyalty_modifier",
	"yearly_government_power",
	"horde_unity",
	"advisor_cost",
	"imperial_authority",
	"republican_tradition",
	"build_time",
}

var tradeMods = []string{
	"caravan_power",
	"global_prov_trade_power_modifier",
	"justify_trade_conflict_cost",
	"embargo_efficiency",
	"merchants",
	"max_absolutism",
	"trade_steering",
	"center_of_trade_upgrade_cost",
	"global_foreign_trade_power",
	"global_trade_power",
	"global_ship_trade_power",
	"ship_power_propagation",
	"global_trade_goods_size_modifier",
	"global_own_trade_power",
	"privateer_efficiency",
	"trade_efficiency",
	"trade_company_investment_cost",
	"trade_range_modifier",
}

var colonialMods = []string{
	"colonists",
	"native_uprising_chance",
	"may_explore",
	"global_colonial_growth",
	"global_tariffs",
	"idea_claim_colonies",
	"native_assimilation",
	"may_establish_frontier",
	"range",
	"auto_explore_adjacent_to_colony",
}

var religiousMods = []string{
	// Missionary
	"missionary_maintenance_cost",
	"global_missionary_strength",
	"missionaries",
	"global_heretic_missionary_strength",

	//Tolerance & Unity
	"tolerance_own",
	"tolerance_heathen",
	"tolerance_heretic",
	"tolerance_of_heathens_capacity",
	"tolerance_of_heretics_capacity",
	"religious_unity",

	// Church Specifc
	"monthly_piety",
	"church_power_modifier",
	"curia_powers_cost",
	"church_loyalty_modifier",
	"yearly_patriarch_authority",
	"monthly_fervor_increase",
	"monthly_piety_accelerator",

	// Misc
	"no_religion_penalty",
	"enforce_religion_cost",
	"manpower_in_true_faith_provinces",
}

var navalMods = []string{
	// Ship modifiers
	"flagship_cost",
	"ship_durability",
	"galley_cost",
	"transport_cost",
	"heavy_ship_power",
	"heavy_ship_cost",
	"galley_power",
	"light_ship_power",
	"global_ship_cost",
	"light_ship_cost",

	// Leader & combat
	"leader_naval_shock",
	"leader_naval_fire",
	"leader_naval_manuever",
	"disengagement_chance",
	"capture_ship_chance",
	"global_naval_engagement_modifier",

	// Tradition & morale
	"naval_morale",
	"navy_tradition",
	"naval_tradition_from_battle",
	"navy_tradition_decay",
	"recover_navy_morale_speed",
	"sunk_ship_morale_hit_recieved",

	// Sailors & FL
	"global_sailors_modifier",
	"allowed_marine_fraction",
	"has_geobukseon",
	"naval_forcelimit_modifier",
	"siege_blockade_progress",
	"sailors_recovery_speed",
	"naval_maintenance_modifier",
	"sailor_maintenance_modifer",

	// Misc
	"blockade_efficiency",
	"global_ship_recruit_speed",
	"movement_speed_in_fleet_modifier",
	"naval_attrition",
	"own_coast_naval_combat_bonus",
	"may_perform_slave_raid",
	"may_perform_slave_raid_on_same_religion",
}

var armyMods = []string{
	// Combat
	"discipline",
	"cavalry_power",
	"artillery_power",
	"cavalry_cost",
	"cavalry_flanking",
	"cavalry_fire",
	"fire_damage_received",
	"artillery_fire",
	"artillery_shock",
	"shock_damage",
	"fire_damage",
	"infantry_power",
	"morale_damage_received",
	"infantry_cost",
	"backrow_artillery_damage",
	"cav_to_inf_ratio",
	"shock_damage_received",
	"morale_damage",
	"land_morale",
	"recover_army_morale_speed",

	// Leaders
	"general_cost",
	"leader_cost",
	"leader_land_fire",
	"leader_siege",
	"leader_land_shock",
	"leader_land_manuever",

	// Merc & Special
	"mercenary_manpower",
	"merc_maintenance_modifier",
	"amount_of_banners",
	"mercenary_cost",
	"mercenary_discipline",
	"can_recruit_hussars",
	"allowed_samurai_fraction",
	"has_samurai",

	// Tradition & Professionalism
	"yearly_army_professionalism",
	"army_tradition",
	"army_tradition_decay",
	"army_tradition_from_battle",
	"drill_gain_modifier",

	// Forts
	"rival_border_fort_maintenance",
	"fort_maintenance_modifier",
	"defensiveness",
	"garrison_size",
	"siege_ability",
	"garrison_damage",
	"global_garrison_growth",

	// Manpower, army maintanance and FL
	"vassal_forcelimit_bonus",
	"land_forcelimit_modifier",
	"land_maintenance_modifier",
	"global_regiment_cost",
	"global_supply_limit_modifier",
	"global_manpower_modifier",
	"reinforce_speed",
	"global_regiment_recruit_speed",
	"reinforce_cost_modifier",

	// Misc
	"hostile_attrition",
	"max_hostile_attrition",
	"war_taxes_cost_modifier",
	"free_mil_policy",
	"may_recruit_female_generals",
	"war_exhaustion_cost",
	"war_exhaustion",
	"loot_amount",
	"movement_speed",
	"prestige_from_land",
}

var diploMods = []string{
	"reduced_liberty_desire",
	"spy_offence",
	"ae_impact",
	"reduced_liberty_desire_on_same_continent",
	"liberty_desire_from_subject_development",
	"envoy_travel_time",
	"accept_vassalization_reasons",
	"diplomatic_upkeep",
	"diplomats",
	"improve_relation_modifier",
	"free_dip_policy",
	"stability_cost_to_declare_war",
	"fabricate_claims_cost",
	"diplomatic_annexation_cost",
	"vassal_income",
	"global_spy_defence",
	"possible_dip_policy",
	"unjustified_demands",
	"warscore_cost_vs_other_religion",
	"rebel_support_efficiency",
	"years_to_integrate_personal_union",
	"diplomatic_reputation",
	"province_warscore_cost",
}
