# Тестовое задание на позицию стажёра-бэкендера
## Микросервис для работы с балансом пользователей
Сервис для взаимодействия с балансом пользователей
Позволяет хранить и начислять деньги пользователям;
Выполнять операции со счетом пользователя(совершение транзакий, замораживание счета);
Получение месячного отчета о совершенных транзакциях

## Используемые технологии:

- [Go](https://go.dev)
- [Gin](https://github.com/gin-gonic/gin) - Фреймворк для работы с http
- [Docker](https://www.docker.com)
- [PostgreSQL](https://www.postgresql.org)
- [Viper](https://github.com/spf13/viper) - Создание и работа с файлами конфигурации
- [sqlx](https://github.com/jmoiron/sqlx) - Библиотека для работа с базой данных
- [sqltocsv](https://github.com/joho/sqltocsv) - Библиотека для работы с csv файлами

## Запуск приложения:
```sh
docker build -t projects .
docker run --name=projects -p 8081:8081 projects
docker-compose up --build projects
```
## Реализованные API:
 - /user
    - /:id - Возвращает данные о пользователе с id=id - GET
    - /:id/balance - Возвращает баланс пользователя с id=id - GET
    - /create-user - Создание пользователя - POST
    - /users - Получение списка всех пользователей - GET
    - /:id - Удалить пользователя с id=id - DELETE
 - /transaction
   - /send - Зачисление баланса на счет пользователя - POST
   - /reserve - Снять деньги со счета пользователя и положить на зарезервированный счет
   - /confirm - Подтверждение операции оплаты и размораживание денег со счета
 - /report/get-report - Получение отчета о транзакциях

# Примеры запросов

##  Пользователи

### Получить пользователя
```sh
GET http://localhost:8081/user/1
```

### Создать пользователя
```sh
POST http://localhost:8081/user/create-user
{
    "id":1,
    "balance":100
}
```

### Вернуть баланс пользователя
```sh
GET http://localhost:8081/user/1/balance
```

### Получить список всех пользователей
```sh
GET http://localhost:8081/user/users
```

### Удалить пользователя
```sh
DELETE http://localhost:8081/user/1
```
## Транзакции
### Зачислить сумму на баланс пользователя
```sh
POST http://localhost:8081/transaction/send
{
    "id":1,
    "balance":100
}
```
### Перевести деньги пользователя на резервный счет
```sh
POST http://localhost:8081/transaction/reserve
{
    "id":1,
    "id_service":1,
    "id_order":1,
    "balance":100
}
```

### Подтвердить транзакцию и снять деньги с резервного счета
```sh
POST http://localhost:8081/transaction/confirm
{
    "id":1,
    "id_service":1,
    "id_order":1,
    "balance":100
}
```
### Положить деньги пользователя на резервный счет
```sh
POST http://localhost:8081/transaction/reserve
{
    "id":1,
    "id_service":1,
    "id_order":1,
    "balance":100
}
```

### Получить отчет о транзакциях
```sh
POST http://localhost:8081/report/get-report?year=2022&month=11
```

