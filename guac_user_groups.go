package guacamole

import (
	"encoding/json"
)

func (g *Guacamole) CreateUserGroup(group *GuacUserGroup) (GuacUserGroup, error) {
	ret := GuacUserGroup{}
	resp, err := g.Call("POST", "/api/session/data/mysql/userGroups", nil, group)
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guacamole) ReadUserGroup(group *GuacUserGroup) (GuacUserGroup, error) {
	ret := GuacUserGroup{}
	resp, err := g.Call("GET", "/api/session/data/mysql/userGroups/"+group.Identifier, nil, nil)
	if err != nil {
		return GuacUserGroup{}, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guacamole) UpdateUserGroup(group *GuacUserGroup) error {
	_, err := g.Call("PUT", "/api/session/data/mysql/userGroups/"+group.Identifier, nil, group)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guacamole) DeleteUserGroup(group *GuacUserGroup) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/userGroups/"+group.Identifier, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guacamole) ListUserGroups() ([]GuacUserGroup, error) {
	ret := []GuacUserGroup{}
	marshalledResponse := map[string]GuacUserGroup{}
	grp_tree, err := g.Call("GET", "/api/session/data/mysql/userGroups", nil, nil)
	err = json.Unmarshal(grp_tree, &marshalledResponse)
	if err != nil {
		return []GuacUserGroup{}, err
	} else {
		for _, datablock := range marshalledResponse {
			ret = append(ret, datablock)
		}
	}
	return ret, nil
}
