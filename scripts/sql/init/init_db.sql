DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login TEXT NOT NULL UNIQUE DEFAULT '',
    password bytea NOT NULL DEFAULT '',
    role TEXT NOT NULL DEFAULT 'USER',
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_at timestamptz NOT NULL DEFAULT now(),
    email TEXT NOT NULL UNIQUE DEFAULT '',
    organization_name TEXT NOT NULL DEFAULT '',
    first_name TEXT NOT NULL DEFAULT '',
    second_name TEXT NOT NULL DEFAULT ''
);

DROP TABLE IF EXISTS satellites CASCADE;
CREATE TABLE IF NOT EXISTS satellites (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_satellite_id TEXT UNIQUE NOT NULL,
    satellite_name TEXT UNIQUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
    );

DROP TABLE IF EXISTS gnss_coords CASCADE;
CREATE TABLE IF NOT EXISTS gnss_coords (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satellite_id TEXT NOT NULL,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    z DOUBLE PRECISION NOT NULL,
    coordinate_measurement_time timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
    );

DROP TABLE IF EXISTS devices CASCADE;
CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    token TEXT UNIQUE NOT NULL,
    description TEXT,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    z DOUBLE PRECISION NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
    );

DROP TABLE IF EXISTS tasks CASCADE;
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    satellite_id TEXT NOT NULL,
    device_id TEXT NOT NULL,
    signal_type TEXT NOT NULL,
    grouping_type TEXT NOT NULL,
    start_at timestamptz NOT NULL,
    end_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
    );

DROP TABLE IF EXISTS hardware_measurements CASCADE;
CREATE TABLE IF NOT EXISTS hardware_measurements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token TEXT NOT NULL,
    start_at timestamptz NOT NULL,
    end_at timestamptz NOT NULL,
    group_type TEXT NOT NULL, 
    signal TEXT NOT NULL,
    satellite_name TEXT NOT NULL,
    measurement_power_id UUID DEFAULT uuid_generate_v4(),
    measurement_spectrum_id UUID DEFAULT uuid_generate_v4(),
    created_at timestamptz NOT NULL DEFAULT now()
);

DROP TABLE IF EXISTS measurements_power CASCADE;
CREATE TABLE IF NOT EXISTS measurements_power (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    power FLOAT[] NOT NULL,
    started_at timestamptz NOT NULL,
    time_step timestamptz NOT NULL
);

DROP TABLE IF EXISTS measurements_spectrum CASCADE;
CREATE TABLE IF NOT EXISTS measurements_spectrum (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    spectrum FLOAT[] NOT NULL,
    start_freq FLOAT NOT NULL,
    freq_step FLOAT NOT NULL,
    started_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX tas_device_id_start_end_satellite_id_signal_type ON tasks (device_id, satellite_id, signal_type, start_at, end_at);
-- CREATE UNIQUE INDEX tas_satellite_idx ON satellites (satellite_name, external_satellite_id);

INSERT INTO profile(login, password, role, email, organization_name, first_name, second_name) VALUES ('admin', '\xc7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec', 'ADMIN', 'admin@mail.ru', 'gnss-company', 'admin', 'admin');

INSERT INTO satellites (external_satellite_id, satellite_name) VALUES
                                                                   ('PC06', 'PC06'),
                                                                   ('PC07', 'PC07'),
                                                                   ('PC08', 'PC08'),
                                                                   ('PC09', 'PC09'),
                                                                   ('PC10', 'PC10'),
                                                                   ('PC11', 'PC11');

INSERT INTO gnss_coords (satellite_id, x, y, z, coordinate_measurement_time) VALUES
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC06'), -16806.320344, 29291.120310, -25355.710938, now()),
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC07'), -6959.418476, 39332.954409, -13000.851001, now()),
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC08'), -1908.204600, 21553.224987, 36203.881809, now()),
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC09'), -11202.586298, 28046.331947, -29182.143554, now()),
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC10')   -917.431406, 41238.966109, -6711.991412, now()),
                                                    ((SELECT id FROM satellites WHERE external_satellite_id = 'PC11'), -16138.177056, -3913.891460, -22348.411693, now());

INSERT INTO devices (name, token, description, x, y, z) VALUES
                                                            ('device1', uuid_generate_v4(), 'desc1', 10.0, 20.0, 30.0),
                                                            ('device2', uuid_generate_v4(), 'desc2', 15.0, 25.0, 35.0),
                                                            ('device3', uuid_generate_v4(), 'desc3', 20.0, 30.0, 40.0);

INSERT INTO tasks (satellite_id, device_id, title, description, signal_type, grouping_type, start_at, end_at) VALUES
                                                   ((SELECT id FROM satellites WHERE external_satellite_id = 'PC06'), (SELECT id FROM devices WHERE name = 'device1'), 'Задание 1', 'Описание 1', 'SIGNAL_TYPE_L1', 'GROUPING_TYPE_GPS', now(), now() + interval '2 days'),
                                                   ((SELECT id FROM satellites WHERE external_satellite_id = 'PC07'), (SELECT id FROM devices WHERE name = 'device2'), 'Задание 2', 'Описание 2', 'SIGNAL_TYPE_L2', 'GROUPING_TYPE_GLONASS', now(), now() + interval '3 days'),
                                                   ((SELECT id FROM satellites WHERE external_satellite_id = 'PC08'), (SELECT id FROM devices WHERE name = 'device3'), 'Задание 3', '', 'SIGNAL_TYPE_L3', 'GROUPING_TYPE_GLONASS', now(), now() + interval '4 days');