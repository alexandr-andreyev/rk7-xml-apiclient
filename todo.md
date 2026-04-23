# TODO

## Критично
- [x] `getRefData.go`, `categList.go`, `orderMenu.go` — исправить глотание ошибки: `return &result, nil` → `return &result, err`
- [ ] `rk7client.go:25` — убрать `InsecureSkipVerify: true`, добавить опцию или инъекцию `*http.Client`

## Высокий приоритет
- [x] `request.go` — убрать `fmt.Println` из библиотечного кода (ошибки и так возвращаются)
- [x] Все файлы — привести receiver всех методов к `*Client` (сейчас смешаны `Client` и `*Client`)
- [x] Добавить тесты: минимум XML-маршалинг/анмаршалинг без сетевых вызовов

## Средний приоритет
- [x] `types.go` — переименовать поле `OnlyActrive` → `OnlyActive`
- [x] `types.go:108` — убрать хардкод `PRICETYPES-3` в теге поля `Price`

## Низкий приоритет
- [x] Удалить весь закомментированный код (`request.go`, `getRefData.go`, `categList.go`, `orderMenu.go`)
- [x] `types.go:91` — проверить `TotalItemCount`: атрибут или элемент?
- [ ] `types.go:61` — доработать или удалить неиспользуемый stub `Dishes`
