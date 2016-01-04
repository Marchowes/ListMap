use listmap;
-- since way too many ISPs dont support INNODB, we dont use foreign keys.

create table `users` (
	user_id		bigint unsigned not NULL auto_increment,
	email		varchar(255) not NULL,
	password	varchar(255) not NULL,
	user_level	int not NULL,
	creation_date	datetime not NULL,
	last_login	datetime,
	user_name	varchar(70) not NULL,
	primary key	(user_id)
) engine=MyISAM
COMMENT = 'USER TABLE';

create table `pins` (
	pin_id		bigint unsigned not NULL auto_increment,
	name		varchar(255) not NULL,
	latitude	varchar(15) not NULL,
	longitude	varchar(15) not NULL,
	date_added	datetime not NULL,
	creator		bigint unsigned not NULL,	
	pin_type	varchar(50) not NULL,
	primary key	(pin_id)
) engine=MyISAM
COMMENT = 'Base PIN Super Class, further data is pulled in from data pulled in from pin_type';

create table `glue_map` (
	glue_id		bigint unsigned not NULL auto_increment,
	data_type	varchar(50) not NULL,
	table_name	varchar(50) not NULL,
	primary key	(glue_id)
) engine=MyISAM
COMMENT = 'Glue record. data_type is a string used in a source table. table name refers to the table name of that additional data set';

create table `peaks` (
	peak_id		bigint unsigned not NULL auto_increment,	
        pin_id		bigint unsigned not NULL,
	elevation_ft    varchar(15) not NULL,
        elevation_m     varchar(15),
	state		varchar(30),
	prominence	varchar(15),
	col		varchar(15),
	primary key 	(peak_id)
) engine=MyISAM
COMMENT = 'Incomplete peak record. its reasonable to assume this will get much much larger';

create table `list` (
	list_id		bigint unsigned not NULL auto_increment,
	name		varchar(50) not NULL,
	creator_user	bigint unsigned not NULL,
	creation_date	datetime,
	primary key	(list_id)
) engine=MyISAM
COMMENT = 'List Data Super Class';

create table `list_linker` (
	linker_id	bigint unsigned not NULL auto_increment,
	list_id		bigint unsigned not NULL,
	user_id		bigint unsigned not NULL,
	pin_id		bigint unsigned not NULL,
	link_date	datetime,
	primary key	(linker_id)
) engine=MyISAM
COMMENT = 'Linker Class, links pin objects with list objects';

create table `list_permission` (
	list_permission_id	bigint unsigned not NULL auto_increment,
	list_id		bigint unsigned not NULL,
	user_id		bigint unsigned,
	user_level      bigint unsigned,	
	primary key 	(list_permission_id)
) engine=MyISAM
COMMENT = 'grants non creator users permission to see lists';

create table `list_entry` (
	entry_id	bigint unsigned not NULL auto_increment,
	pin_id		bigint unsigned not NULL,
	user_id		bigint unsigned not NULL,
	pin_date	varchar(30),
	list_entry_date	datetime,
	primary key	(entry_id)
) engine=MyISAM
COMMENT = 'Actual user entry for pin';

	

