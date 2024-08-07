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

  2.73,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "drip_coffee_maker",
  "Drop Coffee Maker",
  "A generic Drip Coffee maker",

  0.68,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "pour_over",
  "Pour Over",
  "Pour Over",

  0.74,

  TRUE
);


INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "french_press",
  "French Press",
  "A French Press coffee maker",

  0.89,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "aeropress",
  "AeroPress",
  "An AeroPress coffee maker",

  1.36,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "moka_pot",
  "Moka Pot",
  "A Moka pot",

  1.64,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "cezve",
  "Cezve",
  "A Turkish Coffee Cezve",

  1.5,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_multiplier_per_ml,

	is_hot
) VALUES (
  "cold_brew",
  "Cold Brew",
  "A Cold Brew",

  1.12,

  TRUE
);


