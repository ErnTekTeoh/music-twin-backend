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

