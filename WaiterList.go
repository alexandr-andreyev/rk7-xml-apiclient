package rk7client

func (c Client) GetWaiterlist(registeredOnly bool) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:            RK7CMD_GETWAITERLIST,
				RegisteredOnly: REGISTEREDONLY_TRUE,
			},
		},
	}
	req, err := c.newRequest("POST", cmd)
	if err != nil {
		return nil, err
	}
	result := RK7QueryResult{}
	_, err = c.do(req, &result)
	//defer resp.Body.Close()
	return &result, nil
}
