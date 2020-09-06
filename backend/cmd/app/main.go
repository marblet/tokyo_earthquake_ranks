package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"local.packages/apis"
	"log"
	"net/http"
)

func main () {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/towns", apis.GetTowns),
		rest.Get("/api/towns/:id", apis.GetTownInfo),
		rest.Get("/api/municipalities/:areacode/towns", apis.GetMunicipalityTowns),
	)
	if err != nil {
		log.Fatal(err)
	}
	apis.LoadTownFile()
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}