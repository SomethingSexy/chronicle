-- Represents an overall game
CREATE TABLE game (
  id   BIGSERIAL PRIMARY KEY,
  game_id uuid NOT NULL,
  name text      NOT NULL,
  type text      NOT NULL
);

-- Represents the game world in general
-- This may or may not be necessary
CREATE TABLE world (
  id BIGSERIAL PRIMARY KEY,
  world_id uuid NOT NULL,
  game_id BIGSERIAL NOT NULL REFERENCES game(id)
);

CREATE TABLE location (
  ud BIGSERIAL PRIMARY KEY,
  location_id uuid NOT NULL,
  world_ID BIGSERIAL NOT NULL REFERENCES world(id),
  type text NOT NULL,
  name text NOT NULL,
  path ltree
);

create index location_path_idx on location using gist (path);