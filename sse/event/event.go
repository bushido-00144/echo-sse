package event

type Event struct {
	IDField        string
	EventTypeField string
	DataField      string
	RetryField     string
}

// Event Structure convert to String
func ToString(event Event) string {
	var eventString string = ""
	if event.IDField != "" {
		eventString += event.IDField
	}
	if event.EventTypeField != "" {
		eventString += event.EventTypeField
	}
	if event.DataField != "" {
		eventString += event.DataField
	}
	if event.RetryField != "" {
		eventString += event.RetryField
	}
	return eventString + "\n"
}
