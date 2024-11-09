DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   login TEXT NOT NULL UNIQUE DEFAULT '',
   password bytea NOT NULL DEFAULT '',
   role TEXT NOT NULL DEFAULT 'user'
);

DROP TABLE IF EXISTS gnss_coords CASCADE;
CREATE TABLE IF NOT EXISTS gnss_coords (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satellite_id TEXT NOT NULL,
    satellite_name TEXT NOT NULL,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    z DOUBLE PRECISION NOT NULL
);

DROP TABLE IF EXISTS devices CASCADE;
CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    token TEXT UNIQUE NOT NULL,
    description TEXT,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    z DOUBLE PRECISION NOT NULL
);

INSERT INTO profile(login, password, role) VALUES ('admin', '\xc7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec', 'admin');

INSERT INTO gnss_coords (satellite_id, satellite_name, x, y, z) VALUES
                                                                    ('PC06', 'PC06', -16806.320344, 29291.120310, -25355.710938),
                                                                    ('PC07', 'PC07', -6959.418476, 39332.954409, -13000.851001),
                                                                    ('PC08', 'PC08', -1908.204600, 21553.224987, 36203.881809),
                                                                    ('PC09', 'PC09', -11202.586298, 28046.331947, -29182.143554),
                                                                    ('PC10', 'PC10', -917.431406, 41238.966109, -6711.991412),
                                                                    ('PC11', 'PC11', -16138.177056, -3913.891460, -22348.411693),
                                                                    ('PC12', 'PC12', -997.099233, -19759.345910, -19638.934483),
                                                                    ('PC13', 'PC13', 5858.392549, 25505.986419, 33308.911170),
                                                                    ('PC14', 'PC14', -17706.605729, -14691.268566, 15829.680477),
                                                                    ('PC16', 'PC16', -22387.055407, 28560.640995, -21454.026667);