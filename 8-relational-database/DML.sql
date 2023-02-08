-- insert / memasukkan data ke table
INSERT INTO guru(id, nama, telepon)
VALUES ("GR0001", "Budi", "0812345");

INSERT INTO guru(id, nama, telepon, email)
VALUES ("GR0002", "Rudi", "0812346","rudi@mail.com");

INSERT INTO guru(id, nama, telepon, email)
VALUES ("GR0003", "Adi", "0812346","adi@mail.com");

INSERT INTO guru(id, nama, telepon, email)
VALUES ("GR0004", "Andi", "0812346","andi@mail.com");

INSERT INTO mata_pelajaran(nama_mapel, kkm, keterangan)
VALUES ("Matematika Kelas 1", 85, "Mata pelajaran Matematika untuk kelas 1"),
("Matematika Kelas 2", 85, "Mata pelajaran Matematika untuk kelas 2"),
("Bahasa Indonesia Kelas 1", 85, "Mata pelajaran Bahasa Indonesia untuk kelas 1");


INSERT INTO mata_pelajaran(nama_mapel, kkm, keterangan)
VALUES ("Pemrograman Dasar", 70, "Mata pelajaran Pemrograman Dasar");
INSERT INTO mata_pelajaran(nama_mapel, kkm, keterangan)
VALUES ("Statistika", 50, "Mata pelajaran Statistika");

-- error: nama_mapel tidak boleh kosong
INSERT INTO mata_pelajaran(kkm, keterangan)
VALUES ( 85, "Mata pelajaran Matematika untuk kelas 1");

INSERT INTO data_mengajar(guru_id, mapel_id, jam, hari)
VALUES ("GR0001", 1, "07:00", "SENIN");

INSERT INTO data_mengajar(guru_id, mapel_id, jam, hari)
VALUES ("GR0001", 2, "07:00", "SELASA");

INSERT INTO data_mengajar(guru_id, mapel_id, jam, hari)
VALUES ("GR0003", 3, "09:00", "SELASA");

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

-- DML Statement
SELECT * FROM guru;
SELECT * FROM mata_pelajaran;
SELECT * FROM data_mengajar;
-- LIKE
SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%';
SELECT * FROM mata_pelajaran where nama_mapel like '%Kelas%';

-- BETWEEN
SELECT * FROM mata_pelajaran where id BETWEEN 1 and 2;

-- AND / OR
-- and : semua kondisi harus terpenuhi
-- or : salah satu terpenuhi, maka data akan muncul
SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%' AND kkm=85;
SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%' OR kkm=85;

SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%' AND kkm=85 OR kkm=70;

-- ORDER BY
SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%' AND kkm=85 OR kkm=70 ORDER BY kkm ASC;
select * from mata_pelajaran order by id DESC;

-- LIMIT
SELECT * FROM mata_pelajaran where nama_mapel like 'Matematika%' AND kkm=85 OR kkm=70 ORDER BY kkm ASC LIMIT 1;
select * from mata_pelajaran order by id DESC LIMIT 3;

-- JOIN
-- INNER JOIN
SELECT data_mengajar.id, data_mengajar.guru_id, guru.nama, data_mengajar.mapel_id, data_mengajar.jam, data_mengajar.hari FROM data_mengajar
INNER JOIN guru ON data_mengajar.guru_id = guru.id;

-- membuat alias table, untuk join
SELECT dm.id, dm.guru_id, gr.nama, dm.mapel_id, dm.jam, dm.hari FROM data_mengajar dm
INNER JOIN guru gr ON dm.guru_id = gr.id;
-- inner join ke lebih dari 2 table
SELECT dm.id, dm.guru_id, gr.nama, dm.mapel_id, mp.nama_mapel, mp.kkm, dm.jam, dm.hari FROM data_mengajar dm
INNER JOIN guru gr ON dm.guru_id = gr.id
INNER JOIN mata_pelajaran mp ON dm.mapel_id = mp.id;

-- LEFT JOIN
-- menampilkan seluruh data di table kiri
SELECT gr.nama, gr.telepon, gr.email, dm.id, dm.guru_id, dm.mapel_id, dm.jam, dm.hari FROM data_mengajar dm
LEFT JOIN guru gr ON dm.guru_id = gr.id;

