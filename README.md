# To-Do List API

## Описание:
Этот проект представляет собой REST API для управления задачами (To-Do List), написанный на Go.

___

### Запуск:
Требования:
- Go 1.16+
- PostgreSQL

### Установка:

Клонируйте репозиторий:

git clone https://github.com/Bars-377/Golang_REST_API.git

Настройте базу данных, используя файл base.sql

```
CREATE DATABASE todo_list_db;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

Запустите проект:

go run main.go

API будет доступен по адресу http://localhost:8080


Эндпоинты API

Создание новой задачи

Метод: POST

URL: /tasks

Описание: Создать новую задачу.

Запрос:

- Заголовки:

- Content-Type: application/json

- Тело:
```
{
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)"
}
```
Ответ:

Успех (201 Created):

```
{
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
}
```

Ошибка (400 Bad Request): Неправильный формат данных.

Ошибка (500 Internal Server Error): Проблема на сервере.

Просмотр списка задач

Метод: GET

URL: /tasks

Описание: Получить список всех задач.

Запрос:

- Заголовки:

  Content-Type: application/json

Ответ:

Успех (200 OK):

```
[
    {
        "id": "int",
        "title": "string",
        "description": "string",
        "due_date": "string (RFC3339 format)",
        "created_at": "string (RFC3339 format)",
        "updated_at": "string (RFC3339 format)"
    }
]
```

Ошибка (500 Internal Server Error): Проблема на сервере.

Просмотр задачи

Метод: GET

URL: /tasks/{id}

Описание: Получить задачу по ID.

Запрос:

- Параметры пути:

    id: ID задачи (int)

- Заголовки:

  Content-Type: application/json

Ответ:

Успех (200 OK):

```
{
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
}
```

Ошибка (404 Not Found): Задача не найдена.

Ошибка (500 Internal Server Error): Проблема на сервере.

Обновление задачи

Метод: PUT

URL: /tasks/{id}

Описание: Обновить задачу по ID.

Запрос:

- Параметры пути:

  id: ID задачи (int)

- Заголовки:

  Content-Type: application/json

Тело:

```
{
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)"
}
```

Ответ:

Успех (200 OK):

```
{
    "id": "int",
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)",
    "created_at": "string (RFC3339 format)",
    "updated_at": "string (RFC3339 format)"
}
```

Ошибка (400 Bad Request): Неправильный формат данных.

Ошибка (404 Not Found): Задача не найдена.

Ошибка (500 Internal Server Error): Проблема на сервере.

Удаление задачи

Метод: DELETE

URL: /tasks/{id}

Описание: Удалить задачу по ID.

Запрос:

- Параметры пути:

  id: ID задачи (int)

- Заголовки:

  Content-Type: application/json

Ответ:

Успех (204 No Content): Задача удалена.

Ошибка (404 Not Found): Задача не найдена.

Ошибка (500 Internal Server Error): Проблема на сервере.