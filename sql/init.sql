CREATE DATABASE IF NOT EXISTS with_b;
USE with_b;
create table if not exists lectures
(department varchar(10),semester varchar(10)
,dayofweek varchar(10),timed varchar(10),teacher varchar(20)
,lectureName varchar(20),indexLectureNumber INT AUTO_INCREMENT,PRIMARY KEY (indexLectureNumber));

insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','0','月曜日','3限','高木健太郎','日本国憲法');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','火曜日','4限','中村','情報数学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','1','水曜日','3限','秦','ネットワーク');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','木曜日','5限','鈴木','データベース');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','金曜日','2限','福田','アーキテクチャ');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','月曜日','1限','外山','幾何学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','火曜日','2限','山口','オブジェクト志向');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','水曜日','3限','松田','モバイル');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','水曜日','3限','春名','知的財産権');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','1','金曜日','1限','山本','メディア文化論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','1','木曜日','3限','久留宮','インターネットビジネス論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','0','月曜日','1限','中條','組み込みシステム概論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ev','1','水曜日','4限','森','電気実用英語');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fa','0','火曜日','2限','河路','建築設備Ⅰ');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fa','1','火曜日','1限','杉野','日本建築史1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fa','1','月曜日','3限','野澤','都市計画');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fa','0','月曜日','3限','安井','建築設計2');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cb','0','金曜日','2限','岡本','工学基礎数理');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cb','1','木曜日','1限','平野','無機化学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cb','0','木曜日','4限','大澤','薬理学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cb','1','木曜日','2限','飯島','生体情報科学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cc','1','金曜日','1限','佐藤','高分子化学1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cc','1','火曜日','1限','手嶋','分析化学1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cc','1','金曜日','4限','小林','基礎科学実験');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('cc','1','金曜日','2限','村田','有機化学演習');









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
insert into grade values (2,3);
insert into grade values (2,4);
insert into grade values (2,5);
insert into grade values (4,5);
insert into grade values (4,6);
insert into grade values (2,7);
insert into grade values (4,7);
insert into grade values (4,8);
insert into grade values (3,9);
insert into grade values (4,9);
insert into grade values (2,10);
insert into grade values (3,10);
insert into grade values (4,10);
insert into grade values (3,11);
insert into grade values (4,11);
insert into grade values (3,12);
insert into grade values (3,12);
insert into grade values (2,13);
insert into grade values (3,13);
insert into grade values (4,13);
insert into grade values (1,14);
insert into grade values (2,14);
insert into grade values (3,14);
insert into grade values (4,14);
insert into grade values (3,15);
insert into grade values (4,15);
insert into grade values (3,16);
insert into grade values (4,16);
insert into grade values (3,17);
insert into grade values (1,18);
insert into grade values (2,18);
insert into grade values (1,19);
insert into grade values (2,19);
insert into grade values (3,19);
insert into grade values (4,19);
insert into grade values (4,20);
insert into grade values (2,21);
insert into grade values (3,21);
insert into grade values (4,21);
insert into grade values (2,22);
insert into grade values (2,23);
insert into grade values (3,23);
insert into grade values (4,23);
insert into grade values (2,24);
insert into grade values (3,24);
insert into grade values (4,24);
insert into grade values (2,25);
insert into grade values (3,25);
insert into grade values (4,25);










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
insert into reviews (indexLectureNumber,reviewStar,sentence) values (3,3,'普通');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (4,2,'難しい');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (5,5,'インターネットビジネスって感じだった');
insert into reviews (indexLectureNumber,reviewStar,sentence) values (6,3,'とてもタメになった');






























