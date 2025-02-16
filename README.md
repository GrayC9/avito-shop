# Avito Shop 🛒

Сервис для внутреннего магазина мерча в Авито. Пользователи могут покупать товары за виртуальные монеты, переводить монеты друг другу, просматривать историю транзакций и список купленных товаров.

## 📌 Функциональность

- 📥 Регистрация и авторизация через JWT
- 💰 Покупка товаров за монеты
- 🔄 Перевод монет между пользователями
- 🛍️ Просмотр списка купленных товаров
- 📜 История транзакций (от кого/кому переведены монеты)
- ❌ Нельзя иметь отрицательный баланс
- 🏪 Магазин содержит 10 видов товаров с фиксированными ценами и бесконечным запасом

## 🚀 Запуск проекта

### 1️⃣ Клонирование репозитория
```sh
git clone git@github.com:GrayC9/avito-shop.git
cd avito-shop
```

### 2️⃣ Запуск с использованием Docker
```sh
docker compose up app —build
```

### Docker-compose.yml
```sh
version: '3.8'

services:
  db:
    image: postgres:15
    container_name: avito-shop-db
    restart: always
    environment:
      POSTGRES_DB: avito_shop
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - pg_data2:/var/lib/postgresql/data

  app:
    build: .
    container_name: avito-shop-app
    restart: always
    depends_on:
      - db
    ports:
      - "8080:8080"
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: admin
      DB_PASSWORD: password
      DB_NAME: avito_shop
    command: ["/app/main"]

volumes:
  pg_data2:
```
### Запуск тестов
```sh
 go test ./...  
```

