package main

import (
	"./sse"
	"./sse/event"
	"./sse/field"
	"fmt"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

func main() {
	e := echo.New()

	e.File("/", "public/index.html")

	e.GET("/sse", func(c echo.Context) error {
		// Get LastEventID
		lastEventID := sse.LastEventID(c.Request())
		fmt.Println(lastEventID)

		// Prepare queue for Event
		queue := event.NewQueue()

		// Notification channel for detection disconnection
		connectionCloseNotify := sse.ConnectionCloseNotify(c.Response())
		var connectionFlag bool = true

		// Init process for SSE
		sse.ConnectionInit(c.Response())

		// Setting Keep Alive (for HTTP Proxy)
		go sse.SendHeartBeat(15, &connectionFlag, c.Response())

		var i int = 1000
		for i > 0 {
			// Generate Content
			var content string = strconv.Itoa(i)
			var dataField string = field.DataField(content)

			// Generate ID
			var idField string = field.IDField(content)

			// Generate Event Filed
			var eventType string = "hoge"
			var eventTypeField string = field.EventTypeField(eventType)

			// Generate Event
			var ev event.Event = event.Event{
				DataField:      dataField,
				IDField:        idField,
				EventTypeField: eventTypeField,
			}

			// Generate Block for queuing events by using id as a key
			var id uint64
			id, _ = strconv.ParseUint(field.GetFieldContent(idField), 10, 64)
			var eventBlock event.EventBlock = event.EventBlock{ID: id, EventContent: ev}
			event.Enqueue(eventBlock, queue)

			select {
			case <-connectionCloseNotify:
				// Connection closed by client
				fmt.Println("Connection Closed")
				connectionFlag = false
			default:
				if connectionFlag {
					// Dequeue and Send to client
					var eventString = event.ToString(event.Dequeue(queue).EventContent)
					sse.Write(eventString, c.Response())
				}
			}

			time.Sleep(1 * time.Second)
			i--
		}

		return nil
	})

	e.Start(":1323")
}
