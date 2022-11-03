package helpers

import "errors"

func OperationEnum(s string) (string, error) {
	var response string
	types := [...]string{"addition", "subtraction", "multiplication"}

	for _, v := range types {
		if v == s {

			response = s

			return response, nil
		}
	}
	response = ""

	return response, errors.New("incorrect input field")
}
