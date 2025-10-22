# Go Task Tracker

Простой трекер задач на Go с использованием Gin и SQLite.

## Описание

Проект представляет собой REST API для управления списком задач.  
Позволяет создавать новые задачи и получать список существующих. Все данные хранятся в SQLite базе данных (`tasks.db`).


## Установка и запуск

1. Клонируйте репозиторий:

```bash
git clone git@github.com:provsenya/go-task-tracker.git
cd go-task-tracker
```

2. Установите зависимости:

```bash
go mod tidy
```

3. Запустите сервер:

```bash
go run .
```

По умолчанию сервер запускается на http://localhost:8080.

# API
## GET/tasks

Возвращает список всех задач.
Пример запроса:
```bash
curl http://localhost:8080/tasks
```

Пример ответа:
```bash
[
  {"id":1,"title":"Buy milk","completed":false},
  {"id":2,"title":"Buy milk","completed":false}
]
```

## POST/tasks

Добавляет новую задачу.
Пример запроса:
```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Buy milk"}'
```

Пример ответа:
```bash
{
  "status": "task added"
}
```

# База данных

```SQLite файл tasks.db создаётся автоматически при первом запуске.```

## Таблица tasks содержит следующие поля:

- id — уникальный идентификатор задачи

- title — текст задачи

- completed — логическое поле, отмечает, выполнена задача или нет
