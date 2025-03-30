DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS task CASCADE;
CREATE TABLE IF NOT EXISTS task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL DEFAULT '',
    time_start timestamptz NOT NULL DEFAULT now(),
    time_end timestamptz NOT NULL DEFAULT now(),
    description TEXT NOT NULL DEFAULT '',
    creator_id UUID NOT NULL
);

DROP TABLE IF EXISTS task_satellites CASCADE;
CREATE TABLE IF NOT EXISTS task_satellites (
    task_id UUID NOT NULL REFERENCES task(id),
    satellite_name TEXT NOT NULL DEFAULT 'ALL'
);