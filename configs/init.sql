create table users (
  email text not null unique,
  primary key(email),
  name text
);

create table events (
  id serial not null unique,
  primary key(id),
  title text,
  description text,
  cost integer,
	location text,
  organizerName text,
	organizerEmail text,
  foreign key(organizerEmail) REFERENCES users(email),
  attendees text[],
	startTime text,
	endTime text
);