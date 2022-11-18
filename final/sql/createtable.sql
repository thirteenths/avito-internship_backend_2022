DROP TABLE balance CASCADE ;
DROP TABLE payroll CASCADE ;
DROP TABLE reserve CASCADE ;

CREATE TABLE balance (
	id SERIAL PRIMARY KEY,
	id_user integer NOT NULL,
	summa REAL NOT NULL,
	date_upgrate TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE payroll (
	id SERIAL PRIMARY KEY,
	id_balance integer REFERENCES balance (id),
	summa REAL NOT NULL,
	date_payroll TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE  TABLE reserve(
	id SERIAL PRIMARY KEY ,
	id_balance integer REFERENCES balance(id),
	summa REAL NOT NULL,
	data_reserve TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
	id_order integer NOT NULL,
	id_service integer NOT NULL 
);
