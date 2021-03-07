CREATE KEYSPACE IF NOT EXISTS shrink_url WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

CREATE TABLE IF NOT EXISTS shrink_url.user (
id int,
name text,
email text,
last_login date,
creation timestamp,
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS shrink_url.url (
                                               hash text,
                                               original_url text,
                                               expiration_date timestamp,
                                               creation_date timestamp,
                                               PRIMARY KEY (hash)
);


select * from shrink_url.url;
