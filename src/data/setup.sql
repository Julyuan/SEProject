drop table posts;
drop table threads;
drop table sessions;
drop table user;


create table user(
  Uid       varchar (20) not null,
  Upassword varchar(20) not null,
  Uavator   varchar(40) default 'avatar.jpg',
  Unickname varchar(20) default '匿名',
  Urealname varchar(10) default '',
  Usex      varchar(1)  default 'M',
  Uaddress  varchar(80) default '',
  Uphone    varchar(15) default '',
  Uemail    varchar(30) default '',
  Ubirthday date        default NULL,
  primary key (Uid)
);

create table sessions(
  id       serial primary key,
  uuid     varchar (64) not null unique,
  Uemail    varchar (255),
  Uid  integer references user(id),
  created_at timestamp not null
);