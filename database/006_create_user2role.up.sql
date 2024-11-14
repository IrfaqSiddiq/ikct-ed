

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

INSERT INTO public.permission (id,permission) VALUES
(1,'job'),
(2,'blog'),
(3,'sitemap'),
(4,'users'),
(5,'synthetic_job'),
(6,'bounty_company'),
(7,'permissionNrole'),
(8,'payout');

INSERT into public.role2permission 
(role_id,permission_id,allow_create,allow_read,allow_update,allow_delete) VALUES
(1,1,true,true,true,true),
(1,2,true,true,true,true),
(1,3,true,true,true,true),
(1,4,true,true,true,true),
(1,5,true,true,true,true),
(1,6,true,true,true,true),
(1,7,true,true,true,true),
(1,8,true,true,true,true);