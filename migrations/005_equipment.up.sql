CREATE TABLE IF NOT EXISTS equipment (
  id                           INTEGER  PRIMARY KEY AUTOINCREMENT,
	type                         STRING   NOT NULL,

  name                         STRING   NOT NULL,
	description                  STRING   NOT NULL,

	purchase_date                DATE     NOT NULL,
	decommission_date            DATE     DEFAULT NULL,

	price_usd_ct                 INTEGER  NOT NULL DEFAULT 0,
	price_sats                   INTEGER  NOT NULL DEFAULT 0,

  timestamp                    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
