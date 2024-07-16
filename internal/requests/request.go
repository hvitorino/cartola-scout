package requests

import "strings"

type Request struct {
	Url    string
	Method string
}

func (r *Request) queryParams(params map[string]string) {
	r.Url = parseString(r.Url, params)
}

func parseString(text string, params map[string]string) string {
	parsedText := text

	for k, v := range params {
		parsedText = strings.ReplaceAll(parsedText, k, v)
	}

	return parsedText
}
