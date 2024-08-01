CREATE TABLE IF NOT EXISTS cups (
  id         INTEGER          PRIMARY KEY AUTOINCREMENT,
  coffee_id  INTEGER          NOT NULL,
  drink      TEXT             NOT NULL,
  timestamp  TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(coffee_id) REFERENCES coffees(id)
);

