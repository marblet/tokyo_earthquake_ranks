package apis

import (
	"github.com/ant0ine/go-json-rest/rest"
	"math"
	"strconv"
	"strings"
)

func GetTowns (w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	address := v.Get("address")
	page, _ := strconv.Atoi(v.Get("page"))
	displaynum, _ := strconv.Atoi(v.Get("displaynum"))
	if page < 1 || displaynum < 1 {
		w.WriteJson(Result{MatchedNum: 0, MatchedTowns: nil})
		return
	}
	foundTowns := addressFilter(address)
	foundNum := len(foundTowns)
	left := int(math.Min(float64(foundNum), float64((page - 1) * displaynum)))
	right := int(math.Min(float64(foundNum), float64(page * displaynum)))
	foundTowns = foundTowns[left: right]
	result := Result{MatchedNum: uint(foundNum), MatchedTowns: foundTowns}
	w.WriteJson(result)
}

func addressFilter (address string) []SimpleTownInfo {
	var ret []SimpleTownInfo
	for _, v := range simpleTownInfos {
		if strings.Contains(v.Municipality + v.TownName, address) {
			ret = append(ret, v)
		}
	}
	return ret
}