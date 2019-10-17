CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id         UUID                        DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT users_pk PRIMARY KEY,
    login      VARCHAR(50)                                            NOT NULL,
    first_name VARCHAR(50)                                            NOT NULL,
    last_name  VARCHAR(50)                                            NOT NULL,
    password   VARCHAR(100)                                           NOT NULL,

    deleted    BOOLEAN                     DEFAULT FALSE              NOT NULL,

    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS users_id_uindex
    ON users (id);
CREATE UNIQUE INDEX IF NOT EXISTS users_login_uindex
    ON users (login);

CREATE TABLE IF NOT EXISTS roles
(
    id          UUID                        DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT roles_pk PRIMARY KEY,
    title       VARCHAR(50)                                            NOT NULL,
    description VARCHAR(50)                                            NOT NULL,

    super       BOOLEAN                     DEFAULT FALSE              NOT NULL,

    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS roles_id_uindex
    ON roles (id);
CREATE UNIQUE INDEX IF NOT EXISTS roles_title_uindex
    ON roles (title);

CREATE TABLE IF NOT EXISTS actions
(
    title       VARCHAR(100)                              NOT NULL
        CONSTRAINT actions_pk PRIMARY KEY,
    description VARCHAR(100)                              NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS actions_title_uindex
    ON actions (title);

CREATE TABLE IF NOT EXISTS users_roles
(
    user_id    UUID                                      NOT NULL
        CONSTRAINT users_roles_user_id_fk
            REFERENCES users
            ON UPDATE CASCADE ON DELETE CASCADE,
    role_id    UUID                                      NOT NULL
        CONSTRAINT users_roles_role_id_fk
            REFERENCES roles
            ON UPDATE CASCADE ON DELETE CASCADE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS users_roles_user_id_role_id_uindex
    ON users_roles (user_id, role_id);

CREATE TABLE IF NOT EXISTS actions_roles
(
    action_title VARCHAR(100)                              NOT NULL
        CONSTRAINT actions_roles_action_title_fk
            REFERENCES actions
            ON UPDATE CASCADE ON DELETE CASCADE,
    role_id      UUID                                      NOT NULL
        CONSTRAINT actions_roles_role_id_fk
            REFERENCES roles
            ON UPDATE CASCADE ON DELETE CASCADE,
    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS actions_roles_action_title_role_id_uindex
    ON actions_roles (action_title, role_id);

CREATE TABLE IF NOT EXISTS sessions
(
    id                UUID                        DEFAULT uuid_generate_v4() NOT NULL
        CONSTRAINT sessions_pk PRIMARY KEY,
    user_id           UUID                                                   NOT NULL
        CONSTRAINT sessions_user_id_fk
            REFERENCES users
            ON UPDATE CASCADE ON DELETE CASCADE,
    access_token      VARCHAR(100)                                           NOT NULL,
    access_token_ttl  INT                                                    NOT NULL,
    refresh_token     VARCHAR(100)                                           NOT NULL,
    refresh_token_ttl INT                                                    NOT NULL,
    created_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL,
    updated_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now()              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS sessions_id_uindex
    ON sessions (id);
CREATE UNIQUE INDEX IF NOT EXISTS access_token_uindex
    ON sessions (access_token);
CREATE UNIQUE INDEX IF NOT EXISTS refresh_token_uindex
    ON sessions (refresh_token);
