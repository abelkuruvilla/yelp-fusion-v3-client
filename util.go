package yelp

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func structToMap(s interface{}) (map[string]interface{}, error) {

	m := make(map[string]interface{})
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, err
}

func structToQueryParams(s interface{}) (string, error) {

	query := "?"

	m, err := structToMap(s)
	if err != nil {
		return "", err
	}
	for k, v := range m {
		if reflect.ValueOf(v).Kind() == reflect.String {
			v = strings.Join(strings.Split(v.(string), " "), "+")
		}
		query = query + fmt.Sprintf("%v=%v&", k, v)
	}
	t := strings.LastIndex(query, "&")
	if t > 0 {
		query = query[:t]
	}

	return query, nil
}
