-- DROP SCHEMA bogoso;

CREATE SCHEMA bogoso AUTHORIZATION postgres;

-- bogoso.elevy_request definition

-- Drop table

-- DROP TABLE bogoso.elevy_request;

CREATE TABLE bogoso.cv_files (
	id int NOT NULL,
	applicant_name varchar not null,
	email varchar NOT NULL,
	phone varchar NOT NULL,
	file_name varchar NOT NULL,
	created timestamptz NULL DEFAULT now()
);

CREATE TABLE bogoso.temp_token (
	email varchar NOT NULL,
	token varchar NOT NULL,
	created timestamptz NULL DEFAULT now()
);

