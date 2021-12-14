USE go_database;

INSERT
    users (name, token) VALUE ('太郎', UUID());

INSERT INTO
    rarities (rarity, probability)
VALUES
    ("SSR", 5),
    ("SR", 10),
    ("R", 20),
    ("N", 40);

INSERT INTO
    characters (name, rarity_id)
VALUES
    ("Ken", 1),
    ("Tom", 2),
    ("Jaon", 2),
    ("Amy", 3),
    ("Jack", 3),
    ("Kebin", 3),
    ("Tommy", 4),
    ("Boby", 4),
    ("Kenny", 4),
    ("Ren", 4);

INSERT INTO
    user_character_possessions (user_id, character_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3);
