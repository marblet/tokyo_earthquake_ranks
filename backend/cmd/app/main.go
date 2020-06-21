package main

import (
	"encoding/json"
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type AreaInfo struct {
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

type Result struct {
	MatchedNum		uint		`json:"matched_num"`
	MatchedAreas	[]AreaInfo	`json:"matched_areas"`
}

var areaData []AreaInfo

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/areas", getAreas),
		rest.Get("/api/towninfo", getTownInfo),
	)
	if err != nil {
		log.Fatal(err)
	}
	loadAreaFile()
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getTownInfo(w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	id, _ := strconv.Atoi(v.Get("id"))
	if id < 1 || id > len(areaData) {
		w.WriteJson(map[string]string{"Error": "Resouce not found"})
		return
	}
	townInfo := areaData[id - 1]
	w.WriteJson(townInfo)
}

func getAreas(w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	address := v.Get("address")
	page, _ := strconv.Atoi(v.Get("page"))
	displaynum, _ := strconv.Atoi(v.Get("displaynum"))
	if page < 1 || displaynum < 1 {
		w.WriteJson(Result{MatchedNum: 0, MatchedAreas: nil})
		return
	}
	foundAreas := filter(address)
	foundNum := len(foundAreas)
	left := int(math.Min(float64(foundNum), float64((page - 1) * displaynum)))
	right := int(math.Min(float64(foundNum), float64(page * displaynum)))
	foundAreas = foundAreas[left: right]
	result := Result{MatchedNum: uint(foundNum), MatchedAreas: foundAreas}
	w.WriteJson(result)
}

func filter(address string) []AreaInfo {
	var ret []AreaInfo
	for _, v := range areaData {
		if strings.Contains(v.Municipality + v.TownName, address) {
			ret = append(ret, v)
		}
	}
	return ret
}

func loadAreaFile() {
	raw, err := ioutil.ReadFile("assets/all2.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &areaData)
}