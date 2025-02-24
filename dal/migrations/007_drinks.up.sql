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

  compatible_methods           STRING   NOT NULL,
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "espresso",
  "Espresso",
  "A single-shot Espresso",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  0, -- TODO
  0,

  TRUE,
  TRUE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "cappuccino",
  "Cappuccino",
  "A single-shot Cappuccino",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  100, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "latte",
  "Latte",
  "A single-shot Latte",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  180, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "flat_white",
  "Flat White",
  "A single-shot Flat White",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  180, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "macchiato",
  "Macchiato",
  "A single-shot Macchiato",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  20, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "latte_macchiato",
  "Latte Macchiato",
  "A single-shot Latte Macchiato",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  20, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "cortado",
  "Cortado",
  "A single-shot Cortado",

  9, -- TODO
  25, -- TODO
  25, -- TODO
  60, -- TODO
  0,

  TRUE,
  FALSE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder frother"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "drip_coffee",
  "Drip Coffee",
  "A regular Drip Coffee",

  10, -- TODO
  180, -- TODO
  180, -- TODO
  0, -- TODO
  0,

  TRUE,
  TRUE,
  TRUE,

  "drip_coffee_maker pour_over",
  "coffee_maker grinder"
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

  compatible_methods,
  compatible_equipment
) VALUES (
  "americano",
  "Americano",
  "A regular Americano",

  10, -- TODO
  180, -- TODO
  180, -- TODO
  0, -- TODO
  0,

  TRUE,
  TRUE,
  TRUE,

  "espresso_maker",
  "espresso_maker grinder"
);


