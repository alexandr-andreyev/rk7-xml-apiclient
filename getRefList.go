package rk7client

func (c Client) GetRefList() (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{CMD: "GetRefList"},
		},
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
