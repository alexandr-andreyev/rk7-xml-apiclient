package rk7client

func (c Client) GetCateglist() (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:            RK7CMD_GETREFDATA,
				RefName:        RK7REF_CATEGLIST,
				OnlyActrive:    ONLY_ACTIVE_TRUE,
				WithChildItems: WITHCHILDITEMS_NO_CHILDREN,
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
	if err != nil {
		return nil, err
	}
	return &result, nil
}
