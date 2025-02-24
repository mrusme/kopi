CREATE TABLE IF NOT EXISTS coffees (
  id               INTEGER          PRIMARY KEY AUTOINCREMENT,
	roaster          TEXT             NOT NULL,
	name             TEXT             NOT NULL,
	origin           TEXT             NOT NULL,
	altitude_lower_m INTEGER          NOT NULL DEFAULT 0,
	altitude_upper_m INTEGER          NOT NULL DEFAULT 0,
	level            TEXT             NOT NULL,
	flavors          TEXT             NOT NULL DEFAULT "",
	info             TEXT             NOT NULL DEFAULT "",
	decaf            BOOLEAN          NOT NULL,
  timestamp        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

