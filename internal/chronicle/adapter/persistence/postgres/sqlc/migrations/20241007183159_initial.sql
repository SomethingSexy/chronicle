CREATE EXTENSION "ltree";

-- Create "game" table
CREATE TABLE "public"."game" (
  "id" bigserial NOT NULL,
  "game_id" uuid NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "game_game_id_key" UNIQUE ("game_id")
);
-- Create "world" table
CREATE TABLE "public"."world" (
  "id" bigserial NOT NULL,
  "world_id" uuid NOT NULL,
  "game_id" bigserial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "world_world_id_key" UNIQUE ("world_id"),
  CONSTRAINT "world_game_id_fkey" FOREIGN KEY ("game_id") REFERENCES "public"."game" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "world_game_id" to table: "world"
CREATE INDEX "world_game_id" ON "public"."world" ("game_id");
-- Create "location" table
CREATE TABLE "public"."location" (
  "id" bigserial NOT NULL,
  "location_id" uuid NOT NULL,
  "game_id" bigserial NOT NULL,
  "world_id" bigserial NOT NULL,
  "type" text NOT NULL,
  "name" text NOT NULL,
  "path" public.ltree NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "location_location_id_key" UNIQUE ("location_id"),
  CONSTRAINT "location_game_id_fkey" FOREIGN KEY ("game_id") REFERENCES "public"."game" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "location_world_id_fkey" FOREIGN KEY ("world_id") REFERENCES "public"."world" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "location_game_id" to table: "location"
CREATE INDEX "location_game_id" ON "public"."location" ("game_id");
-- Create index "location_path_idx" to table: "location"
CREATE INDEX "location_path_idx" ON "public"."location" USING gist ("path");
-- Create index "location_world_id" to table: "location"
CREATE INDEX "location_world_id" ON "public"."location" ("world_id");
-- Create "character" table
CREATE TABLE "public"."character" (
  "id" bigserial NOT NULL,
  "character_id" uuid NOT NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "character_character_id_key" UNIQUE ("character_id")
);
-- Create "world_character" table
CREATE TABLE "public"."world_character" (
  "id" bigserial NOT NULL,
  "world_character_id" uuid NOT NULL,
  "character_id" bigserial NOT NULL,
  "world_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "world_character_world_character_id_key" UNIQUE ("world_character_id"),
  CONSTRAINT "world_character_character_id_fkey" FOREIGN KEY ("character_id") REFERENCES "public"."character" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "world_character_world_id_fkey" FOREIGN KEY ("world_id") REFERENCES "public"."world" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "world_character_character_id_world_id_idx" to table: "world_character"
CREATE UNIQUE INDEX "world_character_character_id_world_id_idx" ON "public"."world_character" ("character_id", "world_id");
