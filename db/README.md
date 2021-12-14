## コンテナの中に入る
docker container exec -it techtrain-mission_db_1 bash

## mysqlに接続
mysql -u root -p

## パスワード
root

## 使用するDBの選択
use go_database;

## キャラクター情報取得のSQL(id, name, rarity, probability)
select characters.id, characters.name, rarities.rarity, rarities.probability from characters, rarities where characters.rarity_id = rarities.id;

## inner join版
SELECT
 characters.id, characters.name, rarities.rarity, rarities.probability
FROM characters
INNER JOIN rarities
ON characters.rarity_id = rarities.id;

## useridと一致するキャラクターのidを取得するSQL
select * from users, user_character_possessions where users.id = user_character_possessions.user_id;

##
select users.id, users.name, characters.name from users, user_character_possessions, characters where users.id = user_character_possessions.user_id and user_character_possessions.character_id = characters.id;

## ユーザーが所有しているキャラを表示
select characters.id, characters.name, rarities.rarity, rarities.probability from user_character_possessions, characters, rarities where user_character_possessions.user_id = 1 and user_character_possessions.character_id = characters.id and characters.rarity_id = rarities.id;

## inner join版
SELECT characters.id, characters.name, rarities.rarity, rarities.probability
FROM user_character_possessions
INNER JOIN characters
ON user_character_possessions.character_id = characters.id
INNER JOIN rarities
ON characters.rarity_id = rarities.id
WHERE user_character_possessions.user_id = 1;
