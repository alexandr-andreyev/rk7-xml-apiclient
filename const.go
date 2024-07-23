package rk7client

type rk7cmd string

const (
	RK7CMD_GETREFDATA     rk7cmd = "GetRefData"
	RK7CMD_GETWAITERLIST  rk7cmd = "GetWaiterList"
	RK7CMD_GETREFLIST     rk7cmd = "GetRefList"
	RK7CMD_GETSYSTEMINFO2 rk7cmd = "GetSystemInfo2"
	RK7CMD_GETORDERMENU   rk7cmd = "GetOrderMenu"
	RK7CMD_GETORDERLIST2  rk7cmd = "GetOrderList2"
)

type rk7ref string

const (
	RK7REF_EMPLOYEES rk7ref = "Employees"
	RK7REF_CATEGLIST rk7ref = "CATEGLIST"
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

type registeredOnly string

const (
	REGISTEREDONLY_TRUE  = "1"
	REGISTEREDONLY_FAlSE = "0"
)

type onlyOpened string

const (
	ONLY_OPENED_TRUE  = "true"
	ONLY_OPENED_FALSE = "false"
)
