-- DROP SCHEMA bogoso;

CREATE SCHEMA bogoso AUTHORIZATION postgres;

-- bogoso.elevy_request definition

-- Drop table

-- DROP TABLE bogoso.cv_files;

CREATE TABLE bogoso.cv_files (
	id varchar NOT NULL,
	applicant_name varchar NOT NULL,
	email varchar NOT NULL,
	phone varchar NOT NULL,
	file_name varchar NOT NULL,
	created timestamptz NULL DEFAULT now(),
	CONSTRAINT cv_files_pk PRIMARY KEY (id)
);

insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2949','John Bolton','j.bolton@politics.com','0210641124','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3054','Don Lenon','dlenon@news.com','0211090013','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3159','Brian Adams','bradams@music.com','0211538902','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3264','Marvin Gaye','mgaye@arts.com','0211987791','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3369','Abraham Lincoln','alincoln@lead.com','0212436680','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3474','Majorie Greene','mgreene@lies.com','0212885569','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3579','Kevin Taylor','ktaylor@yahoo.com','0213334458','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3684','John Plumber','jplumber@yahoo.com','0213783347','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3789','Eric Carpenter','ecarpenter@yahoo.com','0214232236','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3894','Joseph Mason','j.mason@icloud.com','0214681125','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3999','Alexander Graham','a.graham@arts.com','0215130014','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('4104','Kojo Antwi','kantwi@music.com','0215578903','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('4209','Rex Omar','romar@music.com','0216027792','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0020','Nat Brew','nbrew@music.com','0216476681','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0031','Celine Dion','cdion@music.com','0216925570','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0042','Nicholas Taleb','Nntb@data.com','0217374459','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0053','Peter Diamandis','peter@abundance.com','0217823348','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0064','Elon Musk','elon@rockets.com','0218272237','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0075','Jack Folson','Jack.folson@live.com','0218721126','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0086','Mark Zuck','mzuck@meta.com','0219170015','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0100','John Doe','jdoe@outlook.com','0243010123','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2010','Jane Doh','jdoh@live.com','0201231234','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0960','Kofi Ghana','kghana@yahoo.com','0243112244','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('0090','Amina Zongo','aminaz@icloud.com','0284993254','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1140','Esi Fadama','esifa@gmail.com','0326874264','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2190','Eric Tula','etula@gmail.com','0368755274','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('3240','Kofi Lion','kolion@hotmail.com','0410636284','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('4290','Jake Doe','jaked@outlook.com','0452517294','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('5340','Kevin Ghana','kevghana@live.com','0494398304','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('6390','Joyce Doe','joyced@yandex.com','0536279314','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('7440','Janet Doh','janetdoe@alibaba.com','0578160324','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('8490','June Tulsa','jtulsa@icloud.com','0201663344','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('9540','Lenny Kravitz','lkravitz@music.com','0202112233','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1059','Lucky Dube','ldube@reggae.com','0202561122','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1164','Jimmy Cliff','jcliff@reggae.com','0203010011','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1269','Peter Tosh','ptosh@reggae.com','0203458900','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1374','Steel Pulse','spulse@reggae.com','0203907789','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1479','Kenny G','kg@jazz.com','0204356678','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1584','Afua Sampana','afua.sampana@live.com','0204805567','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1689','Ata kakra','kakraa@yandex.com','0205254456','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1794','Kweku Ananse','kweku.ananse@arts.com','0205703345','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('1899','Kwame Ntekuma','k.ntekuma@icloud.com','0206152234','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2004','Thomas Sankara','tsankara@politics.com','0206601123','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2109','Patrice Lumumba','patlumumba@live.com','0207050012','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2214','Che Cuevera','ccuevera@politics.com','0207498901','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2319','Kwame Nkrumah','kwame.nkrumah@live.com','0207947790','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2424','Malcom X','malcomx@outlook.com','0208396679','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2529','Richard Nixon','r.nixon@icloud.com','0208845568','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2634','Donald Trump','dtrump@lies.com','0209294457','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2739','Barack Obama','bobama@hope.com','0209743346','uploaded file');
insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values('2844','Margaret Thatcher','mthatcher@lead.com','0210192235','uploaded file');

CREATE TABLE bogoso.temp_token (
	email varchar NOT NULL,
	token varchar NOT NULL,
	created timestamptz NULL DEFAULT now()
);

