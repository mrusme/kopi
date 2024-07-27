CREATE TABLE IF NOT EXISTS coffee (
  id             INTEGER          PRIMARY KEY AUTOINCREMENT,
	roaster        TEXT             NOT NULL,
	name           TEXT             NOT NULL,
	origin         TEXT             NOT NULL,
	altitude_lower INTEGER,
	altitude_upper INTEGER,
	level          TEXT             NOT NULL,
	flavors        TEXT             NOT NULL,
	info           TEXT,
	roasting_date  DATE             NOT NULL,
  timestamp      TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

