-- See https://dbdiagram.io/d/60df48cd0b1d8a6d39649fa4 for designer

-------------------------
-- USER TABLES
-------------------------
CREATE TABLE IF NOT EXISTS users (
   id SERIAL PRIMARY KEY,
   name VARCHAR (32) NOT NULL,
   email VARCHAR(64) UNIQUE NOT NULL,
   creation TIMESTAMPTZ NOT NULL,
   password TEXT NOT NULL,
   last_login TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS user_groups (
   id SERIAL PRIMARY KEY,
   name VARCHAR (16) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_group_assignments (
   id SERIAL PRIMARY KEY,
   user_id SERIAL REFERENCES users(id),
   group_id SERIAL REFERENCES user_groups(id)
);

-------------------------
-- WORKSPACE TABLES
-------------------------
CREATE TABLE IF NOT EXISTS workspaces (
   id SERIAL PRIMARY KEY,
   name VARCHAR(32),
   creation TIMESTAMPTZ NOT NULL,
   last_used TIMESTAMPTZ NOT NULL,
   floors INT NOT NULL
);

CREATE TABLE IF NOT EXISTS workspace_groups (
   id SERIAL PRIMARY KEY,
   name VARCHAR(16) NOT NULL,
);

CREATE TABLE IF NOT EXISTS user_workspace_assignments (
   id SERIAL PRIMARY KEY,
   user_id SERIAL REFERENCES users(id),
   workspace_id SERIAL REFERENCES workspaces(id),
   workspace_group_id SERIAL REFERENCES workspace_groups(id)
)

CREATE TABLE IF NOT EXISTS desks (
   id SERIAL PRIMARY KEY,
   name VARCHAR(16) NOT NULL,
   creation TIMESTAMPTZ NOT NULL,
   workspace_id SERIAL REFERENCES workspaces(id),
   floor INT NOT NULL,
   location_x INT NOT NULL,
   location_y INT NOT NULL,
   updated_by SERIAL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS desk_bookings (
   id SERIAL PRIMARY KEY,
   desk_id SERIAL REFERENCES desks(id),
   user_id SERIAL REFERENCES users(id),
   start_date TIMESTAMPTZ NOT NULL,
   end_date TIMESTAMPTZ
);
