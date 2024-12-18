

CREATE TABLE permission (
    id SERIAL PRIMARY KEY,
    permission character varying(100) NOT NULL
);

CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    role character varying NOT NULL
);

CREATE TABLE user2role (
    user_id integer,
    role_id integer,
    id SERIAL PRIMARY KEY,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES role(id)
);

CREATE TABLE role2permission (
    role_id integer,
    permission_id integer,
    id SERIAL PRIMARY KEY,
    allow_create boolean,
    allow_read boolean,
    allow_update boolean,
    allow_delete boolean,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES role(id),
    CONSTRAINT fk_permission FOREIGN KEY (permission_id) REFERENCES permission(id)
);

