BEGIN;

CREATE TABLE "person" (
    "id" serial PRIMARY KEY,
    "uuid" uuid NOT NULL UNIQUE,
    "first_name" varchar(150) NOT NULL,
    "last_name" varchar(150) NOT NULL,
    "email" varchar(255) NOT NULL UNIQUE,
    "phone" varchar(30) NOT NULL UNIQUE,
    "ssn" varchar(30) NOT NULL UNIQUE,
    "address_id" integer NOT NULL REFERENCES "address" ("id") ON DELETE RESTRICT
);

COMMIT;