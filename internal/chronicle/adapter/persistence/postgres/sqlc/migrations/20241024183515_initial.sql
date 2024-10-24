CREATE EXTENSION "ltree";

-- Create "world" table
CREATE TABLE "public"."world" (
  "id" bigserial NOT NULL,
  "world_id" uuid NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "world_world_id_key" UNIQUE ("world_id")
);
-- Create "game" table
CREATE TABLE "public"."game" (
  "id" bigserial NOT NULL,
  "game_id" uuid NOT NULL,
  "world_id" bigserial NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "game_game_id_key" UNIQUE ("game_id"),
  CONSTRAINT "game_world_id_fkey" FOREIGN KEY ("world_id") REFERENCES "public"."world" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "game_character_game_id_idx" to table: "game"
CREATE INDEX "game_character_game_id_idx" ON "public"."game" ("id");
-- Create index "game_world_id_idx" to table: "game"
CREATE INDEX "game_world_id_idx" ON "public"."game" ("game_id");
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
-- Create index "game_character_character_id_idx" to table: "character"
CREATE INDEX "game_character_character_id_idx" ON "public"."character" ("id");
-- Create "game_character" table
CREATE TABLE "public"."game_character" (
  "id" bigserial NOT NULL,
  "game_character_id" uuid NOT NULL,
  "game_id" bigserial NOT NULL,
  "character_id" bigserial NOT NULL,
  "character_type" text NOT NULL,
  "character" jsonb NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "game_character_game_character_id_key" UNIQUE ("game_character_id"),
  CONSTRAINT "game_character_character_id_fkey" FOREIGN KEY ("character_id") REFERENCES "public"."character" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "game_character_game_id_fkey" FOREIGN KEY ("game_id") REFERENCES "public"."game" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "game_character_character_id_game_id_idx" to table: "game_character"
CREATE UNIQUE INDEX "game_character_character_id_game_id_idx" ON "public"."game_character" ("character_id", "game_id");
-- Create index "game_character_game_character_id_idx" to table: "game_character"
CREATE INDEX "game_character_game_character_id_idx" ON "public"."game_character" ("game_character_id");
-- Create "location" table
CREATE TABLE "public"."location" (
  "id" bigserial NOT NULL,
  "location_id" uuid NOT NULL,
  "world_id" bigserial NOT NULL,
  "type" text NOT NULL,
  "name" text NOT NULL,
  "path" public.ltree NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "location_location_id_key" UNIQUE ("location_id"),
  CONSTRAINT "location_world_id_fkey" FOREIGN KEY ("world_id") REFERENCES "public"."world" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "location_path_idx" to table: "location"
CREATE INDEX "location_path_idx" ON "public"."location" USING gist ("path");
-- Create index "location_world_id_idx" to table: "location"
CREATE INDEX "location_world_id_idx" ON "public"."location" ("world_id");
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
