package response

import (
	"fmt"

	statuscode "http/status-code"
)

func addContentLenghtHeader(response string, body string) string {
	contentLength := fmt.Sprintf("%d", len(body))
	return fmt.Sprintf("%sContent-Length: %s\r\n", response, contentLength)
}

func addHeader(response string, key string, value string) string {
	return fmt.Sprintf("%s%s: %s\r\n", response, key, value)
}

func addBody(response string, body string) string {
	return fmt.Sprintf("%s\r\n%s", response, body)
}

func createResponse(statusCode int) string {
	code, phrase := statuscode.GetStatusLine(statusCode)
	return fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, phrase)
}

func Response(statusCode int, headers map[string]string, body string) string {
	responseLine := createResponse(statusCode)

	for key, value := range headers {
		responseLine = addHeader(responseLine, key, value)
	}

	responseLine = addBody(responseLine, body)

	return responseLine
}
