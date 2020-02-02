
create table msg(
	id    		bigserial	    primary key,
	ts	        bigint		    not null,
	from_addr	varchar(100)    not null,
	to_addr     varchar(100),
	topic	    varchar(100)    not null,
	type	    varchar(100)    not null,
    sha256      bytea           not null,
    psha256     bytea,
	content	    text            not null
);

-- create table msg_attachment(
-- 	msg_id		bitint		references msg.id,
-- 	filename	text,
-- 	filepath	text,
-- 	filesize	
-- );