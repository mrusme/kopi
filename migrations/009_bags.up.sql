CREATE TABLE IF NOT EXISTS bags (
  id               INTEGER          PRIMARY KEY AUTOINCREMENT,
  coffee_id        INTEGER          NOT NULL,

	weight_g         INTEGER          NOT NULL,
	grind            STRING           NOT NULL,

	roast_date       DATE             NOT NULL,
	open_date        DATE             NOT NULL,
	empty_date       DATE             DEFAULT NULL,
	purchase_date    DATE             NOT NULL,

	price_usd_ct     INTEGER          NOT NULL DEFAULT 0,
	price_sats       INTEGER          NOT NULL DEFAULT 0,

  timestamp        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(coffee_id) REFERENCES coffees(id)
);


