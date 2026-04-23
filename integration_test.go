//go:build integration

package rk7client

import (
	"testing"
)

const (
	integHost = "192.168.0.91"
	integPort = 8086
	integUser = "http"
	integPass = "http"
)

func newIntegClient() *Client {
	return NewClient(integHost, integPort, integUser, integPass)
}

func requireOk(t *testing.T, result *RK7QueryResult, method string) {
	t.Helper()
	if result.Status != "Ok" {
		t.Errorf("%s: Status=%q", method, result.Status)
	}
	if result.CommandResult == nil {
		t.Fatalf("%s: CommandResult is nil", method)
	}
	if result.CommandResult.ErrorText != "" {
		t.Logf("%s: ErrorText=%s", method, result.CommandResult.ErrorText)
	}
}

func TestIntegration_GetSystemInfo(t *testing.T) {
	c := newIntegClient()
	result, err := c.GetSystemInfo()
	if err != nil {
		t.Fatalf("GetSystemInfo error: %v", err)
	}
	requireOk(t, result, "GetSystemInfo")
	si := result.CommandResult.SystemInfo
	if si == nil {
		t.Fatal("SystemInfo is nil")
	}
	t.Logf("Server: %s, RestCode: %s, ShiftDate: %s, CashGroup: %s",
		result.NetName, si.RestCode, si.ShiftDate, si.CashGroup.Name)
}

func TestIntegration_GetCateglist(t *testing.T) {
	c := newIntegClient()
	result, err := c.GetCateglist()
	if err != nil {
		t.Fatalf("GetCateglist error: %v", err)
	}
	requireOk(t, result, "GetCateglist")
	ref := result.CommandResult.RK7Reference
	if ref == nil {
		t.Fatal("RK7Reference is nil")
	}
	items := ref.Items.Item
	t.Logf("Categlist: TotalItemCount=%s, received=%d", ref.TotalItemCount, len(items))
	if len(items) == 0 {
		t.Error("expected at least one category, got 0")
	}
}

func TestIntegration_GetSelectorGroups(t *testing.T) {
	c := newIntegClient()
	result, err := c.GetSelectorGroups("")
	if err != nil {
		t.Fatalf("GetSelectorGroups error: %v", err)
	}
	requireOk(t, result, "GetSelectorGroups")
	ref := result.CommandResult.RK7Reference
	if ref == nil {
		t.Fatal("RK7Reference is nil")
	}
	items := ref.Items.Item
	t.Logf("SelectorGroups: TotalItemCount=%s, received=%d", ref.TotalItemCount, len(items))
	for _, item := range items {
		t.Logf("  Group Ident=%s Name=%q Selectors=%d", item.Ident, item.Name, len(item.Selectors))
	}
}

func TestIntegration_GetRefData_MenuItems(t *testing.T) {
	c := newIntegClient()
	result, err := c.GetRefData("MenuItems", 1)
	if err != nil {
		t.Fatalf("GetRefData(MenuItems) error: %v", err)
	}
	requireOk(t, result, "GetRefData(MenuItems)")
	ref := result.CommandResult.RK7Reference
	if ref == nil {
		t.Fatal("RK7Reference is nil")
	}
	items := ref.Items.Item
	t.Logf("MenuItems: TotalItemCount=%s, received=%d", ref.TotalItemCount, len(items))
	if len(items) > 0 {
		first := items[0]
		t.Logf("  First item: Ident=%s Name=%q Price(1)=%s", first.Ident, first.Name, first.GetPrice(1))
	}
}
