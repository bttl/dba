package dba

import (
	"io/ioutil"
	"encoding/json"
)

func readTable(tblName string, mdl interface{}) error {
	content, err := ioutil.ReadFile("./data/" + tblName + ".json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, mdl)

	if err != nil {
		return err
	}

	return nil
}
