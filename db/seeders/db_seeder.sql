INSERT INTO users (username, password)
VALUES
    ('user1', 'password'),
    ('user2', 'password');


INSERT INTO films (user_id, title, release_date, genre, director, cast, synopsis)
VALUES
    (1, 'The Shawshank Redemption', '1994-09-23', 'Drama', 'Frank Darabont', 'Tim Robbins;Morgan Freeman;Bob Gunton', 'Two imprisoned'),
    (1, 'The Godfather', '1972-03-24', 'Crime', 'Francis Ford Coppola', 'Marlon Brando;Al Pacino;James Caan', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.'),
    (1, 'The Dark Knight', '2008-07-18', 'Action', 'Christopher Nolan', 'Christian Bale;Heath Ledger;Aaron Eckhart', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept'),
    (1, 'Pulp Fiction', '1994-10-14', 'Crime', 'Quentin Tarantino', 'John Travolta;Uma Thurman;Samuel L. Jackson', 'The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.'),
    (1, 'Fight Club', '1999-10-15', 'Drama', 'David Fincher', 'Brad Pitt;Edward Norton;Helena Bonham Carter', 'An insomniac office worker and a devil-may-care soap maker form an underground fight club that evolves into much more.'),
    (1, 'Forrest Gump', '1994-07-06', 'Drama', 'Robert Zemeckis', 'Tom Hanks;Robin Wright;Gary Sinise', 'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold from the perspective of an Alabama man with an IQ of 75.'),
    (1, 'Inception', '2010-07-16', 'Sci-Fi', 'Christopher Nolan', 'Leonardo DiCaprio;Joseph Gordon-Levitt;Elliot Page', 'A thief who enters the dreams of others to steal secrets is given the inverse task of planting an idea into the mind of a CEO.'),
    (1, 'The Matrix', '1999-03-31', 'Sci-Fi', 'The Wachowskis', 'Keanu Reeves;Laurence Fishburne;Carrie-Anne Moss', 'A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.'),
    (1, 'Goodfellas', '1990-09-19', 'Crime', 'Martin Scorsese', 'Robert De Niro;Ray Liotta;Joe Pesci', 'The story of Henry Hill and his life in the mafia, covering his relationship with his wife Karen Hill and his mob partners.'),
    (1, 'The Silence of the Lambs', '1991-02-14', 'Thriller', 'Jonathan Demme', 'Jodie Foster;Anthony Hopkins;Scott Glenn', 'A young FBI cadet must receive the help of an incarcerated and manipulative cannibal killer to catch another serial killer.'),
    (1, 'Se7en', '1995-09-22', 'Thriller', 'David Fincher', 'Morgan Freeman;Brad Pitt;Kevin Spacey', 'Two detectives, a rookie and a veteran, hunt a serial killer who uses the seven deadly sins as his motives.'),
    (1, 'The Green Mile', '1999-12-10', 'Drama', 'Frank Darabont', 'Tom Hanks;Michael Clarke Duncan;David Morse', 'The lives of guards on Death Row are affected by one of their charges: a black man accused of child murder and rape, yet who has a mysterious gift.'),
    (1, 'Interstellar', '2014-11-07', 'Sci-Fi', 'Christopher Nolan', 'Matthew McConaughey;Anne Hathaway;Jessica Chastain', 'A team of explorers travel through a wormhole in space in an attempt to ensure humanity''s survival.'),
    (1, 'The Usual Suspects', '1995-07-19', 'Mystery', 'Bryan Singer', 'Kevin Spacey;Gabriel Byrne;Chazz Palminteri', 'A sole survivor tells of the twisty events leading up to a horrific gun battle on a boat, which began when five criminals met at a police lineup.'),
    (1, 'Gladiator', '2000-05-05', 'Action', 'Ridley Scott', 'Russell Crowe;Joaquin Phoenix;Connie Nielsen', 'A former Roman General sets out to exact vengeance against the corrupt emperor who murdered his family and sent him into slavery.'),
    (1, 'Saving Private Ryan', '1998-07-24', 'War', 'Steven Spielberg', 'Tom Hanks;Matt Damon;Tom Sizemore', 'Following the Normandy landings, a group of U.S. soldiers go behind enemy lines to retrieve a paratrooper whose brothers have been killed in action.'),
    (1, 'The Departed', '2006-10-06', 'Crime', 'Martin Scorsese', 'Leonardo DiCaprio;Matt Damon;Jack Nicholson', 'An undercover cop and a mole in the police attempt to identify each other while infiltrating an Irish gang in South Boston.'),
    (1, 'Whiplash', '2014-10-10', 'Drama', 'Damien Chazelle', 'Miles Teller;J.K. Simmons;Paul Reiser', 'A promising young drummer enrolls at a cut-throat music conservatory where his dreams of greatness are mentored by an instructor who will stop at nothing to realize a student''s potential.'),
    (1, 'Schindler''s List', '1993-12-15', 'Biography', 'Steven Spielberg', 'Liam Neeson;Ralph Fiennes;Ben Kingsley', 'In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.'),
    (1, 'Parasite', '2019-05-30', 'Thriller', 'Bong Joon Ho', 'Song Kang-ho;Lee Sun-kyun;Cho Yeo-jeong', 'Greed and class discrimination threaten the newly formed symbiotic relationship between the wealthy Park family and the destitute Kim clan.');
