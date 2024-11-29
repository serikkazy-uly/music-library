# Music Library API 🎶

## Описание
Сервис для работы с библиотекой песен. Поддерживаются следующие возможности:
- Получение списка песен с фильтрацией и пагинацией.
- Получение текста песни по куплетам.
- Добавление новой песни с обогащением данных через внешний API.
- Редактирование и удаление песен.
- Сохранение данных в PostgreSQL.

## Технологии
- **Go**: 
- **PostgreSQL**:
- **Docker** и **Docker Compose**: 
- **Swagger**: 

## Установка и запуск

1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/serikkazy-uly/music-library.git
   cd /music-library
   

2. Сборка и запуск контейнеров с помощью Makefile:
   ```make run
   
3. без make, с использованием docker compose:
   ```docker-compose up --build

4. Для выполнения миграций базы данных: 
   ```make migrate

5. Для генерации документации Swagger:
   ```make swagger

# Swagger UI будет доступен по адресу: http://localhost:8080/swagger/