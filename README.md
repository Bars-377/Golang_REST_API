# To-Do List API

## Описание
Этот проект представляет собой REST API для управления задачами (To-Do List), написанный на Go.

## Запуск

### Требования
- Go 1.16+
- PostgreSQL

### Установка
1. Клонируйте репозиторий.
2. Настройте базу данных:
   ```sql
   CREATE DATABASE todo_list_db;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

Запустите проект:
go run main.go

API будет доступен по адресу http://localhost:8080

