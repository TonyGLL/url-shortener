-- Create schema
CREATE SCHEMA url_shortener_schema;

-- Tables creation
CREATE TABLE url_shortener_schema.sites (
    id SERIAL PRIMARY KEY,
    key varchar(20) UNIQUE NOT NULL,
    long_url text NOT NULL,
    salt bigint NOT NULL,
    expiration timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE url_shortener_schema.searches (
    id SERIAL PRIMARY KEY,
    ip_address varchar(15),
    browser varchar(20),
    site_id integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);

-- Add foreign keys
ALTER TABLE url_shortener_schema.searches
  ADD CONSTRAINT fk_site_id FOREIGN KEY (site_id) REFERENCES url_shortener_schema.sites (id);