# analytics-service
Сервис для сбора и хранение информации о действиях совершенных пользователями

## Подключение
- Скопировать .env.dist в корень проекта и переименовать в .env
- В .env изменить значения по необходимости
- Запустить докер из корневой папки проекта
```docker-compose up -d```

С данными из .env.dist, запрос проходит по `http://localhost:8008/analitycs`
Пример POST запроса:
```
curl -location -request POST 'http://localhost:8080/analitycs' \
--header 'X-Tantum-UserAgent: DeviceID=G1752G75-7C56-4G49-BGFA-
5ACBGC963471;DeviceType=iOS;OsVersion=15.5;AppVersion=4.3 (725)' \
--header 'X-Tantum-Authorization: 2daba111-1e48-4ba1-8753-2daba1119a09' \
--header 'Content-Type: application/json' \
--data-raw '{
"module" : "settings",
"type" : "alert",
"event" : "click",
"name" : "подтверждение выхода",
"data" : {"action" : "cancel"}
}'
```
Получить данные можно по `http://localhost:8008/analitycs` или `http://localhost:8008/analitycs/{id}`, где id - идентификатор необходимой записи
Пример записи: 
```json
{
        "id": "1",
        "user_id": "2daba111-1e48-4ba1-8753-2daba1119a09",
        "CreatedAt": "2023-07-14T09:38:46.647082Z",
        "data": {
            "body": {
                "data": {
                    "action": "cancel"
                },
                "event": "click",
                "module": "settings",
                "name": "подтверждение выхода",
                "type": "alert"
            },
            "headers": {
                "Accept": [
                    "*/*"
                ],
                "Accept-Encoding": [
                    "gzip, deflate, br"
                ],
                "Cache-Control": [
                    "no-cache"
                ],
                "Connection": [
                    "keep-alive"
                ],
                "Content-Length": [
                    "146"
                ],
                "Content-Type": [
                    "application/json"
                ],
                "Postman-Token": [
                    "3375701a-c8c0-4daa-a837-8173e69d35b0"
                ],
                "User-Agent": [
                    "PostmanRuntime/7.32.3"
                ],
                "X-Tantum-Authorization": [
                    "2daba111-1e48-4ba1-8753-2daba1119a09"
                ],
                "X-Tantum-Useragent": [
                    "DeviceID=G1752G75-7C56-4G49-BGFA-"
                ]
            }
        }
    }
```
