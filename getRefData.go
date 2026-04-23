package rk7client

import "fmt"

// GetSelectorGroups возвращает справочник групп селекторов (SelectorGroups) со вложенными TSelector и TSelectorDetail.
// rootIdent ограничивает выборку конкретным корневым элементом; пустая строка возвращает все группы.
func (c *Client) GetSelectorGroups(rootIdent string) (*RK7QueryResult, error) {
	cmd := RK7Query{
		RK7Command: []RK7Command{{
			CMD:            "GetRefData",
			RefName:        "SelectorGroups",
			WithChildItems: "2",
			WithMacroProp:  "1",
			RefItemIdent:   rootIdent,
		}},
	}
	req, err := c.newRequest("POST", cmd)
	if err != nil {
		return nil, err
	}
	result := RK7QueryResult{}
	_, err = c.do(req, &result)
	return &result, err
}

// GetRefData возвращает произвольный справочник R-Keeper 7 по имени refName.
// priceType указывает тип цены для включения в PropMask (например, 3 для основного прайс-листа).
// Возвращает только активные элементы (OnlyActive=true).
func (c *Client) GetRefData(refName string, priceType int) (*RK7QueryResult, error) {
	propMask := fmt.Sprintf("items.(Ident,GUIDString,Code,Name,MainParentIdent,Status,Parent,PriceTypes^%d,CategPath,ModiScheme)", priceType)
	cmd := RK7Query{
		RK7Command: []RK7Command{
			{
				CMD:            "GetRefData",
				RefName:        refName,
				OnlyActive:     "true",
				WithChildItems: "2",
				WithMacroProp:  "1",
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
	return &result, err
}
