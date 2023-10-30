package guacamole

import (
	"encoding/json"
)

func (g *Guacamole) GetConnectionTree() (GuacConnectionGroup, error) {
	body, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/ROOT/tree", nil, nil)

	var resp GuacConnectionGroup

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func flatten(nested []GuacConnectionGroup) ([]GuacConnection, []GuacConnectionGroup, error) {
	flat_conns := []GuacConnection{}
	flat_grps := []GuacConnectionGroup{}
	for _, groups := range nested {
		flat_grps = append(flat_grps, groups)
		if len(groups.ChildGroups) > 0 {
			conns, subgrps, err := flatten(groups.ChildGroups)
			if err != nil {
				return nil, nil, err
			}
			for _, c := range conns {
				flat_conns = append(flat_conns, c)
			}
			for _, g := range subgrps {
				flat_grps = append(flat_grps, g)
			}
		}
		if len(groups.ChildConnections) > 0 {
			for _, c := range groups.ChildConnections {
				flat_conns = append(flat_conns, c)
			}
		}
	}
	return flat_conns, flat_grps, nil
}

func (g *Guacamole) CreateConnectionGroup(conn *GuacConnectionGroup) (GuacConnectionGroup, error) {
	ret := GuacConnectionGroup{}
	resp, err := g.Call("POST", "/api/session/data/mysql/connectionGroups", nil, conn)
	if err != nil {
		return GuacConnectionGroup{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacConnectionGroup{}, err
	} else {
		return ret, nil
	}
}

func (g *Guacamole) ReadConnectionGroup(conn *GuacConnectionGroup) (GuacConnectionGroup, error) {
	ret := GuacConnectionGroup{}
	resp, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, nil)
	if err != nil {
		return GuacConnectionGroup{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacConnectionGroup{}, err
	} else {
		return ret, nil
	}

}

func (g *Guacamole) UpdateConnectionGroup(conn *GuacConnectionGroup) error {
	_, err := g.Call("PUT", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, conn)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (g *Guacamole) DeleteConnectionGroup(conn *GuacConnectionGroup) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, conn)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (g *Guacamole) ListConnectionGroups() ([]GuacConnectionGroup, error) {
	ret := []GuacConnectionGroup{}
	conn_tree, err := g.GetConnectionTree()
	if err != nil {
		return []GuacConnectionGroup{}, err
	}

	_, flattened_grps, err := flatten([]GuacConnectionGroup{conn_tree})
	for _, grps_from_grps := range flattened_grps {
		ret = append(ret, grps_from_grps)
	}

	return ret, nil
}
