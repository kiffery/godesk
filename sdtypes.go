package godesk

type ServiceDesk struct {
	Hostname      string
	TechnicianKey string
}

type SDRequester struct {
	Operation struct {
		Name      string `json:"name"`
		Module    string `json:"module"`
		TotalRows string `json:"TotalRows"`
		Result    struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"result"`
		Details []struct {
			Userid            int           `json:"userid"`
			Username          string        `json:"username"`
			Emailid           string        `json:"emailid"`
			Department        string        `json:"department"`
			Sitename          string        `json:"sitename"`
			Loginname         string        `json:"loginname"`
			Domainname        string        `json:"domainname"`
			Employeeid        string        `json:"employeeid"`
			Jobtitle          string        `json:"jobtitle"`
			Landline          string        `json:"landline"`
			Mobile            string        `json:"mobile"`
			Secondaryemailids []interface{} `json:"secondaryemailids"`
			Isvipniluser      bool          `json:"isvipuser"`
		} `json:"details"`
		Resources   []interface{} `json:"resources"`
		SSPSETTINGS struct {
		} `json:"SSP_SETTINGS"`
	} `json:"operation"`
}
