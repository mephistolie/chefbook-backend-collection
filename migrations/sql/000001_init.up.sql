CREATE TYPE visibility as ENUM ('private', 'link', 'public');

CREATE TABLE collections
(
    collection_id UUID PRIMARY KEY NOT NULL,
    owner_id      UUID             NOT NULL,
    coauthors     UUID[]           NOT NULL DEFAULT '{}'::uuid[],
    name          VARCHAR(64)      NOT NULL,
    visibility    visibility       NOT NULL DEFAULT 'private',
    emoji         VARCHAR(30)               DEFAULT NULL
);

CREATE INDEX collections_owner_id_key ON collections (owner_id);

CREATE TABLE collections_users
(
    collection_id UUID REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL,
    user_id       UUID                                                          NOT NULL,
    UNIQUE (user_id, collection_id)
);

CREATE TABLE collections_recipes
(
    collection_id UUID REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL,
    recipe_id     UUID                                                          NOT NULL,
    UNIQUE (recipe_id, collection_id)
);

CREATE TABLE inbox
(
    message_id UUID PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);
