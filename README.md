Hello reader.
This is my first project related to a telegram bot that can save, delete, list, output a random page (bookmark).
It is assumed that you will use this project as a template for your projects

How to start my project:
The _exe file contains everything you need to run
- data.json - data to connect to the database, get telegram API and specify batch limit updates
- version-1.0.exe - app
- run.bat - launching an empty console (for convenience)

JSON
- "tgBotHost": "api.telegram.org",
- "PSQLconnection": "user=username dbname=dbname password=pass host=ip port=5432 sslmode=disable",
- "batchSize": 100   // batchSize - updatesBatchLimit, between 1 - 100, defaults to 100


Start
You need to run the application with the -tg-bot-token <token> flag
```
start version-1.0.exe -tg-bot-token <token>
```
