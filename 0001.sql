CREATE TABLE users
(
 ID         SERIAL        PRIMARY KEY,
 username   VARCHAR(20)   UNIQUE NOT NULL,
 firstname  VARCHAR(20)   NOT NULL,
 lastname   VARCHAR(20)   NOT NULL,
 email      VARCHAR(50)   UNIQUE NOT NULL,
 password   VARCHAR(256)  NOT NULL,
 permission INTEGER       NULL
);

CREATE TABLE news
(
 ID       SERIAL      PRIMARY KEY,
 title    VARCHAR(50) NOT NULL,
 text     TEXT        NOT NULL,
 author   INTEGER     NOT NULL REFERENCES users(id)
);

CREATE TABLE sessions
(
  id            SERIAL        PRIMARY KEY,
  refresh_token VARCHAR(255)  NOT NULL,
  user_id       INTEGER       NOT NULL REFERENCES users(id),
  expires_time  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
