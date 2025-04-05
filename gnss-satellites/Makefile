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
    ACTIVATE = $(VENV)/Scripts/activate
else
    PYTHON = $(VENV)/bin/python
    PIP = $(VENV)/bin/pip
    ACTIVATE = . $(VENV)/bin/activate
endif

.PHONY: setup install run format lint clean help

# Установка виртуального окружения и зависимостей
setup:
ifeq ($(SYSTEM),Windows)
	python -m venv $(VENV)
else
	python3 -m venv $(VENV)
endif
	@$(MAKE) install

# Установка зависимостей
install:
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt

# Запуск сервера FastAPI
run:
	$(PYTHON) -m uvicorn app.main:app --reload

# Установка pre-commit
pre-commit:
	$(PYTHON) -m pre_commit install

# Форматирование кода
format:
	$(PYTHON) -m black .
	$(PYTHON) -m ruff format .

# Проверка стиля кода
lint:
	$(PYTHON) -m ruff check .
	$(PYTHON) -m mypy .

# Запуск тестов (пример)
test:
	$(PYTHON) -m pytest

# Очистка проекта
clean:
	rm -rf $(VENV)
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete

# Помощь
help:
	@echo "Доступные команды:"
	@echo "  make setup       - Установка виртуального окружения и зависимостей"
	@echo "  make install     - Установка зависимостей"
	@echo "  make run         - Запуск сервера FastAPI"
	@echo "  make test        - Запуск тестов"
	@echo "  make pre-commit  - Установка pre-commit хуков"
	@echo "  make format      - Форматирование кода"
	@echo "  make lint        - Проверка стиля кода"
	@echo "  make clean       - Очистка проекта"
	@echo "  make help        - Показать эту справку"