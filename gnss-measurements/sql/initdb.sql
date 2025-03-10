DROP TABLE IF EXISTS file_meta CASCADE;
CREATE TABLE IF NOT EXISTS profile (
    filename TEXT NOT NULL DEFAULT '',
    minio_name TEXT NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now()
);
