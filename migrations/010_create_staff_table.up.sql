BEGIN;

CREATE TABLE "staff" (
    "id" serial NOT NULL,
    "uuid" uuid NOT NULL UNIQUE,
    "role" varchar(50) NOT NULL,
    "person_id" integer NOT NULL UNIQUE REFERENCES "person" ("id") ON DELETE CASCADE
);

COMMIT;