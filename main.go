package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// This variable could be included in the html template to be defined by the user in the future
var redfishURL = "http://192.168.122.1:8000"

func fetchRedfishSystems() ([]SystemDetail, error) {
	baseURL := redfishURL + "/redfish/v1/Systems"
	var response RedfishResponse

	// Fetch the base systems list
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch systems list: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse systems list: %v", err)
	}

	var systems []SystemDetail
	for _, member := range response.Members {
		// Fetch each system detail
		systemResp, err := http.Get(redfishURL + member.OdataID)
		if err != nil {
			log.Printf("failed to fetch system %s: %v", member.OdataID, err)
			continue
		}
		defer systemResp.Body.Close()

		var system SystemDetail
		err = json.NewDecoder(systemResp.Body).Decode(&system)
		if err != nil {
			log.Printf("failed to parse system %s: %v", member.OdataID, err)
			continue
		}

		systems = append(systems, system)
	}

	return systems, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	systems, err := fetchRedfishSystems()
	if err != nil {
		http.Error(w, "Unable to fetch systems", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	pageData := PageData{Systems: systems}
	tmpl.Execute(w, pageData)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Server running at port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
