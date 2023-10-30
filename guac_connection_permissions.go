package guacamole

import (
	"encoding/json"
)

func (g *Guacamole) SendUserConnectionPermissionChanges(user string, p []GuacPermissionItem) error {
	url := "/api/session/data/mysql/users/" + user + "/permissions"
	_, err := g.Call("PATCH", url, nil, p)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guacamole) GetUserConnectionPermissions(user string) (GuacPermissionData, error) {
	ret := GuacPermissionData{}
	url := "/api/session/data/mysql/users/" + user + "/permissions"
	resp, err := g.Call("GET", url, nil, nil)
	if err != nil {
		return GuacPermissionData{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacPermissionData{}, err
	}

	return ret, nil
}
