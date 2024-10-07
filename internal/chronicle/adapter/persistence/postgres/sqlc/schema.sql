CREATE EXTENSION "ltree";

-- Represents an overall game
CREATE TABLE game (
  id   BIGSERIAL PRIMARY KEY,
  game_id uuid UNIQUE NOT NULL,
  name text      NOT NULL,
  type text      NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Represents the game world in general
-- This may or may not be necessary
CREATE TABLE world (
  id BIGSERIAL PRIMARY KEY,
  world_id uuid UNIQUE NOT NULL,
  game_id BIGSERIAL NOT NULL REFERENCES game(id),
  name text NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create index world_game_id on world(game_id);

CREATE TABLE location (
  id BIGSERIAL PRIMARY KEY,
  location_id uuid UNIQUE NOT NULL,
  -- Adding game_id here for now, incase we want locations independent of world?
  game_id BIGSERIAL NOT NULL REFERENCES game(id),
  world_id BIGSERIAL NOT NULL REFERENCES world(id),
  type text NOT NULL,
  name text NOT NULL,
  path ltree,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create index location_game_id on location(game_id);
create index location_world_id on location(world_id);
create index location_path_idx on location using gist(path);

CREATE TABLE character (
  id BIGSERIAL PRIMARY KEY,
  character_id uuid UNIQUE NOT NULL,
  name text NOT NULL,
  description text,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE world_character (
  id BIGSERIAL PRIMARY KEY,
  world_character_id uuid UNIQUE NOT NULL,
  character_id BIGSERIAL NOT NULL REFERENCES character(id),
  world_id BIGSERIAL NOT NULL REFERENCES world(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create unique index world_character_character_id_world_id_idx on world_character(character_id,world_id);