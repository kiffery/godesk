package godesk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	V1API = "/sdpapi"
	V3API = "/api/v3"
)

// creates a new SD object for using methods and calling api
// takes a key and servicedesk url
func CreateSDObject(key string, host string) *ServiceDesk {
	SD := new(ServiceDesk)
	SD.Hostname = host
	SD.TechnicianKey = key
	return SD
}

// Makes a URL
func (SD *ServiceDesk) makeURL(apiType, apiCall, id string, parameters map[string]string) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = SD.Hostname
	if id == "" {
		u.Path = fmt.Sprintf("%s/%s", apiType, apiCall)
	} else {
		u.Path = fmt.Sprintf("%s/%s/%s", apiType, apiCall, id)
	}
	query := u.Query()
	for key, value := range parameters {
		query.Add(key, value)
	}
	u.RawQuery = query.Encode()
	return u.String()
}

// calls api
func (SD *ServiceDesk) callAPI(method, apiType, apiCall, id string, parameters map[string]string) ([]byte, error) {
	parameters["format"] = "json"
	parameters["TECHNICIAN_KEY"] = SD.TechnicianKey
	URL := SD.makeURL(apiType, apiCall, id, parameters)
	fmt.Println(URL)
	req, err := http.NewRequest(method, URL, strings.NewReader(""))
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
