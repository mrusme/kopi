CREATE TABLE IF NOT EXISTS methods (
	id                           STRING   PRIMARY KEY,
	name                         STRING   NOT NULL,
	description                  STRING   NOT NULL,

	caffeine_mg_extraction_yield_per_g INTEGER NOT NULL,
	caffeine_loss_factor         FLOAT    NOT NULL,

	is_hot                       BOOLEAN  NOT NULL
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "espresso_maker",
  "Espresso Maker",
  "A generic Espresso maker",

  10,
  0.05,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "drip_coffee_maker",
  "Drop Coffee Maker",
  "A generic Drip Coffee maker",

  8,
  0.10,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "pour_over",
  "Pour Over",
  "Pour Over",

  8,
  0.10,

  TRUE
);


INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "french_press",
  "French Press",
  "A French Press coffee maker",

  8,
  0.05,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "aeropress",
  "AeroPress",
  "An AeroPress coffee maker",

  9,
  0.10,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "moka_pot",
  "Moka Pot",
  "A Moka pot",

  10,
  0.08,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "cezve",
  "Cezve",
  "A Turkish Coffee Cezve",

  11,
  0.05,

  TRUE
);

INSERT INTO methods (
	id,
	name,
	description,

	caffeine_mg_extraction_yield_per_g,
  caffeine_loss_factor,

	is_hot
) VALUES (
  "cold_brew",
  "Cold Brew",
  "A Cold Brew",

  7,
  0.15,

  TRUE
);


