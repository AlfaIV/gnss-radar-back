DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   login TEXT NOT NULL UNIQUE DEFAULT '',
   password bytea NOT NULL DEFAULT '',
   role TEXT NOT NULL DEFAULT 'user'
);

INSERT INTO profile(login, password, role) VALUES ('admin', '\xc7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec', 'admin');