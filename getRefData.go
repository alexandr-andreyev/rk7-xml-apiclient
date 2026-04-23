package rk7client

func (c *Client) GetSelectorGroups(rootIdent string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{{
			CMD:            RK7CMD_GETREFDATA,
			RefName:        rk7ref("SelectorGroups"),
			WithChildItems: WITHCHILDITEMS_2,
			WithMacroProp:  "1",
			RefItemIdent:   rootIdent,
		}},
	}
	req, err := c.newRequest("POST", cmd)
	if err != nil {
		return nil, err
	}
	result := RK7QueryResult{}
	_, err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c Client) GetRefData(input []RK7Command) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: input,
	}
	req, err := c.newRequest("POST", cmd)
	if err != nil {
		return nil, err
	}
	result := RK7QueryResult{}
	_, err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
