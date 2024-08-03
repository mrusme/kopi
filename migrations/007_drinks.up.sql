CREATE TABLE IF NOT EXISTS drinks (
	id                           STRING   PRIMARY KEY,
	name                         STRING   NOT NULL,
	description                  STRING   NOT NULL,

	coffee_g                     INTEGER  NOT NULL,
	brew_ml                      INTEGER  NOT NULL,
	water_ml                     INTEGER  NOT NULL,
	milk_ml                      INTEGER  NOT NULL,
	sugar_g                      INTEGER  NOT NULL,

	is_hot                       BOOLEAN  NOT NULL,
	is_always_vegan              BOOLEAN  NOT NULL,
	can_be_vegan                 BOOLEAN  NOT NULL,

  compatible_equipment         STRING   NOT NULL
);

INSERT INTO drinks (
	id,
	name,
	description,

	coffee_g,
	brew_ml,
	water_ml,
	milk_ml,
	sugar_g,

	is_hot,
	is_always_vegan,
	can_be_vegan,

  compatible_equipment
) VALUES (
  "espresso",
  "Espresso",
  "A regular Espresso",

  9,
  37,
  37,
  0,
  0,

  TRUE,
  TRUE,
  TRUE,

  "espresso_maker grinder"
);
