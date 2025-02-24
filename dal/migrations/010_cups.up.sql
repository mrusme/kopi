CREATE TABLE IF NOT EXISTS cups (
  id                     INTEGER   PRIMARY KEY AUTOINCREMENT,
  bag_id                 INTEGER   NOT NULL,

  method                 STRING    NOT NULL,
  drink                  STRING    NOT NULL,

  equipment_ids          STRING    NOT NULL DEFAULT "",

  vegan                  BOOLEAN   NOT NULL,
	coffee_g               INTEGER   NOT NULL,
	brew_ml                INTEGER   NOT NULL,
	water_ml               INTEGER   NOT NULL,
	milk_ml                INTEGER   NOT NULL,
  milk_type              STRING    NOT NULL DEFAULT "none",
	sugar_g                INTEGER   NOT NULL,
  rating                 INTEGER   NOT NULL DEFAULT 0,
  timestamp              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(bag_id) REFERENCES bags(id),
  FOREIGN KEY(method) REFERENCES methods(id),
  FOREIGN KEY(drink) REFERENCES drinks(id)
);

