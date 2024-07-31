-- Create "game" table
CREATE TABLE "public"."game" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL,
  PRIMARY KEY ("id")
);
