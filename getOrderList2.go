package rk7client

func (c Client) GetOrderList2(waiterCode string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD: RK7CMD_GETORDERLIST2,
				Waiter: &Waiter{
					Code: waiterCode,
				},
				OnlyOpened: ONLY_OPENED_TRUE,
			},
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
	//defer resp.Body.Close()
	return &result, nil
}
