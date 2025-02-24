CREATE TABLE IF NOT EXISTS equipment_logs (
  id                           INTEGER  PRIMARY KEY AUTOINCREMENT,
  equipment_id                 INTEGER  NOT NULL,

  key                          STRING   NOT NULL,
	value                        STRING   NOT NULL,

  timestamp                    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(equipment_id) REFERENCES equipment(id)
);

