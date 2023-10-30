package guacamole

import (
	"encoding/json"
)

func (g *Guacamole) CreateUser(user *GuacUser) (GuacUser, error) {
	ret := GuacUser{}
	resp, err := g.Call("POST", "/api/session/data/mysql/users", nil, user)
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guacamole) ReadUser(user *GuacUser) (GuacUser, error) {
	ret := GuacUser{}
	resp, err := g.Call("GET", "/api/session/data/mysql/users/"+user.Username, nil, nil)
	if err != nil {
		return GuacUser{}, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guacamole) UpdateUser(user *GuacUser) error {
	_, err := g.Call("PUT", "/api/session/data/mysql/users/"+user.Username, nil, user)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guacamole) DeleteUser(user *GuacUser) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/users/"+user.Username, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g *Guacamole) ListUsers() ([]GuacUser, error) {
	ret := []GuacUser{}
	marshalledResponse := map[string]GuacUser{}
	user_tree, err := g.Call("GET", "/api/session/data/mysql/users", nil, nil)

	err = json.Unmarshal(user_tree, &marshalledResponse)
	if err != nil {
		return []GuacUser{}, err
	} else {
		for _, datablock := range marshalledResponse {
			ret = append(ret, datablock)
		}
	}
	return ret, nil
}
