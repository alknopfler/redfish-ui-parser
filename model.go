package main

// RedfishResponse is the response from the Redfish API
type RedfishResponse struct {
	Members []struct {
		OdataID string `json:"@odata.id"`
	} `json:"Members"`
}

// SystemDetail is the detailed information about a system
type SystemDetail struct {
	UUID       string `json:"UUID"`
	Name       string `json:"Name"`
	PowerState string `json:"PowerState"`
	Status     struct {
		State        string `json:"State"`
		Health       string `json:"Health"`
		HealthRollUp string `json:"HealthRollUp"`
	} `json:"Status"`
}

// PageData is the data to be passed to the template html file
type PageData struct {
	Systems []SystemDetail
}
