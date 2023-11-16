package rk7client

func (c Client) GetSystemInfo() (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{CMD: "GetSystemInfo"},
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
