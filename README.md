# Проект REST API на Go для управления базой данных Postgres

Этот проект представляет собой REST API, написанный на языке Go, который предназначен для управления базой данных Postgres. Он использует следующие библиотеки:

- [github.com/jackc/pgx/v4](https://github.com/jackc/pgx/v4)
- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

Проект также развертывается на [Vercel](https://vercel.com/) с использованием файла `vercel.json`, который настроен следующим образом:

```json
{
    "rewrites": [
      { "source": "(.*)", "destination": "api/index.go" }
    ]
}
```

**API позволяет управлять таблицами разными запросами (POST, GET, DELETE, PATCH) по адресам в следующем формате:**

- <u>**/:table_name/**</u> - Обращение к таблице table_name.
- <u>**/:table_name/1**</u> - Обращение к элементу с индексом 1 в таблице table_name.
- <u>**/:table_name?name=*&sortBy=name**</u> - Выводит только элементы с именем * и сортирует их по имени.

**Пример запроса:**
- <u>**/table_name/1/**</u>  - обращение к таблице table_name элементу с индексом 1
- <u>**/table_name?name=**</u>  - обращение к таблице table_name к элементом с определенным именем ( в случае * это любое имя)
- <u>**/table_name?sortBy=name**</u>  - отсортирует по имени
>(вы можете комбинировать например <u>**/table_name/1?name=*&sortBy=name**</u>)


### Структура проекта
```plaintext
Project:.
├───api
│   └───index.go   // Главный файл API
├───config
│   └───config.go   // Конфиг базы данных
├───db
│   └───operations   // Файлы с операциями с базой данных
└───routers   // Файлы с операциями на серверной части
```

