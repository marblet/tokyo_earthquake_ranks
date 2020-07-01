package apis

import (
	"github.com/ant0ine/go-json-rest/rest"
	"math"
	"strconv"
)

func GetMunicipalityTowns (w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	areaCode := r.PathParam("areacode")
	page, _ := strconv.Atoi(v.Get("page"))
	displaynum, _ := strconv.Atoi(v.Get("displaynum"))
	foundTowns := areacodeFilter(areaCode)
	foundNum := len(foundTowns)
	left := int(math.Min(float64(foundNum), float64((page - 1) * displaynum)))
	right := int(math.Min(float64(foundNum), float64(page * displaynum)))
	foundTowns = foundTowns[left: right]
	result := Result{MatchedNum: uint(foundNum), MatchedTowns: foundTowns}
	w.WriteJson(result)
}

func areacodeFilter (areaCode string) []SimpleTownInfo {
	var ret []SimpleTownInfo
	for _, v := range simpleTownInfos {
		if v.AreaCode == areaCode {
			ret = append(ret, v)
		}
	}
	return ret
}