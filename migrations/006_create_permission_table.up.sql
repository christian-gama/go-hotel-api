BEGIN;

CREATE TABLE "permission" (
    "id" serial PRIMARY KEY,
    "uuid" uuid NOT NULL UNIQUE,
    "level" integer NOT NULL UNIQUE,
    "description" text NOT NULL
);

COMMIT;