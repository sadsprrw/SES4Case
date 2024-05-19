# SES4 Case

## Опис

Це API сервіс, який дозволяє отримати поточний курс долара США до гривні та підписати емейл на отримання інформації про зміну курсу.

## Запуск проекту

1. Клонуйте репозиторій
2. Виконайте команду `docker-compose up --build`
3. API буде доступне за адресою `http://localhost:8080`

## Використання API

### Отримання поточного курсу
```GET /api/rate```<br>
Використовує third party APi https://currencybeacon.com/api-documentation
### Підписка на отримання курсу
```POST /api/subscribe```
Параметри: email (required, string)
