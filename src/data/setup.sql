drop table posts;
drop table threads;
drop table sessions;
drop table user;

CREATE TABLE address (
  Aid varchar(50) NOT NULL,
  Uid varchar(20) NOT NULL,
  Areceivername varchar(10) NOT NULL,
  Aaddress varchar(100) NOT NULL,
  Acode varchar(10) NOT NULL,
  Aphone varchar(15) DEFAULT '',
  Afixphone varchar(15) DEFAULT '',
  Aprovince varchar(10) NOT NULL,
  Acity varchar(10) NOT NULL,
  Atown varchar(10) NOT NULL,
  Acheck tinyint(2) NOT NULL,
  PRIMARY KEY (Aid),
  FOREIGN KEY (Uid) REFERENCES user (Uid)
);

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
  id       int(5) primary key not null auto_increment,
  uuid     varchar (64) not null unique,
  Uemail    varchar (255),
  Uid  varchar(20) references user(Uid),
  created_at timestamp not null
);

create table shop(
  Sid int(5) not null auto_increment,
  Sname varchar(50) not null default '匿名店铺',
  Uid   varchar(20) default null references user(Uid),
  Sicon varchar(25) default 'shopicon.jpg',
  Ssummary  varchar(1000) default '暂无',
  Sactivity varchar(100)  default '暂无',
  Stransprice float(10, 2) unsigned zerofill default '0000000.00',
  primary key(Sid),
  FOREIGN KEY (Uid) REFERENCES user (Uid)
);

create table book(
  Bid varchar(50) not null,
  Bimage varchar(50) not null,
  Bprice float(10,2) not null,
  Bname  varchar(100) not null,
  Btype  varchar(10)  not null,
  Sid    int(5) not null,
  Bauthor varchar(30) default '',
  Bpublisher  varchar(30) default '',
  Bpublishdate date default '2004-10-3',
  Bsalednum int(5) default null,
  Bcommentnum int(5) default null,
  Boriprice float(10,2) default '0.00',
  Bversion  int(5) default '0',
  Bpagenum int(11) default '0',
  Bwordnum int(11) default '0',
  Bprintdate  date default '2004-2-3',
  Bsize   int(5) default '0',
  Bpaperstyle varchar(10) default '',
  Bprintnum int(5) default null,
  Bpackage  varchar(10) default '',
  Bisbn varchar(20) default '',
  Bcontentsummary varchar(1000) default '暂无',
  Bauthorsummary varchar(1000) DEFAULT '暂无',
  Bmediacomment varchar(1000) DEFAULT '暂无',
  Btastecontent varchar(8000) DEFAULT '暂无',
  Bstocknum int(5) DEFAULT '0',
  Buploaddate date DEFAULT '1993-11-08',
  PRIMARY KEY (Bid),
  FOREIGN KEY (Sid) REFERENCES shop (Sid)
);

CREATE TABLE `orderform` (
  `Oid` varchar(50) NOT NULL,
  `Uid` varchar(20) DEFAULT NULL,
  `Sid` int(5) DEFAULT NULL,
  `Aid` varchar(50) DEFAULT NULL,
  `Ototalbooksprice` float(10,2) DEFAULT NULL,
  `Ototaltransprice` float(10,2) DEFAULT NULL,
  `Osubmittime` datetime DEFAULT NULL,
  `Opaytime` datetime DEFAULT NULL,
  `Oreceivetime` datetime DEFAULT NULL,
  `Ofinishedtime` datetime DEFAULT NULL,
  `Ostatus` tinyint(2) DEFAULT NULL,
  PRIMARY KEY (`Oid`),

  FOREIGN KEY (`Aid`) REFERENCES `address` (`Aid`) ON DELETE SET NULL,
  FOREIGN KEY (`Sid`) REFERENCES `shop` (`Sid`),
  FOREIGN KEY (`Uid`) REFERENCES `user` (`Uid`)
);

CREATE TABLE `transaction` (
  `Tid` int(5) NOT NULL AUTO_INCREMENT,
  `Bid` varchar(50) NOT NULL,
  `Oid` varchar(50) DEFAULT NULL,
  `Uid` varchar(20) DEFAULT NULL,
  `Tstatus` tinyint(2) DEFAULT NULL,
  `Tboughtnum` int(5) NOT NULL,
  `Tmark` int(1) DEFAULT NULL,
  `Tcomment` varchar(1000) DEFAULT '',
  `Tsubmittime` datetime DEFAULT NULL,
  `Tpaytime` datetime DEFAULT NULL,
  `Treceivetime` datetime DEFAULT NULL,
  `Tcommenttime` datetime DEFAULT NULL,
  PRIMARY KEY (`Tid`),
  FOREIGN KEY (`Oid`) REFERENCES `orderform` (`Oid`) ON DELETE SET NULL,
  FOREIGN KEY (`Uid`) REFERENCES `user` (`Uid`)
);

