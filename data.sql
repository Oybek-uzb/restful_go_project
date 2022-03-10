CREATE TABLE public.author (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name VARCHAR(100) NOT NULL,
                             author_id UUID NOT NULL ,
                             CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('Abdulla Qodiriy');
INSERT INTO author (name) VALUES ('Akrom Malik');
INSERT INTO author (name) VALUES ('George Orwell');

INSERT INTO book (name, author_id) VALUES ('Mehrobdan Chayon', '89f44248-c821-40d6-bb1c-9fc7d1f747c4');
INSERT INTO book (name, author_id) VALUES ('Halqa', '57d88b1b-7305-46ed-a5a4-f5e8200ec46d');
INSERT INTO book (name, author_id) VALUES ('1984', 'a9dd2677-6d9a-4d12-9034-f1ee8f4b4b11');
