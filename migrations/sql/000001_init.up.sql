CREATE TABLE categories
(
    category_id uuid PRIMARY KEY NOT NULL,
    user_id     uuid             NOT NULL,
    name        text             NOT NULL,
    emoji       text DEFAULT NULL
);

CREATE INDEX categories_user_id_key ON categories (user_id);

CREATE TABLE inbox
(
    message_id uuid PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);
