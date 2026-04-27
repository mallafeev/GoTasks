Малафеев Леонид

Тестовое задание на Golang

---

## Описание проекта

REST API сервис управления группами людей.

Функционал:
- CRUD групп (с иерархией)
- CRUD людей (Имя, Фамилия, Год рождения)
- Привязка человека к группе
- Просмотр групп с количеством участников
- Просмотр людей в группе (только группа / включая дочерние)

---

## Подготовительные действия

Установлено:
- Docker Desktop
- VS Code
- Git (опционально)

---

## Как запустить проект

git clone <repos>
cd task9
docker-compose up --build

# API
Примеры запросов API
Создание группы
curl -X POST http://localhost:8080/groups \
-H "Content-Type: application/json" \
-d "{\"name\":\"Engineering\",\"parent_id\":null}"
Обновление группы
curl -X PUT http://localhost:8080/groups/1 \
-H "Content-Type: application/json" \
-d "{\"name\":\"New Engineering\",\"parent_id\":null}"
Удаление группы
curl -X DELETE http://localhost:8080/groups/1
Создание человека
curl -X POST http://localhost:8080/people \
-H "Content-Type: application/json" \
-d "{\"first_name\":\"Ivan\",\"last_name\":\"Ivanov\",\"birth_year\":1990,\"group_id\":2}"
Обновление человека (включая смену группы)
curl -X PUT http://localhost:8080/people/1 \
-H "Content-Type: application/json" \
-d "{\"first_name\":\"Petr\",\"last_name\":\"Petrov\",\"birth_year\":1995,\"group_id\":3}"
Удаление человека
curl -X DELETE http://localhost:8080/people/1
Работа с людьми в группах
Получить людей только из конкретной группы
curl http://localhost:8080/groups/1/people?mode=direct
Получить людей из группы и всех дочерних групп
curl http://localhost:8080/groups/1/people?mode=all
Получить список групп
curl http://localhost:8080/groups

