BEGIN;

CREATE TABLE "room_restriction" (
  "id" serial PRIMARY KEY,
  "uuid" uuid NOT NULL UNIQUE,
  "room_id" integer NOT NULL UNIQUE REFERENCES "room" ("id") ON DELETE CASCADE,
  "restriction_id" integer NOT NULL REFERENCES "restriction" ("id") ON DELETE CASCADE
);

COMMIT;