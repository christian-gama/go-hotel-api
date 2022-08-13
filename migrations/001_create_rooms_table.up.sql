BEGIN;

CREATE TABLE rooms (
  id serial PRIMARY KEY,
  uuid uuid UNIQUE NOT NULL,
  name varchar(255) UNIQUE NOT NULL,
  description text NOT NULL,
  price decimal(10,2) NOT NULL,
  bed_count integer NOT NULL
);

COMMIT;