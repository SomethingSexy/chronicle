CREATE EXTENSION "ltree";

-- Represents a generic world that can be linked to many games
CREATE TABLE world (
  id BIGSERIAL PRIMARY KEY,
  world_id uuid UNIQUE NOT NULL,
  name text NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE location (
  id BIGSERIAL PRIMARY KEY,
  location_id uuid UNIQUE NOT NULL,
  world_id BIGSERIAL NOT NULL REFERENCES world(id),
  type text NOT NULL,
  name text NOT NULL,
  path ltree,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create index location_world_id_idx on location(world_id);
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

-- Represents an overall game
CREATE TABLE game (
  id   BIGSERIAL PRIMARY KEY,
  game_id uuid UNIQUE NOT NULL,
  world_id BIGSERIAL NOT NULL REFERENCES world(id),
  name text      NOT NULL,
  type text      NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create index game_world_id_idx on game(game_id);

CREATE TABLE game_character (
  id   BIGSERIAL PRIMARY KEY,
  game_character_id uuid UNIQUE NOT NULL,
  game_id BIGSERIAL NOT NULL REFERENCES game(id),
  character_id BIGSERIAL NOT NULL REFERENCES character(id),
  character_type text      NOT NULL,
  character JSONB,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create index game_character_game_character_id_idx on game_character(game_character_id);
create index game_character_game_id_idx on game(id);
create index game_character_character_id_idx on character(id);
create unique index game_character_character_id_game_id_idx on game_character(character_id,game_id);