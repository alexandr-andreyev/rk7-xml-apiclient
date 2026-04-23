package rk7client

// GetOrderMenu возвращает меню заказа для указанной кассовой станции.
// stationCode — код станции (поле Code в справочнике станций).
func (c *Client) GetOrderMenu(stationCode string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:     RK7CMD_GETORDERMENU,
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
	if err != nil {
		return nil, err
	}
	return &result, nil
}
