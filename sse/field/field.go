package field

import (
	"strconv"
	"strings"
)

var DATAFIELDNAME string = "data: "
var EVENTTYPEFIELDNAME string = "event: "
var RETRYFIELDNAME string = "retry: "
var IDFIELDNAME string = "id: "

// Get value from Field
func GetFieldContent(field string) string {
	var content string = strings.Replace(field, DATAFIELDNAME, "", -1)
	content = strings.Replace(content, EVENTTYPEFIELDNAME, "", -1)
	content = strings.Replace(content, RETRYFIELDNAME, "", -1)
	content = strings.Replace(content, IDFIELDNAME, "", -1)
	return content
}

// Generate Field from field name and content
// If field content is more than one line, insert field name before each line
func convertToField(name string, content string) string {
	var splitedContent []string = strings.Split(content, "\n")
	var fields string = name + strings.Join(splitedContent, "\n"+name) + "\n"
	return fields
}

// Generate Data field
func DataField(content string) string {
	return convertToField(DATAFIELDNAME, content)
}

// Generate Event Field
func EventTypeField(content string) string {
	return convertToField(EVENTTYPEFIELDNAME, content)
}

// Generate Retry Field
func RetryField(retryTime uint64) string {
	var retryTimeStr string = strconv.FormatUint(retryTime, 10)
	return convertToField(RETRYFIELDNAME, retryTimeStr)
}

// Generate ID Field
func IDField(idString string) string {
	return convertToField(IDFIELDNAME, idString)
}
