DROP TABLE IF EXISTS author CASCADE;
DROP TABLE IF EXISTS book CASCADE;

CREATE TABLE public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INT,
    is_alive BOOL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE INDEX idx_author_created_at_pagination ON public.author (created_at, id);
CREATE INDEX idx_author_age_pagination ON public.author (age, id);

CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INT,
    author_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);
CREATE INDEX idx_book_created_at_pagination ON public.book (created_at, id);
CREATE INDEX idx_book_age_pagination ON public.book (age, id);

INSERT INTO author (name, age, is_alive) VALUES ('Abdulla Qodiriy', 115, false);
INSERT INTO author (name, age, is_alive) VALUES ('Akrom Malik', 34, true);
INSERT INTO author (name, age, is_alive) VALUES ('George Orwell', 100, false);
INSERT INTO author (name, age, is_alive) VALUES ('Stephen Covey', 60, true);
INSERT INTO author (name, age, is_alive) VALUES ('Javlon Jovliyev', 34, true);
INSERT INTO author (name, age, is_alive) VALUES ('Dan Brown', 65, true);
INSERT INTO author (name, age, is_alive) VALUES ('Joanne Rowling', 60, true);
INSERT INTO author (name, age, is_alive) VALUES ('Alisher Navoiy', 581, false);
INSERT INTO author (name, age, is_alive) VALUES ('Mark Manson', 35, true);
INSERT INTO author (name, age, is_alive) VALUES ('Jaloliddin Rumiy', 630, false);
INSERT INTO author (name, age, is_alive) VALUES ('Cho`lpon', 120, false);
INSERT INTO author (name, age, is_alive) VALUES ('O`tkir Hoshimov', 90, false);
INSERT INTO author (name, age, is_alive) VALUES ('Shukrullo', 100, false);
INSERT INTO author (name, age, is_alive) VALUES ('Holid Husayniy', 45, true);
INSERT INTO author (name, age, is_alive) VALUES ('Stephen Edwin King', 74, true);
INSERT INTO author (name, age, is_alive) VALUES ('Imom Buxoriy', 800, false);
INSERT INTO author (name, age, is_alive) VALUES ('Muhammad Sodiq Muhammad Yusuf', 70, false);
INSERT INTO author (name, age, is_alive) VALUES ('Ramazon Butiy', 75, false);
INSERT INTO author (name, age, is_alive) VALUES ('Tobariy', 700, false);
INSERT INTO author (name, age, is_alive) VALUES ('Muhammad G`azzoliy', 600, false);
INSERT INTO author (name, age, is_alive) VALUES ('Brian Tracy', 60, true);
INSERT INTO author (name, age, is_alive) VALUES ('Fitrat', 115, false);
INSERT INTO author (name, age, is_alive) VALUES ('Oybek', 100, false);
INSERT INTO author (name, age, is_alive) VALUES ('Fyodor Dostayevskiy', 160, false);
INSERT INTO author (name, age, is_alive) VALUES ('Ahmad Lutfiy', 70, true);
INSERT INTO author (name, age, is_alive) VALUES ('Aristotel', 2427, false);
INSERT INTO author (name, age, is_alive) VALUES ('Abdulfattoh Abu G`udda', 105, false);
INSERT INTO author (name, age, is_alive) VALUES ('Ibn Kesir', 721, false);
INSERT INTO author (name, age, is_alive) VALUES ('Nouman Alikhon', 50, true);
INSERT INTO author (name, age, is_alive) VALUES ('Jack London', 166, false);
INSERT INTO author (name, age, is_alive) VALUES ('Theodore Dreiser', 85, false);
INSERT INTO author (name, age, is_alive) VALUES ('Imom Termiziy', 790, false);
INSERT INTO author (name, age, is_alive) VALUES ('Imom Muslim', 790, false);
INSERT INTO author (name, age, is_alive) VALUES ('Ibn Rushd', 894, false);
INSERT INTO author (name, age, is_alive) VALUES ('Abdukarim Mirzayev', 35, true);

INSERT INTO book (name, author_id) VALUES ('Mehrobdan Chayon', '21184482-b503-4dd9-aff8-875be000ded3');
INSERT INTO book (name, author_id) VALUES ('Halqa', '5922e8f0-78de-4578-af0b-1042c7e2be29');
INSERT INTO book (name, author_id) VALUES ('1984', '8c746211-fbea-4070-890d-962a7ae6331e');


-- page 1
SELECT
    *
FROM public.author
LIMIT 10;

-- page 2
SELECT
    *
FROM public.author
OFFSET 10
LIMIT 10;

-- page 3
SELECT
    *
FROM public.author
OFFSET 20
LIMIT 10;

-- page 4
SELECT
    *
FROM public.author
OFFSET 30
LIMIT 10;

-- page 1 (another way)
SELECT
    *
FROM public.author
ORDER BY created_at, id
LIMIT 10;

-- page 2 (another way)
SELECT
    *
FROM public.author
WHERE (created_at, id) > ('2022-03-14 15:37:39.102335' :: timestamp, '0ea33261-2ba0-4c5d-97ba-38f3ff055864')
ORDER BY created_at, id
LIMIT 10;

-- page 3 (another way)
SELECT
    *
FROM public.author
WHERE (created_at, id) > ('2022-03-14 15:37:39.317104' :: timestamp, 'aa3ec83f-0e7a-47f7-a848-ddb77a11c6ef')
ORDER BY created_at, id
LIMIT 10;

-- page 4 (another way)
SELECT
    *
FROM public.author
WHERE (created_at, id) > ('2022-03-14 15:37:39.460198' :: timestamp, 'ffb1fa7f-c8da-4a1f-bf98-fb266024b0a2')
ORDER BY created_at, id
LIMIT 10;