DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type user_role as enum ('ADMIN', 'SUPERVISOR', 'USER');
create type user_status as enum ('APPROVED', 'DECLINED', 'PENDING');

DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login TEXT NOT NULL UNIQUE DEFAULT '',
    password bytea NOT NULL DEFAULT '',
    role user_role NOT NULL DEFAULT 'USER',
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_at timestamptz NOT NULL DEFAULT now(),
    email TEXT NOT NULL UNIQUE DEFAULT '',
    organization_name TEXT NOT NULL DEFAULT '',
    first_name TEXT NOT NULL DEFAULT '',
    second_name TEXT NOT NULL DEFAULT '',
    status user_status NOT NULL DEFAULT 'PENDING'
);

DROP TABLE IF EXISTS role_api CASCADE;
CREATE TABLE IF NOT EXISTS role_api (
    role user_role NOT NULL DEFAULT 'ADMIN',
    api TEXT NOT NULL,
    description TEXT DEFAULT ''
);