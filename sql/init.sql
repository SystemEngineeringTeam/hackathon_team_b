CREATE DATABASE IF NOT EXISTS with_b;
USE with_b;
create table if not exists lectures
(department varchar(10),semester varchar(10)
,dayofweek varchar(10),timed varchar(10),teacher varchar(20)
,lectureName varchar(20),indexLectureNumber INT AUTO_INCREMENT,PRIMARY KEY (indexLectureNumber));

insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','前期','月曜','3限','高木健太郎','日本国憲法');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','後期','火曜','4限','中村','情報数学');




create table if not exists grade
(grade varchar(10),indexLectureNumber INT);

insert into grade values (1,1);
insert into grade values (2,1);
insert into grade values (3,1);
insert into grade values (4,1);
insert into grade values (1,2);
insert into grade values (2,2);
insert into grade values (3,2);
insert into grade values (4,2);





insert into lectures values ('1', 'kk','後期','火曜','4限','中村','情報数学',2);

create table if not exists reviews
(indexLectureNumber int,reviewStar int,sentence varchar(100),id INT AUTO_INCREMENT,PRIMARY KEY (id));


insert into reviews (indexLectureNumber,reviewStar,sentence) values(1,5,'金を払えば単位がもらえる');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (2,1,'教科書の内容が理解できない');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,4,'難しい');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,3,'楽しい');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,4,'意味ワカンナイ');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,5,'虚無の申し子がお前たちを無に返すぞ');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,5,'SDPなしだぞ!?止めろよ');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (1,3,'いいね');





























