drop database if exists `local_search`;

create database `local_search`;

use `local_search`;

-- map
create table `maps`(
   `id` int not null AUTO_INCREMENT comment "id" ,
   'host' varchar(20) null,
   'guest' varchar(20) null ,

-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;
