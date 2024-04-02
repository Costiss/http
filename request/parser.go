package request

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	RawBody string
	Headers map[string]string
	IP      string
}

func splitRequest(data string) ([]string, error) {
	fmt.Printf("data: %s\n", data)
	return strings.Split(data, "\r\n"), nil
}

func requestParser(bytes []byte) (*Request, error) {
	data, err := splitRequest(string(bytes))
	if err != nil {
		return nil, err
	}

	request := Request{}
	parseRequestLine(data[0], &request)

	return &request, nil
}

func parseRequestLine(data string, reques *Request) error {
	return nil
}
