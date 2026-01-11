CREATE SCHEMA music_twin_sg;
USE music_twin_sg;

CREATE TABLE users
(
    user_id               int NOT NULL AUTO_INCREMENT,
    display_name             varchar(100),
    email                   varchar(100),
    salt                    varchar(16),
    hash                    varchar(64),
    profile_image_url       varchar(1024),
    created_at              datetime,
    updated_at              datetime,
    deleted_at              datetime,
    bio varchar(1024),
    gender int,
    location varchar(128),
    user_referral_code varchar(10),
    joining_referral_code varchar(10),
    favourite_artist1_id int,
    favourite_artist1_name varchar(128),
    favourite_artist2_id int,
    favourite_artist2_name varchar(128),
    favourite_artist3_id int,
    favourite_artist3_name varchar(128),
    favourite_genre_name1 varchar(128),
    favourite_genre_name2 varchar(128),
    favourite_genre_name3 varchar(128),
    INDEX(email),
    UNIQUE(user_referral_code),
    PRIMARY KEY (user_id)
);

CREATE TABLE user_top_picks (
    user_top_pick_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    type ENUM('artist', 'song'),
    ranking INT,
    apple_music_artist_image_url VARCHAR(512),
    apple_music_artist_name VARCHAR(128),
    apple_music_external_id VARCHAR(128),
    apple_music_song_name VARCHAR(128),
    apple_music_song_image_url VARCHAR(512),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    INDEX(user_id)
);


CREATE TABLE song_suggestion_cards (
   song_suggestion_card_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   song_title VARCHAR(128),
   artist_name VARCHAR(128),
   image_url VARCHAR(512),
   card_header VARCHAR(128),
   card_subheader VARCHAR(256),
   recommend_reason VARCHAR(256),
   created_at DATETIME,
   updated_at DATETIME
);

CREATE TABLE song_poll_cards (
     song_poll_card_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
     card_header VARCHAR(128),
     card_subheader VARCHAR(256),
     card_disclaimer VARCHAR(256),
     created_at DATETIME,
     updated_at DATETIME
);

CREATE TABLE song_poll_card_options (
    song_poll_card_option_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    song_poll_card_id INT NOT NULL,
    song_title VARCHAR(128),
    artist_name VARCHAR(128),
    image_url VARCHAR(512),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (song_poll_card_id) REFERENCES song_poll_cards(song_poll_card_id)
);

ALTER TABLE users ADD COLUMN instagram_handle varchar(100);
ALTER TABLE users ADD COLUMN whatsapp_handle varchar(100);
ALTER TABLE users ADD COLUMN telegram_handle varchar(100);
ALTER TABLE users ADD COLUMN alternate_email varchar(256);

ALTER TABLE user_top_picks ADD COLUMN image_url varchar(512);