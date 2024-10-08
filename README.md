# Auth Service

### Разворачивание и запуск проекта

- Прежде всего надо:

1. Удостовериться, что установлена утилита `make`.
2. Удостовериться, что установлен `docker`.
3. Запустить `make bin-deps`.

- Запуск приложения, для разработки:

1. `make compose-up` - запустить контейнеры.
2. `make run-app` - запустить приложение (локально).

- Запуск приложения, для продакшена:

`make compose-up-prod`

- Остановка контейнеров:

`make compose-down`

### Workflow

- Получить информацию по командам: `make help`.

- Отформатировать код:

`make linter-fix` - линтер исправит все что может и выведет лог с непофикшенными ошибками.

- Просмотр содержимого БД:

1. `docker ps` - находим Container ID у образа image.
2. `docker exec -it 2089ce2ba4be bash` - вставляем нужный Container ID.
3. `psql -p 5432 user -d postgres` - открывает консоль Postgres.
4. `exit` - чтобы выйти из контейнера.

- Работа с миграция:

1. `make run-app` - автоматически запускает миграции.
2. `make migrate-create name="migration_name"` - создать файлы для миграций.
3. `make migrate-up` - запустить миграцию.

- Удалить содержимое БД:

`make docker-rm-volume`

- Unit тесты

1. `make test` - запустить unit тесты
2. `make mock` - запустить кодогенерацию моков
