CREATE TABLE IF NOT EXISTS cups (
  id                     INTEGER   PRIMARY KEY AUTOINCREMENT,
  coffee_id              INTEGER   NOT NULL,
  drink                  STRING    NOT NULL,
  vegan                  BOOLEAN   NOT NULL,
	coffee_g               INTEGER   NOT NULL,
	brew_ml                INTEGER   NOT NULL,
	water_ml               INTEGER   NOT NULL,
	milk_ml                INTEGER   NOT NULL,
	sugar_g                INTEGER   NOT NULL,
  rating                 INTEGER   NOT NULL DEFAULT 0,
  timestamp              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(coffee_id) REFERENCES coffees(id),
  FOREIGN KEY(drink) REFERENCES drinks(id)
);

