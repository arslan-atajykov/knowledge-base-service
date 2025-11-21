# Knowledge Base Service

Небольшой API-сервис knowledge-base, который работает как банк вопросов и ответов.  
К одному вопросу может быть несколько ответов.  
Сервис реализован на Go, использует PostgreSQL, миграции через goose и запускается полностью через Docker Compose.

## Стек технологий

- Go (net/http)
- GORM
- PostgreSQL
- Goose — миграции
- Docker + docker-compose
- httptest — один юнит-тест

---

# Запуск проекта

Полный запуск (backend + PostgreSQL):

```
docker-compose up --build
```

После запуска сервис доступен по адресу:
```
http://localhost:8080
```

---

# Примеры запросов (cURL)

1. Создать вопрос:

```
curl -X POST http://localhost:8080/questions   -H "Content-Type: application/json"   -d '{"text": "Столица Англии?"}'
```

Ответ:
```
{
  "id": 1,
  "text": "Столица Англии?",
  "created_at": "2025-11-21T15:34:56Z"
}
```

2. Получить список всех вопросов:

```
curl http://localhost:8080/questions
```

3. Получить вопрос по ID (с ответами):

```
curl http://localhost:8080/questions/1
```

4. Добавить ответ к вопросу:

```
curl -X POST http://localhost:8080/questions/1/answers   -H "Content-Type: application/json"   -d '{"user_id": "u123", "text": "London"}'
```

5. Получить ответ по ID:

```
curl http://localhost:8080/answers/1
```

6. Удалить вопрос (удаляются все ответы):

```
curl -X DELETE http://localhost:8080/questions/1
```

7. Удалить ответ:

```
curl -X DELETE http://localhost:8080/answers/1
```

---

# Тесты

В проекте добавлен один HTTP-тест.

Запуск тестов:

```
go test ./...
```

Тест проверяет создание вопроса через POST /questions.

---
