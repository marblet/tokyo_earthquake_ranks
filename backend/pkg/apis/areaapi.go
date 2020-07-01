package apis

import (
	"encoding/json"
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type TownInfo struct {
	ID 				uint	`json:"id"`
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

func GetTowns(w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	address := v.Get("address")
	page, _ := strconv.Atoi(v.Get("page"))
	displaynum, _ := strconv.Atoi(v.Get("displaynum"))
	if page < 1 || displaynum < 1 {
		w.WriteJson(Result{MatchedNum: 0, MatchedTowns: nil})
		return
	}
	foundTowns := filter(address)
	foundNum := len(foundTowns)
	left := int(math.Min(float64(foundNum), float64((page - 1) * displaynum)))
	right := int(math.Min(float64(foundNum), float64(page * displaynum)))
	foundTowns = foundTowns[left: right]
	result := Result{MatchedNum: uint(foundNum), MatchedTowns: foundTowns}
	w.WriteJson(result)
}

func filter(address string) []SimpleTownInfo {
	var ret []SimpleTownInfo
	for _, v := range simpleTownInfos {
		if strings.Contains(v.Municipality + v.TownName, address) {
			ret = append(ret, v)
		}
	}
	return ret
}

func LoadTownFile() {
	raw, err := ioutil.ReadFile("assets/all2.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &townInfos)
	json.Unmarshal(raw, &simpleTownInfos)
}