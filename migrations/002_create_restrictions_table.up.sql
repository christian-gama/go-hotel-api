BEGIN;

CREATE TABLE restrictions (
  id serial PRIMARY KEY,
  uuid uuid NOT NULL,
  name varchar(255) NOT NULL,
  description text NOT NULL
);

COMMIT;