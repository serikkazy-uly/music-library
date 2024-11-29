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
   ```bash
   make run
   
3. без make, с использованием docker compose:
   ```bash
   docker-compose up --build

4. Для выполнения миграций базы данных: 
   ```bash
   make migrate

5. Для генерации документации Swagger:
   ```bash
   make swagger

> Swagger UI будет доступен по адресу: http://localhost:8080/swagger/

### Подключение к внешнему API для получения информации о песнях

1. **Внесите изменения в код для подключения к реальному API**:

   Откройте файл `services.go` и найдите функцию `FetchSongDetailsFromAPI` замените URL на реальный API, например, [API Songs](https://api.example.com/info).

   Пример:
   ```go
   func FetchSongDetailsFromAPI(groupName, songName string) (models.Song, error) {
          apiURL := fmt.Sprintf("https://api.example.com/info?group=%s&song=%s", groupName, songName)
          ...
   }
   
2. Измените конфигурацию в .env:

   EXTERNAL_API_URL=https://api.example.com/info

