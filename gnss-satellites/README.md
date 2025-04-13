# Запуска проекта

Если запуск происходит в __Windows__

```bash
    python -m venv venv
    .\venv\scripts\activate.ps1
    pip install -r requirements.txt
    fastapi.exe dev .\app\main.py
```

Если запуск происходит в __Linux__

```bash
    python3 -m venv venv
    source ./venv/bin/activate
    pip install -r requirements.txt
    fastapi dev ./app/main.py
```

Запуск через Докер образ:

```bash
    docker build -t measurement_services .
    docker run -d -p 8010:8010 --name measurement_services_container measurement_services —network gnss-radar-net
```

Для формирования пре-коммитов можно использовать:

```bash
    pre-commit install
```
