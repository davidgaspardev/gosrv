package helpers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Request struct {
	*http.Request
	params map[string]interface{}
}

// Param is the variables inside the path from request
// E.g:
//  - /v1/estabels/104/resources/inj003
//  - /v1/estabels/105/resources/smd09
// Both are the same route, but with params, in the first example the 104 and inj003 are params.
// And in the second example the 105 and smd09 are params.

// Add param from
func (req *Request) AddParam(key string, value interface{}) {
	if req.params == nil {
		req.params = make(map[string]interface{})
	}

	req.params[key] = value
}

func (req *Request) GetParam(key string) interface{} {
	return req.params[key]
}

func (req *Request) HasRequestOrigin() bool {
	return req.Header.Get(HEADER_ORIGIN) != ""
}

func (req *Request) GetOrigin() string {
	return req.Header.Get(HEADER_ORIGIN)
}

func (req *Request) GetContentType() string {
	return req.Header.Get(HEADER_CONTENT_TYPE)
}

func (req *Request) GetAccept() string {
	return req.Header.Get(HEADER_ACCEPT)
}

// func ExtractRequestBody(request *http.Request, data model.Model) error {
// 	err := json.NewDecoder(request.Body).Decode(data)
// 	if err != nil {
// 		return err
// 	}
// 	if !data.IsValid() {
// 		return fmt.Errorf("invalid request body")
// 	}
// 	return nil
// }

func ExtractRequestBodyMap(request *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	result := make(map[string]interface{})

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for key := range data {
		result[ToSnake(key)] = data[key]
	}

	data = nil
	return result, nil
}

func ToSnake(camel string) (snake string) {
	var b strings.Builder
	diff := 'a' - 'A'
	l := len(camel)
	for i, v := range camel {
		// A is 65, a is 97
		if v >= 'a' {
			b.WriteRune(v)
			continue
		}
		// v is capital letter here
		// irregard first letter
		// add underscore if last letter is capital letter
		// add underscore when previous letter is lowercase
		// add underscore when next letter is lowercase
		if (i != 0 || i == l-1) && (          // head and tail
		(i > 0 && rune(camel[i-1]) >= 'a') || // pre
			(i < l-1 && rune(camel[i+1]) >= 'a')) { //next
			b.WriteRune('_')
		}
		b.WriteRune(v + diff)
	}
	return b.String()
}
