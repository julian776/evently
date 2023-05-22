create table events (
  id serial not null unique,
  title text,
  primary key(id),
  description text,
  organizerName text,
	organizerEmail text,
	location text,
	startTime text,
	endTime text
);


insert into events(title, description, organizerName, organizerEmail, location, startTime, endTime)
values
    ('title', 'The description ...', 'organizerName', 'organizerEmail', 'location', 'startTime', 'endTime');