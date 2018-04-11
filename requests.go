package godesk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (SD *ServiceDesk) GetRequestFilters() (*RequestFilters, error) {
	params := make(map[string]string)
	params["OPERATION_NAME"] = "GET_REQUEST_FILTERS"

	data, err := SD.callAPI("POST", V1API, "request", "", params)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	filters := new(RequestFilters)
	err = json.Unmarshal(data, &filters)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	return filters, err

}

// Requires from(int), limit(int), and requested queue(string)
// from is a starting point 0 is best choice
// limit is how many requests we would get
// queue is the name of queue according to sdesk or "All_Requests"
func (SD *ServiceDesk) GetRequests(from, limit int, queue string) (*SDRequests, error) {
	params := make(map[string]string)
	params["OPERATION_NAME"] = "GET_REQUESTS"

	inputdata := new(RequestInputData)
	inputdata.Operation.Details.From = strconv.Itoa(from)
	inputdata.Operation.Details.Limit = strconv.Itoa(limit)
	inputdata.Operation.Details.Filterby = queue
	jsonData, err := json.Marshal(inputdata)
	if err != nil {
		return nil, err
	}
	params["INPUT_DATA"] = string(jsonData)

	data, err := SD.callAPI("POST", V1API, "request", "", params)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	requests := new(SDRequests)
	err = json.Unmarshal(data, &requests)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	return requests, err
}

// Takes a single valid request number
// returns a *SDRequest object, can be found in requeststypes.go
func (SD *ServiceDesk) GetRequest(id int) (*SDRequest, error) {
	params := make(map[string]string)
	params["OPERATION_NAME"] = "GET_REQUEST"
	data, err := SD.callAPI("POST", V1API, "request", strconv.Itoa(id), params)
	if err != nil {
		return nil, err
	}

	request := new(SDRequest)
	err = json.Unmarshal(data, &request)
	if err != nil {
		return nil, err
	}

	return request, err
}
