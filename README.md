# Running the Application

To run this application, please follow these steps:

```bash
git clone https://github.com/Wefdzen/medods.git
cd medods
docker-compose up -d

```

# Swagger

Once the application is running, you can access the Swagger UI at the following address:

http://localhost:8080/ui-swagger/index.html

or use postman

# Дополнение
У меня сейчас пофик умер ли токен или нет он все равно позволит /refresh
Если надо чтобы менять пару можно было только после невалидности accessToken
то просто в ParseJWT в claims убрать else

Про EmailWarning он у меня StubEmailService and RealEmailService 
stub просто пишет log в консоль, как будто было отправлено сообщение.
Про RealEmailService чтобы он заработал надо in refresh_handler.go
83 строка поменять на тип RealEmailService и уже в email_service.go
устновить свою почту и пароль и чтобы gmail заработал надо получить api key.
И тогда сообщение будет отправлено на саму себя.
