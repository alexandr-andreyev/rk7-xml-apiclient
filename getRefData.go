package rk7client

func (c Client) GetRefData(refName string, onlyActive string, withChildItems string, withMacroProp string, propMask string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:            "GetRefData",
				RefName:        refName,
				OnlyActrive:    onlyActive,
				WithChildItems: withChildItems,
				WithMacroProp:  withMacroProp,
				PropMask:       propMask,
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
