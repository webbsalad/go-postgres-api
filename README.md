# Проект REST API на Go для управления базой данных Postgres

Этот проект представляет собой REST API, написанный на языке Go, который предназначен для управления базой данных Postgres. Он использует следующие библиотеки:

- [github.com/jackc/pgx/v4](https://github.com/jackc/pgx/v4)
- [github.com/gofiber/fiber/v2](https://github.com/gofiber/fiber/v2)
- [github.com/gofiber/adaptor/v2](https://github.com/gofiber/adaptor/v2)
- [github.com/ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)

Проект также развертывается на [Vercel](https://vercel.com/) с использованием файла `vercel.json`, который настроен следующим образом:

```json
{
    "rewrites": [
      { "source": "(.*)", "destination": "api/index.go" }
    ]
}
```

## Секретные данные
Для корректной работы с базой данных и интеграции Vercel с GitHub Secrets, необходимо добавить следующие переменные в GitHub Secrets:
```plaintext
VERCEL_TOKEN
DB_PORT
DB_HOST
DB_NAME
DB_USER
DB_PASSWORD
```
(при локальной работе нужно заменить функцию в файле /config/config.go а данные записать в файл .env)

## API позволяет управлять таблицами разными запросами (POST, GET, DELETE, PATCH) по адресам в следующем формате:

`/:table_name/` - Обращение к таблице table_name.

`/:table_name/1` - Обращение к элементу с индексом 1 в таблице table_name.

`/:table_name?name=*&sortBy=name` - Выводит только элементы с именем * и сортирует их по имени.


## Пример запроса:
`/table_name/1/`  - обращение к таблице table_name элементу с индексом 1

`/table_name?name=`  - обращение к таблице table_name к элементом с определенным именем ( в случае * это любое имя)

`/table_name?sortBy=name`  - отсортирует по имени
>(вы можете комбинировать например `/table_name/1?name=*&sortBy=name`)


## Структура проекта
```plaintext
Project:.
├───api
│   └───index.go   // Главный файл API
├───config
│   └───config.go   // Конфиг базы данных
├───db
│   └───operations   // Файлы с операциями с базой данных
└───routers   // Файлы с операциями на серверной части
└───utils   // Вспомогательные функции
```


