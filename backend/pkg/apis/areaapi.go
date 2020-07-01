package apis

import (
	"encoding/json"
	"io/ioutil"
)

type TownInfo struct {
	ID 				uint	`json:"id"`
	AreaCode		string	`json:"area_code"`
	Municipality	string	`json:"municipality"`
	TownName		string	`json:"town_name"`
	BaseClass		string	`json:"base_class"`
	CollapseHa		float64	`json:"collapse_ha"`
	CollapseOrder	uint	`json:"collapse_order"`
	CollapseRank	string	`json:"collapse_rank"`
	FireHa			float64	`json:"fire_ha"`
	FireOrder		uint	`json:"fire_order"`
	FireRank		string	`json:"fire_rank"`
	Difficulty		float64	`json:"difficulty"`
	DifficultyOrder	uint	`json:"difficulty_order"`
	DifficultyRank	string	`json:"difficulty_rank"`
	TotalHa			float64	`json:"total_ha"`
	TotalOrder		uint	`json:"total_order"`
	TotalRank		string	`json:"total_rank"`
}

type SimpleTownInfo struct {
	ID 				uint	`json:"id"`
	AreaCode		string	`json:"area_code"`
	Municipality	string	`json:"municipality"`
	TownName		string	`json:"town_name"`
	CollapseRank	string	`json:"collapse_rank"`
	FireRank		string	`json:"fire_rank"`
	DifficultyRank	string	`json:"difficulty_rank"`
	TotalRank		string	`json:"total_rank"`
}

type Result struct {
	MatchedNum		uint				`json:"matched_num"`
	MatchedTowns	[]SimpleTownInfo	`json:"matched_towns"`
}

var townInfos []TownInfo
var simpleTownInfos []SimpleTownInfo


func LoadTownFile () {
	raw, err := ioutil.ReadFile("assets/all2.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &townInfos)
	json.Unmarshal(raw, &simpleTownInfos)
}