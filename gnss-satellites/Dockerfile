FROM python:3.10-slim
WORKDIR /app
COPY /gnss-satellites .
RUN pip install --no-cache-dir -r requirements.txt
EXPOSE 8010
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8010"]
