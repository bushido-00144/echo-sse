package sse

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// Set Respone header and return HTTP status to client
func ConnectionInit(response *echo.Response) {
	response.Header().Set(echo.HeaderContentType, "text/event-stream")
	response.WriteHeader(http.StatusOK)
}

// Extract Last Event ID from Request Header
// If Last Event ID Header is null, return Empty Sring
func LastEventID(request *http.Request) string {
	var header http.Header = request.Header
	if header["Last-Event-Id"] != nil {
		return header["Last-Event-Id"][0]
	} else {
		return ""
	}
}

// Ignition channel when disconnection connection to client
func ConnectionCloseNotify(response *echo.Response) <-chan bool {
	return response.CloseNotify()
}

// Write eventString to connection
func Write(eventString string, response *echo.Response) error {
	fmt.Fprintf(response.Writer, eventString)
	response.Flush()
	return nil
}

//Keep Alive for keeping connections when sandwiching Proxy
func SendHeartBeat(keepAliveSec uint64, connectionFlag *bool, response *echo.Response) {
	var keepAliveComment string = ":keep alive"
	for *connectionFlag {
		Write(keepAliveComment, response)
		time.Sleep(time.Duration(keepAliveSec) * time.Second)
	}
}
