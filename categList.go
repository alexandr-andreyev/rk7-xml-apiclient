package rk7client

func (c Client) GetCateglist() (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:            "GetRefData",
				RefName:        "CATEGLIST",
				OnlyActrive:    "true",
				WithChildItems: "0",
				PropMask:       "items.(Ident,GUIDString,Code,Name,MainParentIdent,Status,Parent)",
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
