package rk7client

type rk7cmd string

const (
	RK7CMD_GETREFDATA rk7cmd = "GetRefData"
)

type rk7ref string

const (
	RK7REF_EMPLOYEES rk7ref = "Employees"
)

type onlyactive string

const (
	ONLY_ACTIVE_TRUE  onlyactive = "true"
	ONLY_ACTIVE_FALSE onlyactive = "false"
)

type withChildItems string

const (
	WITHCHILDITEMS_NO_CHILDREN withChildItems = "0"
	WITHCHILDITEMS_1           withChildItems = "1"
	WITHCHILDITEMS_2           withChildItems = "2"
	WITHCHILDITEMS_3           withChildItems = "3"
)

type withMacroProp string

const (
	WITHMACROPROP_TRUE  withMacroProp = "true"
	WITHMACROPROP_FALSE withMacroProp = "false"
)
