BEGIN;

CREATE TABLE "restriction" (
  "id" serial PRIMARY KEY,
  "uuid" uuid NOT NULL UNIQUE,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" text NOT NULL
);

COMMIT;