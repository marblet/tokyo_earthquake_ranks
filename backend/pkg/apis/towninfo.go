package apis

import (
	"github.com/ant0ine/go-json-rest/rest"
	"strconv"
)

func GetTownInfo(w rest.ResponseWriter, r *rest.Request) {
	id, _ := strconv.Atoi(r.PathParam("id"))
	if id < 1 || id > len(townInfos) {
		w.WriteJson(map[string]string{"Error": "Resouce not found"})
		return
	}
	townInfo := townInfos[id - 1]
	w.WriteJson(townInfo)
}