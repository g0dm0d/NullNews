CREATE TABLE users
(
 ID  SERIAL PRIMARY KEY,
 username   varchar(20)   UNIQUE NOT NULL,
 firstname  varchar(20)   NOT NULL,
 lastname   varchar(20)   NOT NULL,
 email      varchar(50)   UNIQUE NOT NULL,
 password   varchar(256)  NOT NULL
);

CREATE TABLE news
(
 ID  SERIAL PRIMARY KEY,
 title    varchar(50) NOT NULL,
 "user"   int         NOT NULL,
 text     text        NOT NULL,
 CONSTRAINT FK_1 FOREIGN KEY ( "user" ) REFERENCES "users" ( id )
);

CREATE INDEX FK_2 ON news
(
 "user"
);
