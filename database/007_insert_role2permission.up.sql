INSERT INTO permission (id,permission) VALUES
(1,'student'),
(2,'school'),
(3,'user');

INSERT INTO role(id, role)VALUES
(1, 'super-admin'),
(2, 'admin');

INSERT INTO user2role (role_id,user_id)VALUES
(1,1),
(2,2);

INSERT into public.role2permission 
(role_id,permission_id,allow_create,allow_read,allow_update,allow_delete) VALUES
(1,1,true,true,true,true),
(1,2,true,true,true,true),
(1,3,true,true,true,true),
(2,1,true,true,true,true),
(2,2,true,true,true,true),
(2,3,false,true,false,false);