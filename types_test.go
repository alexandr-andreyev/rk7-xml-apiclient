package rk7client

import (
	"encoding/xml"
	"strings"
	"testing"
)

const xmlSelectorGroups = `<?xml version="1.0" encoding="utf-8"?>
<RK7QueryResult ServerVersion="7.25.11.2001" XmlVersion="248" NetName="DEMOSERVER"
	Status="Ok" Processed="1" ArrivalDateTime="2026-04-22T17:56:03">
	<CommandResult CMD="GetRefData" Status="Ok" ErrorText="" DateTime="2026-04-22T17:56:03" WorkTime="0">
		<RK7Reference DataVersion="133" TotalItemCount="3">
			<Items>
				<Item Ident="40" Code="40" Name="Дерево Меню" Status="rsActive" Parent="912" MainParentIdent="912" GUIDString="{5DA73247-3416-4A75-82D8-20DB3E370507}">
					<RIChildItems>
						<TSelector ItemIdent="11048" Ident="11048" Name="Начальный выбор" Status="rsActive"/>
						<TSelector ItemIdent="1000439" Ident="1000439" Name="Меню для киоска" Status="rsActive">
							<RIChildItems>
								<TSelectorDetail ItemIdent="1000442" TradeObject="toSelector" ObjectSifr="1000440" OrderNum="231428"/>
								<TSelectorDetail ItemIdent="1000461" TradeObject="toSelector" ObjectSifr="1000449" OrderNum="231458"/>
							</RIChildItems>
						</TSelector>
						<TSelector ItemIdent="1000440" Ident="1000440" Name="Популярные блюда" Status="rsActive">
							<RIChildItems>
								<TSelectorDetail ItemIdent="1000443" TradeObject="toDish" ObjectSifr="1000059" OrderNum="231448"/>
							</RIChildItems>
						</TSelector>
					</RIChildItems>
				</Item>
			</Items>
		</RK7Reference>
	</CommandResult>
</RK7QueryResult>`

const xmlSystemInfo = `<?xml version="1.0" encoding="utf-8"?>
<RK7QueryResult ServerVersion="7.25.11.2001" XmlVersion="248" NetName="DEMOSERVER"
	Status="Ok" Processed="1" ArrivalDateTime="2026-04-22T17:57:31">
	<CommandResult CMD="GetSystemInfo" Status="Ok" ErrorText="" DateTime="2026-04-22T17:57:31" WorkTime="0">
		<SystemInfo SystemTime="3986042251828" ReqSysVer="2" ProcessID="10020" ShiftDate="20260415" RestCode="331950001">
			<BusinessPeriod id="0" code="0"/>
			<CashGroup id="15001" code="15001" name="DEMOSERVER"/>
			<Restaurant id="1" code="1" name="Техлайн"/>
		</SystemInfo>
	</CommandResult>
</RK7QueryResult>`

