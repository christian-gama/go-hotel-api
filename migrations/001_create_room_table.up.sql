BEGIN;

CREATE TABLE "room" (
  "id" serial PRIMARY KEY,
  "uuid" uuid NOT NULL UNIQUE,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" text NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "bed_count" integer NOT NULL
);

COMMIT;