SELECT gr.id as id_guru, gr.nama as nama_guru, gr.telepon, gr.email, dm.id as id_data_mengajar, dm.guru_id, dm.mapel_id, dm.jam, dm.hari FROM guru gr
LEFT JOIN data_mengajar dm ON gr.id = dm.guru_id;

SELECT gr.id as id_guru, gr.nama as nama_guru, gr.telepon, gr.email, dm.id as id_data_mengajar, dm.guru_id, dm.mapel_id, mp.nama_mapel, dm.jam, dm.hari FROM guru gr
LEFT JOIN data_mengajar dm ON gr.id = dm.guru_id
LEFT JOIN mata_pelajaran mp ON dm.mapel_id=mp.id;

-- RIGHT JOIN
SELECT gr.id as id_guru, gr.nama as nama_guru, gr.telepon, gr.email, dm.id as id_data_mengajar, dm.guru_id, dm.mapel_id, dm.jam, dm.hari FROM guru gr
RIGHT JOIN data_mengajar dm ON gr.id = dm.guru_id;

-- dua query ini akan menampilkan data yang sama
-- #1
SELECT gr.nama, gr.telepon, gr.email, dm.id, dm.guru_id, dm.mapel_id, dm.jam, dm.hari FROM data_mengajar dm
RIGHT JOIN guru gr ON dm.guru_id = gr.id;
-- #2
SELECT gr.id as id_guru, gr.nama as nama_guru, gr.telepon, gr.email, dm.id as id_data_mengajar, dm.guru_id, dm.mapel_id, dm.jam, dm.hari FROM guru gr
LEFT JOIN data_mengajar dm ON gr.id = dm.guru_id;

select gr.id as id_guru, gr.nama as nama_guru, gr.telepon, gr.email, dm.id as id_data_mengajar, dm.guru_id, dm.mapel_id, mp.nama_mapel, dm.jam, dm.hari FROM guru gr
LEFT JOIN data_mengajar dm ON gr.id = dm.guru_id
RIGHT JOIN mata_pelajaran mp ON dm.mapel_id=mp.id;


-- UNION
SELECT City FROM Customers
UNION
SELECT City FROM Suppliers
ORDER BY City;

select nama from guru
UNION
select nama_mapel from mata_pelajaran;

-- AGGREGATION
SELECT * FROM guru;
SELECT * FROM mata_pelajaran;
SELECT * FROM data_mengajar;
-- MIN
select MIN(kkm) from mata_pelajaran;
select MIN(created_at) from data_mengajar;

-- MAX
select MAX(kkm) from mata_pelajaran;
select MAX(created_at) from data_mengajar;

-- SUM
select SUM(kkm) from mata_pelajaran where id between 1 and 4;

-- AVG
select AVG(kkm) from mata_pelajaran where id between 1 and 4;

-- COUNT
select count(id) as jumlah_pertemuan from data_mengajar where guru_id='GR0001';

-- HAVING
-- menampilkan data guru yang mempunyai jumlah pertemuan lebih dari 1
select gr.id, gr.nama, gr.email, count(dm.id) as jumlah_mengajar from guru gr
inner join data_mengajar dm ON gr.id = dm.guru_id
GROUP BY gr.id
HAVING count(dm.id) > 1;

select gr.id, gr.nama, gr.email, count(dm.id) as jumlah_mengajar from guru gr
inner join data_mengajar dm ON gr.id = dm.guru_id
GROUP BY gr.id
HAVING count(dm.id) = 1;

-- sample aggregasi count dan sum bersamaan.
select gr.id, gr.nama, gr.email, count(dm.id) as jumlah_mengajar, sum(mp.kkm) as jumlah_kkm from guru gr
inner join data_mengajar dm ON gr.id = dm.guru_id
inner join mata_pelajaran mp ON dm.mapel_id = mp.id
GROUP BY gr.id
HAVING count(dm.id) = 1
ORDER BY gr.id DESC
LIMIT 2;