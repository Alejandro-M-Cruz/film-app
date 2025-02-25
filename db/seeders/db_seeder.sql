INSERT INTO users (username, password)
VALUES
    ('user1', 'password'),
    ('user2', 'password');


INSERT INTO films (user_id, title, release_date, genre, director, cast, synopsis)
VALUES
    (1, 'The Shawshank Redemption', '1994-09-23', 'Drama', 'Frank Darabont', 'Tim Robbins;Morgan Freeman;Bob Gunton', 'Two imprisoned'),
    (1, 'The Godfather', '1972-03-24', 'Crime', 'Francis Ford Coppola', 'Marlon Brando;Al Pacino;James Caan', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.'),
    (1, 'The Dark Knight', '2008-07-18', 'Action', 'Christopher Nolan', 'Christian Bale;Heath Ledger;Aaron Eckhart', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept');
