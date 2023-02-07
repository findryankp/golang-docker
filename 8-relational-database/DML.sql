-- insert / memasukkan data ke table
INSERT INTO guru(id, nama, telepon)
VALUES ("GR0001", "Budi", "0812345");

INSERT INTO guru(id, nama, telepon, email)
VALUES ("GR0002", "Rudi", "0812346","rudi@mail.com");

INSERT INTO guru(id, nama, telepon, email)
VALUES ("GR0003", "Adi", "0812346","adi@mail.com");

INSERT INTO mata_pelajaran(nama_mapel, kkm, keterangan)
VALUES ("Matematika Kelas 1", 85, "Mata pelajaran Matematika untuk kelas 1"),
("Matematika Kelas 2", 85, "Mata pelajaran Matematika untuk kelas 2"),
("Bahasa Indonesia Kelas 1", 85, "Mata pelajaran Bahasa Indonesia untuk kelas 1");

-- error: nama_mapel tidak boleh kosong
INSERT INTO mata_pelajaran(kkm, keterangan)
VALUES ( 85, "Mata pelajaran Matematika untuk kelas 1");

INSERT INTO data_mengajar(guru_id, mapel_id, jam, hari)
VALUES ("GR0001", 1, "07:00", "SENIN");

INSERT INTO data_mengajar(guru_id, mapel_id, jam, hari)
VALUES ("GR0005", 1, "07:00", "SENIN");

-- membaca/menampilkan data
SELECT * FROM guru;
SELECT * FROM mata_pelajaran;
SELECT * FROM data_mengajar;

select id, nama from guru;

select id, nama from guru where id="GR0001";

-- menampilkan data guru yang email nya tidak null
select id, nama, email from guru where email is not null;

-- UPDATE data
UPDATE guru SET 
nama="Andi" 
WHERE id="GR0001";

UPDATE guru SET 
nama="Andi",
email="andi@mail.com"
WHERE id="GR0001";

-- DELETE 
DELETE FROM guru
WHERE id="GR0001";

DELETE FROM data_mengajar
WHERE id=1;