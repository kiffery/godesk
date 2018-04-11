# godesk

It is a api made to call **ManageEngine ServiceDesk**

## Uses

It currently is made to get multiple or singular requests and tasks. More functionality will be created later

## How to use

### go get pkg

```bash
go get github.com/kiffery/godesk
```

### Import pkg

```golang
import (
    "github.com/kiffery/godesk"
)
```

### Example uses

#### Create SDobject
```golang
	helper := godesk.CreateSDObject("key", "host")
	if err != nil {
		fmt.Print("error")
	}
	fmt.Printf("%+v\n", *request)
}
```
#### Get multiple requests
```golang
    // This example will return requests five starting at index 1 from the spcified queue
    requests, err := helper.GetRequests(1, 5, "Queue")
```

#### Get single request
```golang
    // Returns request specified
    request, err := helper.GetRequest(222222)
```

#### Get single task
```golang
    // Returns task specified
    task, err := helper.GetTask(22222)
```

#### Get multiple tasks
```golang
    // This example returns 20 tasks from the specified queue
    tasks, err := helper.GetTasks(20, "queue")
```

#### Get RequestFilters/Queues
```golang
	// This example returns a map of all request filters and queues related to ServiceDesk instance
    filters, err := helper.GetRequestFilters()
	if err != nil {
		fmt.Print("error")
	}
	for key, queue := range filters {
		fmt.Printf("Pretty: %v\nActual: %v\n", key, queue)
    }
```

## TODO

1. Make README better
2. Create better examples
3. Fix spelling errors