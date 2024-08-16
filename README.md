# Auth Service

### Разворачивание и запуск проекта

- Прежде всего надо:

1) Удостовериться, что установлена утилита `make`.  
2) Удостовериться, что установлен `docker`
2) Запустить `make bin-deps`

- Запуск приложения

1) `make compose-up` - запустить контейнеры
2) `make run-app` - запустить приложение (локально)

- Остановка контейнеров

`make compose-down`

### Workflow

- Получить информацию по командам: `make help`

- Отформатировать код

`make linter-fix` - линтер исправит все, что может, и выведет лог с пофикшенными ошибками

- Просмотр содержимого БД

1) `docker ps` - находим Container ID у образа image
2) `docker exec -it 2089ce2ba4be bash` - вставляем нужный Container ID
3) `psql -p 5432 user -d postgres` - открывает консоль Postgres
4) `exit` - чтобы выйти из контейнера

- Работа с миграция:

1) `make run-app` - автоматически запускает миграции
2) `make migrate-create name="migration_name"` - создать файлы для миграций
3) `make migrate-up` - запустить миграцию
