# Переменные
VENV = .venv

# Определение OS
ifeq ($(OS),Windows_NT)
    SYSTEM = Windows
else
    SYSTEM = Linux
endif

# Установка путей в зависимости от ОС
ifeq ($(SYSTEM),Windows)
    PYTHON = $(VENV)/Scripts/python.exe
    PIP = $(VENV)/Scripts/pip.exe
else
    PYTHON = $(VENV)/bin/python
    PIP = $(VENV)/bin/pip
endif

# Установка виртуального окружения и зависимостей
setup:
	ifeq ($(SYSTEM),Windows)
		python -m venv $(VENV)
	else
		python3 -m venv $(VENV)
	endif
		$(PIP) install --upgrade pip
		$(PIP) install -r requirements.txt

activate:
	ifeq ($(SYSTEM),Windows)
		./venv/Scripts/activate.ps1
	else
		source ./venv/bin/activate
	endif

# Запуск сервера FastAPI
run:
	$(PYTHON) -m uvicorn app.main:app --reload

# Создание пре-коммитов
format:
	$(PYTHON) -m pre-commit install

# Форматирование кода с помощью black и isort
format:
	$(PYTHON) -m black .
	$(PYTHON) -m ruff format .

# Проверка стиля кода с помощью flake8
lint:
	$(PYTHON) -m ruff check .
	$(PYTHON) -m mypy .

# Очистка проекта (удаление виртуального окружения и кэша)
clean:
	rm -rf $(VENV)
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete

# Помощь (список доступных команд)
help:
	@echo "Доступные команды:"
	@echo "  make setup       - Установка виртуального окружения и зависимостей"
	@echo "  make run         - Запуск сервера FastAPI"
	@echo "  make test        - Запуск тестов"
	@echo "  make format      - Форматирование кода"
	@echo "  make lint        - Проверка стиля кода"
	@echo "  make clean       - Очистка проекта"
	@echo "  make help        - Показать эту справку"