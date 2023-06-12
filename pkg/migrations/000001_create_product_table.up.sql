CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT Exists products (
                          id uuid DEFAULT uuid_generate_v4 (),
                          name VARCHAR NOT NULL,
                          category VARCHAR NOT NULL,
                          sku VARCHAR NOT NULL,
                          PRIMARY KEY (id)
);

alter table products add column search tsvector generated always as (to_tsvector('english', products.name || ' ' || products.category || ' ' || products.sku)) stored;

create index products_search_idx on products using GIN(search);