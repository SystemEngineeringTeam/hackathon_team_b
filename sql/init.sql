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
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('dd','0','月曜日','1限','倉橋','建設基礎数学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('dd','0','金曜日','3限','小池','土木情報リテラシ');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('dd','0','金曜日','2限','渡邊','土質力学1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('dd','0','金曜日','1限','山崎','交通工学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ee','0','月曜日','2限','島井','電気電子工学基礎');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ee','0','月曜日','2限','箕輪','アナログ回路１');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ee','0','金曜日','2限','道木','ディジタル回路１');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ee','0','火曜日','3限','鳥井','メカトロニクス');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fl','1','水曜日','4限','建部','施設計画1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fl','1','水曜日','3限','河路','住居設備1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fl','1','水曜日','3限','瀬古','木構造');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('fl','0','火曜日','2限','武田','建築図学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mm','1','木曜日','3限','太田','機械製図2');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mm','1','金曜日','4限','三宅','エネルギー変換工学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mm','0','木曜日','3限','田中','計測工学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mm','0','木曜日','1限','高田','先端材料');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mp','0','月曜日','3限','小方','技術者論理');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mp','0','水曜日','2限','山田','機講学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mp','1','木曜日','1限','椿野','制御工学応用');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('mp','1','月曜日','2限','三宅','数値解析法');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ds','0','水曜日','1限','山本','測量学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ds','0','木曜日','1限','倉橋','防災工学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ds','0','火曜日','2限','岩月','鉄筋コンクリート構造１');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ds','0','木曜日','2限','赤堀','水文学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ev','1','水曜日','1限','青木','電気磁気学及び演習');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ev','1','火曜日','3限','五島','電子情報工学実験1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ev','0','金曜日','4限','増田','化学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ev','1','木曜日','1限','徳田','電子デバイス工学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('hh','1','木曜日','2限','菊池','スポーツマネジメント論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('hh','0','木曜日','4限','田中','スポーツ運営論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('hh','0','水曜日','1限','青嶋','ビジネスマネジメント論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('hh','1','金曜日','1限','水谷','生涯スポーツ経営論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ht','0','月曜日','4限','吉成','経営学概論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ht','0','月曜日','4限','ロバート クレイロン','ビジネス英語１');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ht','1','木曜日','3限','小森','会計学基礎論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('ht','1','月曜日','3限','吉成','マクロ経済学');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','0','火曜日','3限','伊藤','コンピュータアーキテクチャ1');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','0','金曜日','5限','内藤','モバイルネットワーク');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','1','水曜日','4限','梶','情報システム概論');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kk','0','月曜日','2限','小野木','言語理論及びコンパイラ');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','0','水曜日','4限','河辺','情報セキュリティ');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','1','水曜日','1限','北坂','映像制作及び演習');
insert into lectures (department,semester,dayofweek,timed,teacher,lectureName) values ('kx','1','金曜日','1限','鳥居 一平','デッサンとモデリング');



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
insert into grade values (1,26);
insert into grade values (2,26);
insert into grade values (3,26);
insert into grade values (4,26);
insert into grade values (1,27);
insert into grade values (2,27);
insert into grade values (3,27);
insert into grade values (4,27);
insert into grade values (2,28);
insert into grade values (3,28);
insert into grade values (4,28);
insert into grade values (3,29);
insert into grade values (4,29);
insert into grade values (1,30);
insert into grade values (2,30);
insert into grade values (3,30);
insert into grade values (4,30);
insert into grade values (2,31);
insert into grade values (3,31);
insert into grade values (4,31);
insert into grade values (2,32);
insert into grade values (3,32);
insert into grade values (4,32);
insert into grade values (3,33);
insert into grade values (4,33);
insert into grade values (2,34);
insert into grade values (3,34);
insert into grade values (4,34);
insert into grade values (2,35);
insert into grade values (3,35);
insert into grade values (4,35);
insert into grade values (2,36);
insert into grade values (3,36);
insert into grade values (4,36);
insert into grade values (1,37);
insert into grade values (2,37);
insert into grade values (3,37);
insert into grade values (4,37);
insert into grade values (1,38);
insert into grade values (3,39);
insert into grade values (4,39);
insert into grade values (3,40);
insert into grade values (4,40);
insert into grade values (4,41);
insert into grade values (1,42);
insert into grade values (2,42);
insert into grade values (3,42);
insert into grade values (4,42);
insert into grade values (2,43);
insert into grade values (3,43);
insert into grade values (4,43);
insert into grade values (3,44);
insert into grade values (4,44);
insert into grade values (3,45);
insert into grade values (4,45);
insert into grade values (1,46);
insert into grade values (2,46);
insert into grade values (3,46);
insert into grade values (4,46);
insert into grade values (2,47);
insert into grade values (3,47);
insert into grade values (4,47);
insert into grade values (3,48);
insert into grade values (4,48);
insert into grade values (3,49);
insert into grade values (4,49);
insert into grade values (1,50);
insert into grade values (3,51);
insert into grade values (4,51);
insert into grade values (1,52);
insert into grade values (2,52);
insert into grade values (3,52);
insert into grade values (4,52);
insert into grade values (3,53);
insert into grade values (4,53);
insert into grade values (1,54);
insert into grade values (2,54);
insert into grade values (3,54);
insert into grade values (4,54);
insert into grade values (3,55);
insert into grade values (4,55);
insert into grade values (3,56);
insert into grade values (4,56);
insert into grade values (2,57);
insert into grade values (3,57);
insert into grade values (4,57);
insert into grade values (1,58);
insert into grade values (2,58);
insert into grade values (3,58);
insert into grade values (4,58);
insert into grade values (2,59);
insert into grade values (3,59);
insert into grade values (4,59);
insert into grade values (1,60);
insert into grade values (2,60);
insert into grade values (3,60);
insert into grade values (4,60);
insert into grade values (2,61);
insert into grade values (3,61);
insert into grade values (4,61);
insert into grade values (2,62);
insert into grade values (3,62);
insert into grade values (4,62);
insert into grade values (2,63);
insert into grade values (3,63);
insert into grade values (4,63);
insert into grade values (2,64);
insert into grade values (3,64);
insert into grade values (4,64);
insert into grade values (2,65);
insert into grade values (3,65);
insert into grade values (4,65);
insert into grade values (2,66);
insert into grade values (3,66);
insert into grade values (4,66);
insert into grade values (3,67);
insert into grade values (4,67);
insert into grade values (1,68);
insert into grade values (2,68);
insert into grade values (3,68);
insert into grade values (4,68);



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
insert into reviews (indexLectureNumber,reviewStar,sentence) values (7,1,"ダメだった");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (8,3,"単位は貰えます");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (9,5,"過去問通りのテストでした");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (10,5,"教授が優しくて最高");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (11,5,"かも");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (12,1,"単位落ちました");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (13,3,"普通の授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (14,4,"何か物足りない感じがする");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (15,2,"単位はくれるけど秀は絶対にくれないです");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (16,3,"普通にやれば普通に成績が出る");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (17,5,"秀もらいやすい授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (18,4,"単位もらいやすい授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (19,1,"やばい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (20,2,"教授が厳しい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (21,1,"辛い");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (22,2,"なかなかしんどい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (23,4,"やりがいはある");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (24,4,"やりがいはある");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (25,3,"しんどいけど単位はもらえる");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (26,3,"good");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (27,5,"流石にこの科目落とすやつはおらんwwwww");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (28,4,"落単");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (29,2,"難しい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (30,2,"つまらん");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (31,3,"面白い");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (32,3,"面白い");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (33,3,"簡単でした！");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (34,1,"この授業はできれば取らない方がいいです");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (35,5,"良かった");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (36,2,"落単したので星2");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (37,1,"むずすぎ");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (38,4,"簡単すぎて逆にほし4");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (39,4,"面白い授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (40,2,"単位落ちました");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (41,4,"単位取れた");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (42,5,"最高");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (43,3,"普通");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (44,4,"教科書通りの授業とテスト");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (45,4,"何か物足りない感じがする");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (46,4,"普通にしていれば大丈夫");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (47,1,"難しい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (48,3,"好き嫌いがあると思います");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (49,2,"教授が厳しい人です");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (50,5,"単位欲しい人にはおすすめ");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (51,4,"楽な授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (52,4,"教授が優しいです");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (53,2,"しんどい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (54,3,"good");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (55,1,"やばい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (56,1,"つまらん");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (57,4,"楽しい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (58,4,"普通にいい授業");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (59,3,"教授ガチャ");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (60,3,"教授ガチャ");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (61,5,"とてもわかりやすい授業だった");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (62,5,"楽しく学べる授業!!");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (63,5,"とてもタメになる授業！！");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (64,5,"この授業だけは取った方がいい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (65,5,"最高");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (66,5,"とてもいい");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (67,5,"教授が良い");
insert into reviews (indexLectureNumber,reviewStar,sentence) values (68,5,"ここでは多く語らない");































