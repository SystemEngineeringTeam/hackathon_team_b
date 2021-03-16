CREATE DATABASE IF NOT EXISTS with_b;
USE with_b;
create table if not exists rectures
(grade varchar(10),department varchar(10),semester varchar(10)
,dayofweek varchar(10),time varchar(10),teacher varchar(20)
,lectureName varchar(20),reviewStarAverage int,indexNumber int);

insert into rectures values ('1', 'kk','前期','月曜','3限','高木健太郎','日本国憲法',5,1);
insert into rectures values ('1', 'kk','後期','火曜','4限','中村','情報数学',4,2);



create table if not exists reviews
(indexNumber int,reviewStar int,sentence varchar(100));

insert into reviews values (1,5,'金を払えば単位がもらえる');
insert into reviews values (2,1,'教科書の内容が理解できない');
























