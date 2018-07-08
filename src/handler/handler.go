package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	data "gitlab.com/ashwini/ProjectMystique/src/databases"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("This page is not implemented")

})

var GetLandingPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	lp := filepath.Join("public", "layout.html")
	fp := filepath.Join("public", "index.html")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

})

var GetRecentBlocks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This returns pre-defined number of recent blocks Say 10")
})

var GetRecentBlocksByNumber = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	numberOfBlocks := params["number"]
	fmt.Println(numberOfBlocks + " asked by the user")
})

var GetSettings = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This returns the setting page TODO: Add the page and serve it ")
})

var GetPackageTable = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	packageId := params["packageId"]
	json.NewEncoder(w).Encode("Pakcage Id " + packageId)
})

var GetAllPackages = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.GetAllPackage())
})

var GetAllBlocks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.GetAllBlock())
})

var TrackPackageWithIdLtd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	packageId := params["packageId"]
	json.NewEncoder(w).Encode(data.GetRecentBlocksForThePackage(packageId))
})

var TrackPackageWithId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	packageId := params["packageId"]
	json.NewEncoder(w).Encode(data.GetAllBlocksForAPackage(packageId))
})

var GetBlockInfoWithId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blockId := params["blockId"]
	json.NewEncoder(w).Encode(data.GetBlockInfoStructById(blockId))
})
var Search = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queryString := params["searchstring"]
	json.NewEncoder(w).Encode("You want to search for " + queryString)
})
