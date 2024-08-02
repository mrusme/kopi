CREATE TABLE IF NOT EXISTS methods (
	id                           STRING   PRIMARY KEY,
	name                         STRING   NOT NULL,
	description                  STRING   NOT NULL,

	caffeine_multiplier_per_ml   FLOAT    NOT NULL,

	is_hot                       BOOLEAN  NOT NULL
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "espresso_maker",
  "Espresso Maker",
  "A generic Espresso maker",

  0.5,

  TRUE
);

