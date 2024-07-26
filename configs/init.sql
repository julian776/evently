CREATE DATABASE evently;

-- With GORM we can skip this step

-- create table users (
--   email text not null unique,
--   primary key(email),
--   name text,
--   password text not null,
--   purpouseOfUse text not null
-- );

-- create table events (
--   id serial not null unique,
--   primary key(id),
--   title text,
--   description text,
--   cost real,
-- 	location text,
--   attendees text[],
--   organizerName text,
-- 	organizerEmail text,
--   foreign key(organizerEmail) REFERENCES users(email),
-- 	startDate text,
-- 	endDate text,
--   startTime text,
-- 	endTime text
-- );