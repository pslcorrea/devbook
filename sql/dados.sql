INSERT INTO usuarios (nome, nick, email, senha)
VALUES
('Paulo','pslcorrea','paulo@gmail.com','$2a$10$Rav1zsQjubcyfngGfOG5RucA3tFIUzVSsVuXgTQoBSSHzqbtCkfju'),
('Ana','anab','ana@gmail.com','$2a$10$Rav1zsQjubcyfngGfOG5RucA3tFIUzVSsVuXgTQoBSSHzqbtCkfju'),
('Iara','iarmsc','iarao@gmail.com','$2a$10$Rav1zsQjubcyfngGfOG5RucA3tFIUzVSsVuXgTQoBSSHzqbtCkfju'),
('Maria','maria','maria@gmail.com','$2a$10$Rav1zsQjubcyfngGfOG5RucA3tFIUzVSsVuXgTQoBSSHzqbtCkfju');

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1,2),
(1,3),
(1,4),
(2,1),
(2,3),
(3,4),
(3,1);