func TestUnmarshal_SelectorGroups(t *testing.T) {
	var result RK7QueryResult
	if err := xml.Unmarshal([]byte(xmlSelectorGroups), &result); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if result.Status != "Ok" {
		t.Errorf("Status: got %q, want %q", result.Status, "Ok")
	}
	if len(result.CommandResult) == 0 {
		t.Fatal("CommandResult is empty")
	}
	cr := result.CommandResult[0]
	if cr.CMD != "GetRefData" {
		t.Errorf("CMD: got %q, want %q", cr.CMD, "GetRefData")
	}
	if len(cr.Data) == 0 {
		t.Fatal("CommandResult.Data is empty")
	}
	ref, ok := cr.Data[0].(RK7Reference)
	if !ok {
		t.Fatalf("Data[0] is not RK7Reference, got %T", cr.Data[0])
	}
	if ref.DataVersion != "133" {
		t.Errorf("DataVersion: got %q, want %q", ref.DataVersion, "133")
	}
	if ref.TotalItemCount != "3" {
		t.Errorf("TotalItemCount: got %q, want %q", ref.TotalItemCount, "3")
	}

	items := ref.Items.Item
	if len(items) != 1 {
		t.Fatalf("items count: got %d, want 1", len(items))
	}
	item := items[0]
	if item.Ident != "40" {
		t.Errorf("Item.Ident: got %q, want %q", item.Ident, "40")
	}
	if item.Name != "Дерево Меню" {
		t.Errorf("Item.Name: got %q, want %q", item.Name, "Дерево Меню")
	}

	if len(item.Selectors) != 3 {
		t.Fatalf("Selectors count: got %d, want 3", len(item.Selectors))
	}

	s0 := item.Selectors[0]
	if s0.Ident != "11048" {
		t.Errorf("Selector[0].Ident: got %q, want %q", s0.Ident, "11048")
	}
	if s0.Name != "Начальный выбор" {
		t.Errorf("Selector[0].Name: got %q, want %q", s0.Name, "Начальный выбор")
	}
	if len(s0.Children) != 0 {
		t.Errorf("Selector[0].Children: got %d, want 0", len(s0.Children))
	}

	s1 := item.Selectors[1]
	if len(s1.Children) != 2 {
		t.Fatalf("Selector[1].Children count: got %d, want 2", len(s1.Children))
	}
	d := s1.Children[0]
	if d.TradeObject != "toSelector" {
		t.Errorf("Detail[0].TradeObject: got %q, want %q", d.TradeObject, "toSelector")
	}
	if d.ObjectSifr != "1000440" {
		t.Errorf("Detail[0].ObjectSifr: got %q, want %q", d.ObjectSifr, "1000440")
	}
	if d.OrderNum != "231428" {
		t.Errorf("Detail[0].OrderNum: got %q, want %q", d.OrderNum, "231428")
	}

	s2 := item.Selectors[2]
	if len(s2.Children) != 1 {
		t.Fatalf("Selector[2].Children count: got %d, want 1", len(s2.Children))
	}
	if s2.Children[0].TradeObject != "toDish" {
		t.Errorf("Detail.TradeObject: got %q, want %q", s2.Children[0].TradeObject, "toDish")
	}
}

func TestUnmarshal_SystemInfo(t *testing.T) {
	var result RK7QueryResult
	if err := xml.Unmarshal([]byte(xmlSystemInfo), &result); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if len(result.CommandResult) == 0 {
		t.Fatal("CommandResult is empty")
	}
	cr := result.CommandResult[0]
	if cr.CMD != "GetSystemInfo" {
		t.Errorf("CMD: got %q, want %q", cr.CMD, "GetSystemInfo")
	}
	if len(cr.Data) == 0 {
		t.Fatal("CommandResult.Data is empty")
	}
	si, ok := cr.Data[0].(SystemInfo)
	if !ok {
		t.Fatalf("Data[0] is not SystemInfo, got %T", cr.Data[0])
	}
	if si.ShiftDate != "20260415" {
		t.Errorf("ShiftDate: got %q, want %q", si.ShiftDate, "20260415")
	}
	if si.ProcessID != "10020" {
		t.Errorf("ProcessID: got %q, want %q", si.ProcessID, "10020")
	}
	if si.CashGroup.Id != "15001" {
		t.Errorf("CashGroup.Id: got %q, want %q", si.CashGroup.Id, "15001")
	}
	if si.Restaurant.Name != "Техлайн" {
		t.Errorf("Restaurant.Name: got %q, want %q", si.Restaurant.Name, "Техлайн")
	}
}

func TestMarshal_RK7Command_GetRefData(t *testing.T) {
	q := RK7Query{
		RK7Command: []RK7Command{{
			CMD:            "GetRefData",
			RefName:        "SelectorGroups",
			WithChildItems: "2",
			WithMacroProp:  "1",
			RefItemIdent:   "40",
		}},
	}
	out, err := xml.Marshal(q)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(out)
	for _, want := range []string{
		`CMD="GetRefData"`,
		`RefName="SelectorGroups"`,
		`WithChildItems="2"`,
		`WithMacroProp="1"`,
		`RefItemIdent="40"`,
	} {
		if !strings.Contains(s, want) {
			t.Errorf("marshaled XML missing %q\ngot: %s", want, s)
		}
	}
}

