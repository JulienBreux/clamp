package functions

import "errors"

func required(msg string, val interface{}) (interface{}, error) {
	if val == nil {
		return val, errors.New(msg)
	}
	if _, ok := val.(string); ok && val == "" {
		return val, errors.New(msg)
	}
	return val, nil
}
