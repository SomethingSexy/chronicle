CREATE EXTENSION "ltree";

-- Represents an overall game
CREATE TABLE game (
  id   BIGSERIAL PRIMARY KEY,
  game_id uuid UNIQUE NOT NULL,
  name text      NOT NULL,
  type text      NOT NULL
);

-- Represents the game world in general
-- This may or may not be necessary
CREATE TABLE world (
  id BIGSERIAL PRIMARY KEY,
  world_id uuid UNIQUE NOT NULL,
  game_id BIGSERIAL NOT NULL REFERENCES game(id),
  name text NOT NULL
);

create index world_game_id on world(game_id);

CREATE TABLE location (
  id BIGSERIAL PRIMARY KEY,
  location_id uuid UNIQUE NOT NULL,
  world_id BIGSERIAL NOT NULL REFERENCES world(id),
  type text NOT NULL,
  name text NOT NULL,
  path ltree
);

create index location_world_id on location(world_id);
create index location_path_idx on location using gist (path);