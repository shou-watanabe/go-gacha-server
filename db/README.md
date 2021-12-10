## コンテナの中に入る
docker container exec -it techtrain-mission_db_1 bash

## mysqlに接続
mysql -u root -p

## パスワード
root

## 使用するDBの選択
use go_database;

## キャラクター情報取得のSQL(id, name, rarity, probability)
select characters.id, characters.name, rarities.rarity, rarities.probability from characters, rarities where characters.rarity_id = rarities.i
d;
