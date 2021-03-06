create database listmap;
use listmap;
-- since way too many ISPs dont support INNODB, we dont use foreign keys.

create table `user` (
	user_id       bigint unsigned not NULL auto_increment,
	email         varchar(255) not NULL,
	password      varchar(255) not NULL,
	user_level    int not NULL,
	creation_date datetime not NULL,
	last_login    datetime,
	user_name     varchar(70) not NULL,
	primary key   (user_id)
) engine=InnoDB
COMMENT = 'USER TABLE';

create table `login` (
	login_id     bigint unsigned not NULL auto_increment,
	user_id      bigint unsigned not NULL,
	ip_address   varchar(48) not NULL,
	created      integer unsigned not NULL default 0,
	modified     integer unsigned not NULL default 0,
	token        varchar(255) not NULL default '',
	primary key  (login_id),
	foreign key  (user_id) REFERENCES user(user_id)
	unique key   (token),
	key          (user_id),
	key          (ip_address)
) engine=InnoDB
COMMENT = `LOGIN TABLE`


create table `peak` (
	peak_id         bigint unsigned not NULL auto_increment,
	elevation_ft    varchar(15) not NULL,
	elevation_m     varchar(15),
	latitude        varchar(15),
	longitude       varchar(15),
	country         varchar(255),
	state           varchar(255),
	county          varchar(255),
	town            varchar(255),
	prominence_ft   varchar(15),
	prominence_m    varchar(15),
	name            varchar(255) not NULL,
	topo            varchar(255),
	col             varchar(15),
	date_added      datetime,
	creator_user_id bigint unsigned not NULL,
	primary key     (peak_id),
	foreign key     (creator_user_id) REFERENCES user(user_id)
) engine=InnoDB
COMMENT = 'Incomplete peak record. its reasonable to assume this will get much much larger';

create table `pin` (
	pin_id      bigint unsigned not NULL auto_increment,
	peak_id     bigint unsigned not NULL,
	user_id     bigint unsigned not NULL,
	name        varchar(255),
	date_added  datetime,
	creator     bigint unsigned not NULL,
	date        datetime,
	primary key (pin_id),
	foreign key (peak_id) REFERENCES peak(peak_id),
	foreign key (user_id) REFERENCES user(user_id)
) engine=InnoDB
COMMENT = 'Base PIN Super Class, further data is pulled in from data pulled in from pin_type';

create table `list` (
	list_id         bigint unsigned not NULL auto_increment,
	name            varchar(50) not NULL,
	creator_user_id bigint unsigned not NULL,
	creation_date   datetime,
	primary key     (list_id),
	foreign key     (creator_user_id) REFERENCES user(user_id)
) engine=InnoDB
COMMENT = 'List Data Super Class';

create table `list_linker` (
	linker_id    bigint unsigned not NULL auto_increment,
	list_id      bigint unsigned not NULL,
	user_id      bigint unsigned not NULL,
	peak_id      bigint unsigned not NULL,
	link_date    datetime,
	primary key  (linker_id),
	foreign key  (list_id) REFERENCES list(list_id),
	foreign key  (user_id) REFERENCES user(user_id),
	foreign key  (peak_id) REFERENCES peak(peak_id)
) engine=InnoDB
COMMENT = 'Linker Class, links pin objects with list objects';

create table `list_permission` (
	list_permission_id  bigint unsigned not NULL auto_increment,
	list_id             bigint unsigned not NULL,
	user_id             bigint unsigned,
	user_level          bigint unsigned,
	primary key         (list_permission_id),
	foreign key         (list_id) REFERENCES list(list_id),
	foreign key         (user_id) REFERENCES user(user_id)
) engine=InnoDB
COMMENT = 'grants non creator users permission to see lists';
