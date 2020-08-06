package functions

import "errors"

func required(msg string, val interface{}) (interface{}, error) {
	if val == nil {
		return val, errors.New(msg)
	} else if _, ok := val.(string); ok {
		if val == "" {
			return val, errors.New(msg)
		}
	}
	return val, nil
}
