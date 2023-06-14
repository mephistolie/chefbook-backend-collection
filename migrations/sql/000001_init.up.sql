CREATE TABLE categories
(
    category_id uuid PRIMARY KEY NOT NULL UNIQUE,
    user_id     uuid PRIMARY KEY NOT NULL,
    name        text             NOT NULL,
    emoji       text DEFAULT NULL,
);

CREATE TABLE inbox
(
    message_id uuid PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);
