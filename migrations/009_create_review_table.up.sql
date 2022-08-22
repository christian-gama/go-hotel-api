BEGIN;

CREATE TABLE "review" (
    "id" serial NOT NULL,
    "uuid" uuid NOT NULL UNIQUE,
    "guest_id" integer NOT NULL,
    "room_id" integer NOT NULL,
    "description" text NOT NULL,
    "rating" integer NOT NULL,
    UNIQUE ("guest_id", "room_id")
);

COMMIT;