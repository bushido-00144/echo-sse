# Echo-SSE
Server side SSE implementation.

# Usage
```go
import (
  "./sse"
)

e.GET("/sse", func(c echo.Context) error {
  queue := event.NewQueue()

  connectionCloseNotify := sse.ConnectionCloseNotify(c.Response())
  var connectionFlag bool = true

  sse.ConnectionInit(c.Response())

  var content string = "Content of data"
  var dataField string = field.DataField(content)

  var ev event.Event = event.Event{DataField: dataField}

  var id uint64 = 1
  var eventBlock event.EventBlock = event.EventBlock{ID: id, EventContent: ev}
  event.Enqueue(eventBlock, queue)

  select {
  case <-connectionCloseNotify:
    fmt.Println("Connection Closed")
    connectionFlag = false
  default:
    if connectionFlag {
      var eventString = event.ToString(event.Dequeue(queue).EventContent)
      sse.Write(eventString, c.Response())
    }
  }

  return nil
})
```
