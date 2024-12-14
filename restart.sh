# Логирование текущего времени
current_time_custom=$(date +"%Y-%m-%d %H:%M:%S restart")
echo "$current_time_custom" >> logs

# Остановка всех контейнеров
docker stop $(docker ps -a -q)

# Удаляем старую директорию и загружаем новый образ
rm -rf gnss-radar
git clone https://github.com/Gokert/gnss-radar.git

cd ./gnss-radar || exit

# Путь к файлу для проверки
file_to_check="./scripts/sql/init/init_db.sql"
hash_file="init_db_hash.txt"

# Вычисляем текущий хеш файла
current_hash=$(md5sum "$file_to_check" | awk '{ print $1 }')

# Проверяем, существует ли файл с сохраненным хешем
if [ -f "$hash_file" ]; then
    saved_hash=$(cat "$hash_file")
    if [ "$current_hash" != "$saved_hash" ]; then
        echo "Файл $file_to_check изменился, пересобираем образ PostgreSQL..."
        # Сохраняем новый хеш
        echo "$current_hash" > "$hash_file"
        rebuild_postgres=true
    else
        echo "Файл $file_to_check не изменился, образ PostgreSQL не пересобирается."
        rebuild_postgres=false
    fi
else
    echo "Сохраняем хеш файла $file_to_check..."
    echo "$current_hash" > "$hash_file"
    rebuild_postgres=true
fi

# Остановка и удаление контейнеров PostgreSQL, если требуется пересборка
if [ "$rebuild_postgres" = true ]; then
    echo "Пересобираем контейнер PostgreSQL..."
    docker-compose up -d --build postgres
else
    echo "Контейнер PostgreSQL не пересобирается."
fi

# Проверка наличия контейнера redis
if [ "$(docker ps -a -q -f name=gnss-radar_redis_1)" ]; then
    echo "Контейнер redis уже существует, не пересобираем."
else
    echo "Контейнер redis не найден, пересобираем..."
    docker-compose up -d --build redis
fi

# Всегда пересобираем и запускаем контейнер app
docker stop gnss-radar_app_1

echo "Перезапускаем контейнер app..."
docker-compose up -d --build app

# Запускаем все контейнеры, чтобы убедиться, что они работают
docker-compose up -d --build nginx