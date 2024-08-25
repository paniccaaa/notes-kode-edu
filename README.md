# Notes-kode-edu

This repository contains source code for the test-task (Kode company)

## Stack:

- Golang
- PostrgeSQL
- Docker
- [go-chi](https://github.com/go-chi/chi) (routing) / slog (logging) / [cleanenv](https://github.com/ilyakaznacheev/cleanenv) (env config reader) / [jwt](https://github.com/golang-jwt/jwt)
- Yandex Speller [API](https://yandex.ru/dev/speller/) (integration for validating notes)

## Endpoints:
- **GET /notes**
- **POST /notes**
- **POST /login**
- **POST /register**

## Local setup

1. **Clone the repository:**

```bash
git clone https://github.com/paniccaaa/notes-kode-edu.git
cd notes-kode-edu
go mod tidy
```

2. **Start containers**
```bash
docker-compose up -d
```

3. **Migration db** (see [Makefile](https://github.com/paniccaaa/notes-kode-edu/blob/main/Makefile) and tool [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)) 
```bash
make migration_up
```

4. **Run app locally**
```bash
make run
```

5. **Templates curl requests**
- **POST /register**: 
```bash
❯ curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"email": "testuser1@gmail.com", "password": "password1"}'
```
Response:
```bash
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyMUBnbWFpbC5jb20iLCJleHAiOjE3MjQ1OTc3NzksInVpZCI6MX0.XgGZ0Z7Uc-sHv2bw7w9WjARnT1sEso0M-jg9DUXclqo",
  "user": {
    "ID": 1,
    "Email": "testuser1@gmail.com",
    "PassHash": null
  }
}
```
- **POST /login**:
```bash
❯ curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"email": "testuser1@gmail.com", "password": "password1"}'
```
Response
```bash
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyMUBnbWFpbC5jb20iLCJleHAiOjE3MjQ1OTc5MjIsInVpZCI6MX0.ELm5_HI0jWuPoOezKnGLfdZ_Sfbh_KOSSKsZY5_unvE",
  "message": "User successfully logged in"
}
```
- **POST /notes** (example validation Yandex Speller)
```bash
❯ curl -X POST http://localhost:8080/notes \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyMUBnbWFpbC5jb20iLCJleHAiOjE3MjQ1OTc5MjIsInVpZCI6MX0.ELm5_HI0jWuPoOezKnGLfdZ_Sfbh_KOSSKsZY5_unvE" \
     -d '{
           "title": "Мой загаловок",
           "description": "Это опесание новой заметки с ошибкой"
         }'
```

Response
```bash
{
  "note": {
    "id": 0,
    "user_id": 0,
    "title": "Мой загаловок",
    "description": "Это опесание новой заметки с ошибкой"
  },
  "errors": {
    "Мой загаловок": [
      {
        "code": 1,
        "pos": 4,
        "len": 9,
        "word": "загаловок",
        "s": [
          "заголовок"
        ]
      }
    ],
    "Это опесание новой заметки с ошибкой": [
      {
        "code": 1,
        "pos": 4,
        "len": 8,
        "word": "опесание",
        "s": [
          "описание",
          "опасение"
        ]
      }
    ]
  }
}
```
- **POST /notes** (without errors)
```bash
❯ curl -X POST http://localhost:8080/notes \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyMUBnbWFpbC5jb20iLCJleHAiOjE3MjQ1OTc5MjIsInVpZCI6MX0.ELm5_HI0jWuPoOezKnGLfdZ_Sfbh_KOSSKsZY5_unvE" \
     -d '{
           "title": "Мой заголовок без ошибок",
           "description": "Это описание новой заметки без ошибки"
         }'
```

Response
```bash
{
  "id": 1,
  "user_id": 1,
  "title": "Мой заголовок без ошибок",
  "description": "Это описание новой заметки без ошибки"
}
```
- **GET /notes**
```bash
❯ curl -X GET http://localhost:8080/notes \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyMUBnbWFpbC5jb20iLCJleHAiOjE3MjQ1OTc5MjIsInVpZCI6MX0.ELm5_HI0jWuPoOezKnGLfdZ_Sfbh_KOSSKsZY5_unvE"
```

Response
```bash
[
  {
    "id": 1,
    "user_id": 1,
    "title": "Мой заголовок без ошибок",
    "description": "Это описание новой заметки без ошибки"
  },
  {
    "id": 2,
    "user_id": 1,
    "title": "Мой второй заголовок без ошибок",
    "description": "Это описание второе"
  }
]
```