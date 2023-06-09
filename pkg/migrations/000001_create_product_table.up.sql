CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT Exists products (
                          id uuid DEFAULT uuid_generate_v4 (),
                          name VARCHAR NOT NULL,
                          category VARCHAR NOT NULL,
                          suk VARCHAR NOT NULL,
                          PRIMARY KEY (id)
);