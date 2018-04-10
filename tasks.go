package godesk

import (
	"encoding/json"
	"strconv"
)

func (SD *ServiceDesk) GetTask(id int) (*SDTask, error) {
	params := make(map[string]string)

	data, err := SD.callAPI("GET", V3API, "tasks", strconv.Itoa(id), params)
	if err != nil {
		return nil, err
	}

	task := new(SDTask)
	json.Unmarshal(data, &task)

	return task, err
}

func (SD *ServiceDesk) GetTasks(rowCount int, filter string) (*SDTasks, error) {
	inputData := new(TaskInputData)
	feildsRequired := []string{"created_by", "created_date", "percentage_completion", "status", "additional_cost", "estimated_effort_days", "estimated_effort_minutes", "id", "estimated_effort_hours", "title", "email_before", "overdue", "description", "request", "owner", "group"}
	inputData.ListInfo.RowCount = strconv.Itoa(rowCount)
	inputData.FieldsRequired = feildsRequired
	inputData.ListInfo.StartIndex = "1"
	inputData.ListInfo.EndIndex = "2"
	inputData.ListInfo.SortColumn = "status"
	inputData.ListInfo.SortOrder = "A"
	inputData.ListInfo.Filter = filter
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return nil, err
	}
	params := make(map[string]string)
	params["INPUT_DATA"] = string(jsonData)

	data, err := SD.callAPI("GET", V3API, "tasks", "", params)
	if err != nil {
		return nil, err
	}

	tasks := new(SDTasks)
	json.Unmarshal(data, &tasks)

	return tasks, err

}
