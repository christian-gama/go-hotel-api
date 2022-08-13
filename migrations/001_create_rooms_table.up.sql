CREATE TABLE rooms (
  id serial PRIMARY KEY,
  uuid uuid NOT NULL,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  price decimal(10,2) NOT NULL,
  bed_count integer NOT NULL
);