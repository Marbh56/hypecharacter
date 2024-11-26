-- +goose Up

CREATE TABLE characters (
    character_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    character_name TEXT NOT NULL,
    str INT NOT NULL,
    dex INT NOT NULL,
    con INT NOT NULL,
    int INT NOT NULL,
    wis INT NOT NULL,
    cha INT NOT NULL,
    save INT NOT NULL,
    ac INT NOT NULL,
    mv INT NOT NULL,
    CONSTRAINT fk_character_user FOREIGN KEY (user_id) REFERENCES users(id)
);


-- +goose Down

DROP TABLE characters;
