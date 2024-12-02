# Тестовое задание

## Эндпоинты

**Port:** 3030

### 1. Выдача токенов

**Метод:** `POST`
**URL:** `api/auth/token`
**Body:**
```json
{
    "ip_address":"1.1.1.1",
    "refresh_token":"",
    "guid":"GUID"
}
```

### 2. Refresh операция

**Метод:** `POST`
**URL:** `api/auth/refresh`
**Body:**
```json
{
    "ip_address":"1.1.1.1",
    "refresh_token":"refresh token",
    "guid":"GUID"
}
```

## Запуск

### Запуск docker compose

```bash
docker-compose -f ./docker/docker-compose.yml up -d
```

### Добавление сервера в pgAdmin

* **Name:** auth service 
* **Host name/address:** postgres_container
* **Port:** 5432
* **Maintenance database:** auth_service 
* **User:** pguser 
* **Password:** 1212

### Включение отправки предупреждений на почту

Для включение отправки предупреждений на почту необходимо:

* Прописать в config.env `SMTP_ENABLE="on"`