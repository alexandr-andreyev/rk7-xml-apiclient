# Changelog

## Unreleased

### Fixed
- `GetRefData`, `GetCateglist`, `GetOrderMenu` — исправлено глотание ошибки: методы теперь возвращают `err` из `do()` вместо `nil`, также удалён закомментированный `defer resp.Body.Close()`
- `request.go` — убраны все `fmt.Println` из библиотечного кода; закомментированный отладочный блок удалён; `errors.New(fmt.Sprintf(...))` заменён на `fmt.Errorf`
- `getRefData.go`, `categList.go`, `orderMenu.go`, `getSystemInfo.go` — все receivers приведены к `*Client`
- `types.go` — опечатка `OnlyActrive` исправлена на `OnlyActive` (обновлены все вызовы в `getRefData.go`, `categList.go`)
- `types.go` — убран хардкод `PRICETYPES-3` из тега поля `Price`; поле удалено, добавлен метод `Item.GetPrice(priceType int) string`

### Breaking changes
- `RK7QueryResult` — удалено поле `CommandResult` (никогда не заполнялось, т.к. сервер не оборачивает ответ в `<CommandResult>`); добавлены поля `CMD`, `ErrorText`, `DateTime`, `WorkTime`, `SystemInfo`, `RK7Reference` напрямую; обновлён `main.go`
- `RK7Reference.TotalItemCount` — исправлен XML-тег: `xml:"TotalItemCount"` → `xml:"TotalItemCount,attr"`
- Добавлены тесты (`types_test.go`): unmarshal SelectorGroups, unmarshal SystemInfo, marshal RK7Command, omitempty, GetPrice
