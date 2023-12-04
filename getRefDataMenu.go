package rk7client

import "fmt"

func (c Client) GetRefDataMenu(priceType int) (*RK7QueryResult, error) {
	propMask := fmt.Sprintf("items.(Ident,GUIDString,Code,Name,MainParentIdent,Status,Parent,PriceTypes^%d,CategPath,ModiScheme)", priceType)
	cmd := RK7Query{
		RK7Command2: []RK7Command2{
			{
				CMD:            "GetRefData",
				RefName:        "MENUITEMS",
				OnlyActrive:    "true",
				WithChildItems: "2",
				WithMacroProp:  "1",
				//PropMask:       "*",
				PropMask: propMask,
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
