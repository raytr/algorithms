package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func main() {
	m := &Manager{
		FullName:       "Jack Oliver",
		Position:       "CEO",
		Age:            44,
		YearsInCompany: 0,
	}

	res, _ := EncodeManager(m)
	fmt.Println(res)
}

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	bs, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(bs)

	return r, nil
}
