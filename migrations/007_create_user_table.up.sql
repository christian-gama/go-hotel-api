BEGIN;

CREATE TABLE "user" (
    "id" serial PRIMARY KEY,
    "uuid" uuid NOT NULL UNIQUE,
    "email" varchar(255) NOT NULL UNIQUE REFERENCES "person" ("email") ON DELETE CASCADE,
    "password" varchar(255) NOT NULL,
    "permission_level" integer NOT NULL REFERENCES "permission" ("level") ON DELETE RESTRICT
);

COMMIT;