func TestMarshal_RK7Command_OmitEmpty(t *testing.T) {
	q := RK7Query{
		RK7Command: []RK7Command{{CMD: "GetSystemInfo"}},
	}
	out, err := xml.Marshal(q)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(out)
	for _, absent := range []string{"RefName", "OnlyActive", "WithChildItems", "WithMacroProp", "PropMask", "RefItemIdent"} {
		if strings.Contains(s, absent) {
			t.Errorf("marshaled XML should not contain %q\ngot: %s", absent, s)
		}
	}
}

func TestItem_GetPrice(t *testing.T) {
	item := Item{
		Attributes: []xml.Attr{
			{Name: xml.Name{Local: "PRICETYPES-3"}, Value: "250.00"},
			{Name: xml.Name{Local: "PRICETYPES-5"}, Value: "300.00"},
		},
	}
	if got := item.GetPrice(3); got != "250.00" {
		t.Errorf("GetPrice(3): got %q, want %q", got, "250.00")
	}
	if got := item.GetPrice(5); got != "300.00" {
		t.Errorf("GetPrice(5): got %q, want %q", got, "300.00")
	}
	if got := item.GetPrice(99); got != "" {
		t.Errorf("GetPrice(99): got %q, want %q", got, "")
	}
}

func TestMarshal_RK7Command_RefItemIdent(t *testing.T) {
	q := RK7Query{
		RK7Command: []RK7Command{{
			CMD:            "GetRefData",
			RefName:        "SelectorGroups",
			WithChildItems: "2",
			WithMacroProp:  "1",
			RefItemIdent:   "40",
		}},
	}
	out, err := xml.Marshal(q)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, `RefItemIdent="40"`) {
		t.Errorf("marshaled XML missing RefItemIdent\ngot: %s", s)
	}
	if strings.Contains(s, "OnlyActive") {
		t.Errorf("marshaled XML should not contain OnlyActive\ngot: %s", s)
	}
}

func TestMarshal_RK7Command_Station(t *testing.T) {
	q := RK7Query{
		RK7Command: []RK7Command{{
			CMD:     "GetOrderMenu",
			Station: &Station{Code: "1"},
		}},
	}
	out, err := xml.Marshal(q)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, `<Station`) {
		t.Errorf("marshaled XML missing Station element\ngot: %s", s)
	}
}

const xmlRefDataWithPrices = `<?xml version="1.0" encoding="utf-8"?>
<RK7QueryResult ServerVersion="7.25.11.2001" XmlVersion="248" NetName="DEMOSERVER"
	Status="Ok" Processed="1" ArrivalDateTime="2026-04-22T17:56:03">
	<CommandResult CMD="GetRefData" Status="Ok" ErrorText="" DateTime="2026-04-22T17:56:03" WorkTime="0">
		<RK7Reference DataVersion="200" TotalItemCount="1">
			<Items>
				<Item Ident="1000059" GUIDString="{AAA}" Code="1059" Name="Борщ" Status="rsActive"
					Parent="100" MainParentIdent="100" PRICETYPES-3="250.00" PRICETYPES-5="300.00"/>
			</Items>
		</RK7Reference>
	</CommandResult>
</RK7QueryResult>`

func TestUnmarshal_RefData_WithPriceTypes(t *testing.T) {
	var result RK7QueryResult
	if err := xml.Unmarshal([]byte(xmlRefDataWithPrices), &result); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if len(result.CommandResult) == 0 {
		t.Fatal("CommandResult is empty")
	}
	cr := result.CommandResult[0]
	if len(cr.Data) == 0 {
		t.Fatal("CommandResult.Data is empty")
	}
	ref, ok := cr.Data[0].(RK7Reference)
	if !ok {
		t.Fatalf("Data[0] is not RK7Reference, got %T", cr.Data[0])
	}
	items := ref.Items.Item
	if len(items) != 1 {
		t.Fatalf("items count: got %d, want 1", len(items))
	}
	item := items[0]
	if item.Name != "Борщ" {
		t.Errorf("Item.Name: got %q, want %q", item.Name, "Борщ")
	}
	if got := item.GetPrice(3); got != "250.00" {
		t.Errorf("GetPrice(3): got %q, want %q", got, "250.00")
	}
	if got := item.GetPrice(5); got != "300.00" {
		t.Errorf("GetPrice(5): got %q, want %q", got, "300.00")
	}
}
