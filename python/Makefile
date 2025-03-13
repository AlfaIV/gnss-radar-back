# Переменные
VENV = .venv
PYTHON = $(VENV)/Script/python
PIP = $(VENV)/Script/pip

# Цели (targets)

# Установка виртуального окружения и зависимостей
setup:
	python -m venv $(VENV)
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt

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