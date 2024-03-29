-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender AS ENUM ('M', 'F');

CREATE TABLE actors (
    id serial NOT NULL,
    name text NOT NULL,
    gender gender,
    birthdate date NOT NULL,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE actors;
DROP TYPE gender;
-- +goose StatementEnd
