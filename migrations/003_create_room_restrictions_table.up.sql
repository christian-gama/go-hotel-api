BEGIN;

CREATE TABLE room_restrictions (
  id serial PRIMARY KEY,
  uuid uuid NOT NULL,
  room_id integer UNIQUE NOT NULL REFERENCES rooms (id) ON DELETE CASCADE,
  restriction_id integer NOT NULL REFERENCES restrictions (id) ON DELETE CASCADE
);

COMMIT;