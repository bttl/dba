package dba

import (
	"errors"
)

var ErrNotFound = errors.New("api: not found")
var ErrSomeError = errors.New("api: not found")

type usr struct {
	UserId int32 `json:"UserId"`
	UserName string `json:"UserName"`
}

func GetUserId(userName string) (
	*int32,
	error) {

	arr := []*usr{}
	
	err := readTable("UserProfile", &arr)
	if err != nil {
		return nil, err
	}

	for _, item := range arr {
		if item.UserName == userName {
			return &item.UserId, nil
		}
	}

	return nil, ErrNotFound
	
	// fmt.Println(len(arr))

	// first := arr[1000]

	// fmt.Println(first)
}

