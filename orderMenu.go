package rk7client

func (c Client) GetOrderMenu(stationCode string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:     "GetOrderMenu",
				Station: &Station{Code: stationCode},
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
