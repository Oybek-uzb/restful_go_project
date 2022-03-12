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

INSERT INTO book (name, author_id) VALUES ('Mehrobdan Chayon', '5b57be70-abe5-4a6d-ada8-534c2e46380a');
INSERT INTO book (name, author_id) VALUES ('Halqa', '1d51b267-8bed-429a-b2b5-80c5db415f23');
INSERT INTO book (name, author_id) VALUES ('1984', '18712c06-caab-4cc4-9a96-cfc5a23494ff');
