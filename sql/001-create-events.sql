create table if not exists events (
	id UUID primary key,
	createdat timestamp not null,
	updatedat timestamp,
	deletedat timestamp,
	occursat timestamp,
	alertevery bigint,
	subject text not null,
	body text,
	duration bigint,
	location text,
	userid UUID not null
);
