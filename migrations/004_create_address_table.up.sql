BEGIN;

CREATE TABLE "address" (
    "id" serial PRIMARY KEY,
    "uuid" uuid NOT NULL UNIQUE,
    "street" varchar(100) NOT NULL,
    "number" varchar(10) NOT NULL,
    "zip_code" varchar(10) NOT NULL,
    "city" varchar(50) NOT NULL,
    "country" varchar(50) NOT NULL,
    "state" varchar(50) NOT NULL
);

COMMIT;