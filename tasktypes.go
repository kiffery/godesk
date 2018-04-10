package godesk

// inputdata for json request
type TaskInputData struct {
	ListInfo       TasksListInfo `json:"list_info"`
	FieldsRequired []string      `json:"fields_required"`
}

// part of inputdata for json request
type TasksListInfo struct {
	Filter     string `json:"filter"`
	RowCount   string `json:"row_count"`
	StartIndex string `json:"start_index"`
	EndIndex   string `json:"end_index"`
	SortColumn string `json:"sort_column"`
	SortOrder  string `json:"sort_order"`
}

// Singular task stuct
type SDTask struct {
	Task           Task               `json:"task"`
	ResponseStatus TaskResponseStatus `json:"response_status"`
}

// multiple tasks struct
type SDTasks struct {
	FeildsRequired []string           `json:"fields_required"`
	ResponseStatus TaskResponseStatus `json:"response_status"`
	Tasks          []Task             `json:"tasks"`
	ListInfo       struct {
		HasMoreRows bool   `json:"has_more_rows"`
		EndIndex    string `json:"end_index"`
		StartINdex  string `json:"start_index" `
		PageNumber  string `json:"page_number"`
		RowCount    string `json:"row_count"`
		TotalCount  string `json:"total_count"`
	} `json:"list_info"`
}

// single task used by both task and multiple task calls
type Task struct {
	CreatedBy struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"created_by"`
	CreatedDate struct {
		DisplayValue string `json:"display_value"`
		Value        string `json:"value"`
	} `json:"created_date"`
	PercentageCompletion string `json:"percentage_completion"`
	Status               struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	AdditionalCost         string      `json:"additional_cost"`
	EstimatedEffortDays    interface{} `json:"estimated_effort_days"`
	EstimatedEffortMinutes interface{} `json:"estimated_effort_minutes"`
	ID                     string      `json:"id"`
	EstimatedEffortHours   interface{} `json:"estimated_effort_hours"`
	Title                  string      `json:"title"`
	EmailBefore            string      `json:"email_before"`
	Overdue                string      `json:"overdue"`
	Description            string      `json:"description"`
	Request                struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"request"`
	Owner struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	AssociatedEntity string `json:"associated_entity"`
	Group            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
}

// part of tast json stuct returned
type TaskResponseStatus struct {
	Status   string `json:"status"`
	Messages []struct {
		Type       string `json:"type"`
		Message    string `json:"message"`
		StatusCode string `json:"status_code"`
	} `json:"messages"`
}
