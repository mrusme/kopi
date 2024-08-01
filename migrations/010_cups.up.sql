CREATE TABLE IF NOT EXISTS cups (
  id                     INTEGER   PRIMARY KEY AUTOINCREMENT,
  coffee_id              INTEGER   NOT NULL,
  drink                  STRING    NOT NULL,
  vegan                  BOOLEAN   NOT NULL,
	override_coffee_g      INTEGER   DEFAULT NULL,
	override_brew_ml       INTEGER   DEFAULT NULL,
	override_water_ml      INTEGER   DEFAULT NULL,
	override_milk_ml       INTEGER   DEFAULT NULL,
	override_sugar_g       INTEGER   DEFAULT NULL,
  rating                 INTEGER   NOT NULL DEFAULT 0,
  timestamp              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(coffee_id) REFERENCES coffees(id),
  FOREIGN KEY(drink) REFERENCES drinks(id)
);

