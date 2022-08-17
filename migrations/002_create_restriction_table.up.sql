BEGIN;

CREATE TABLE restriction (
  id serial PRIMARY KEY,
  uuid uuid UNIQUE NOT NULL,
  name varchar(255) UNIQUE NOT NULL,
  description text NOT NULL
);

COMMIT;