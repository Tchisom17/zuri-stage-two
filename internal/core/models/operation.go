package models

import (
	"regexp"
	"strconv"
)

type Operation struct {
	SlackUsername string `json:"slackUsername"`
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Result        int    `json:"result"`
}

func (o *Operation) ReadString(str string) []int {
	var nums []int
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(str, -1)
	for _, element := range submatchall {
		num, err := strconv.Atoi(element)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}
