CREATE TABLE IF NOT EXISTS timers(
	id UUID,
	step_time INTEGER,
	counter INTEGER,
	created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()::timestamp,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()::timestamp,
	PRIMARY KEY (id)
);