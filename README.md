# BalanceRange
## Техническое задание 
### Проблема
 - У приложения которое обрабатывает запросы на вывод проблема, оно допускает овердрафт

### Задание
 - Необходимо модернизировать функцию вывода так, чтобы не допускать овердрафт
 - Сервис должен меть возможность запуска в нескольких репликах 
 - Нельзя подключать сторонние библиотеки (за исключением тех, что есть в `go.mod`)

### Справочник
 - Овердрафт - уход баланса клиента в минус
 - Баланс клиента - сумма всех вводов минус сумма всех вывводов

### Обязательное условие
 - Тесты
 - Успешное выполнение команд
    ```bash
    $ make lint
    $ make vet
    ```
 - Запуск проекта через `docker-compose`
   ```bash
   $ docker-compose build
   $ docker-compose up
   ```
 - Соблюдение чистоты архитектуры и чистоты кода