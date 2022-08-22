BEGIN;

CREATE TABLE "guest" (
    "id" serial NOT NULL,
    "uuid" uuid NOT NULL UNIQUE,
    "user_id" integer NOT NULL UNIQUE REFERENCES "user" ("id") ON DELETE CASCADE,
    "credits" decimal(10,2) NOT NULL DEFAULT 0.00
);

COMMIT;