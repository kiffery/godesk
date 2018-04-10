package godesk

// input data for getting multiple requests
type RequestInputData struct {
	Operation RequestOperations `json:"operation"`
}

// part of json stucture
type RequestOperations struct {
	Details RequestDetails `json:"details"`
}

// the start, limit, and queue from hich requests are coming from
type RequestDetails struct {
	From     string `json:"from"`
	Limit    string `json:"limit"`
	Filterby string `json:"filterby"`
}

// json format that request is returned in
type SDRequests struct {
	Operation struct {
		Resources []string `json:"resources"`
		Result    struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"result"`
		Module   string    `json:"module"`
		Requests []Request `json:"details"`
	} `json:"operation"`
	Name        string      `json:"name"`
	SSPSettings SSPSettings `json:"SSP_SETTINGS"`
}

// SDrequest feilds for single request
type SDRequest struct {
	Workorderid                  string `json:"WORKORDERID"`
	Requester                    string `json:"REQUESTER"`
	Createdby                    string `json:"CREATEDBY"`
	Createdtime                  string `json:"CREATEDTIME"`
	Duebytime                    string `json:"DUEBYTIME"`
	Responseduebytime            string `json:"RESPONSEDUEBYTIME"`
	Frduetime                    string `json:"FR_DUETIME"`
	Respondedtime                string `json:"RESPONDEDTIME"`
	Resolvedtime                 string `json:"RESOLVEDTIME"`
	Completedtime                string `json:"COMPLETEDTIME"`
	Shortdescription             string `json:"SHORTDESCRIPTION"`
	Timespentonreq               string `json:"TIMESPENTONREQ"`
	Subject                      string `json:"SUBJECT"`
	Requesttemplate              string `json:"REQUESTTEMPLATE"`
	Mode                         string `json:"MODE"`
	SLA                          string `json:"SLA"`
	Asset                        string `json:"ASSET"`
	Department                   string `json:"DEPARTMENT"`
	Editorid                     string `json:"EDITORID"`
	Editingstatus                string `json:"EDITING_STATUS"`
	Iscatalogtemplate            string `json:"IS_CATALOG_TEMPLATE"`
	Site                         string `json:"SITE"`
	Isvipuser                    string `json:"ISVIPUSER"`
	Service                      string `json:"SERVICE"`
	Category                     string `json:"CATEGORY"`
	Subcategory                  string `json:"SUBCATEGORY"`
	Item                         string `json:"ITEM"`
	Technician                   string `json:"TECHNICIAN"`
	Technicianloginname          string `json:"TECHNICIAN_LOGINNAME"`
	Status                       string `json:"STATUS"`
	Priority                     string `json:"PRIORITY"`
	Level                        string `json:"LEVEL"`
	Impact                       string `json:"IMPACT"`
	Urgency                      string `json:"URGENCY"`
	Impactdetails                string `json:"IMPACTDETAILS"`
	Requesttype                  string `json:"REQUESTTYPE"`
	Approvalstatus               string `json:"APPROVAL_STATUS"`
	Closurecode                  string `json:"CLOSURECODE"`
	Closurecomments              string `json:"CLOSURECOMMENTS"`
	Fcr                          string `json:"FCR"`
	Yettoreplycount              string `json:"YETTOREPLYCOUNT"`
	Group                        string `json:"GROUP"`
	Description                  string `json:"DESCRIPTION"`
	FootprintsTicket             string `json:"Footprints Ticket #"`
	FootprintsTeamsAssignees     string `json:"Footprints Teams/Assignees"`
	TicketGrade                  string `json:"Ticket Grade"`
	Knowledgebase                string `json:"Knowledgebase"`
	Tags                         string `json:"Tags"`
	BusinessAnalystGroup         string `json:"Business Analyst Group"`
	EasFunctionalArea            string `json:"EAS Functional Area"`
	DeploymentChecklistCompleted string `json:"Deployment Checklist Completed"`
	Operation                    struct {
		Result struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"result"`
	} `json:"operation"`
}

// part of multiple requests struct
type SSPSettings struct {
	RequesterReopenOption         bool `json:"REQUESTER_REOPEN_OPTION"`
	DisableDefaultRequestTemplate bool `json:"DISABLE_DEFAULT_REQUEST_TEMPLATE"`
	GlobalRequestOnholdOption     bool `json:"GLOBAL_REQUEST_ONHOLD_OPTION"`
	RequesterCloseOption          bool `json:"REQUESTER_CLOSE_OPTION"`
}

// request struct for getting multiple requests
type Request struct {
	Workorderid int    `json:"WORKORDERID"`
	Requester   string `json:"REQUESTER"`
	Createdby   string `json:"CREATEDBY"`
	Createdtime int    `json:"CREATEDTIME"`
	Duebytime   int    `json:"DUEBYTIME"`
	Subject     string `json:"SUBJECT"`
	Technician  string `json:"TECHNICIAN"`
	Status      string `json:"STATUS"`
	Priority    string `json:"PRIORITY"`
}
