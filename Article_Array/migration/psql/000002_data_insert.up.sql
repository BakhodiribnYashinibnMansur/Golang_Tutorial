
INSERT INTO author (firstname,lastname) VALUES ('Aziz','Almamatov') ON CONFLICT DO NOTHING;
INSERT INTO author (firstname,lastname) VALUES ('Diyor','Almamatov') ON CONFLICT DO NOTHING;

INSERT INTO article (title,body,author_id) VALUES ('Talaba','4-kurs',1) ON CONFLICT DO NOTHING;
INSERT INTO article (title,body,author_id) VALUES ('Talaba emas','bitirgan',2) ON CONFLICT DO NOTHING;
