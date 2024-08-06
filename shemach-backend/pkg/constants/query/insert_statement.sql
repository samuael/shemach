insert into users(id, firstname, lastname, phone, email,created_at,password,imageurl,lang) values(2, 'Mesfin','WoldeMariam','+251912131415', 'mesfinwolde@gmail.com', 1649737198 , '$2a$10$TQW6xnGA6xntDzWU21fGg.7.ad8vgfWz7OYvQETIRD/BPrpuvUKvG' ,'','amh');
insert into users(id, firstname, lastname, phone, email,created_at,password,imageurl,lang) values(6, 'Dawit','Teshager','', 'samuaeladnew@gmail.com', 1651306088, '$2a$10$lkeJioro3iuqWnEykSCNz.dvX3DKXz9G0GG7HQlE.P9PubTJrQoCG','' ,'amh');
insert into users(id, firstname, lastname, phone, email,created_at,password,imageurl,lang) values(3, 'Dawit','Teshager','+251992939394', 'abebeteka@gmail.com'   , 1650279386, '$2a$10$OjghnRcNSoU7V4d7m50xp.EdNmpi.qRRRn1Uwe4Xr6cQhkALuSsEC','images/profile/NpKUf.jpg','amh');







insert into messages(id , targets, lang , data,created_by, created_at) values
  (4, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,4, 1650126183),
  (5, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,4, 1650126186),
  (6, ARRAY [-1] , 'amh' , 'Sami tatatfiw'               ,4, 1650126238),
  (7, ARRAY [-1] , 'eng' , 'Sami tatatfiw'               ,4, 1650126244),
  (8, ARRAY [-1] , 'amh' , 'Sami tatatfiw'               ,4, 1650126257),
  (9, ARRAY [-1] , 'oro' , 'Sami tatatfiw'               ,4, 1650126262),
 (10, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131414),
 (11, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131482),
 (12, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131515),
 (13, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131559),
 (14, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131587),
 (15, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131593),
 (16, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131618),
 (17, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131634),
 (18, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650131662),
 (19, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650290359),
 (20, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650290360),
 (21, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,2, 1650290362),
 (22, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,3, 1650790259),
 (23, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,3, 1650790567),
 (24, ARRAY [-1] , 'all' , 'Welcome to shemach systems' ,3, 1650790664);




 insert into superadmin (id,firstname,lastname,phone,email,created_at,password,registered_admins,registered_products, imageurl, lang)
 values
 (2,'Mesfin','WoldeMariam' ,'+251912131415','mesfinwolde@gmail.com',1649737198,'$2a$10$TQW6xnGA6xntDzWU21fGg.7.ad8vgfWz7OYvQETIRD/BPrpuvUKvG',0,0,'' ,'amh')




insert into product(id ,   name   , production_area , unit_id , current_price , created_by , created_at , last_updated_time )
values(2 , 'Teff Aja' , 'Gojam',1,4000 ,2 , 1650009962 ,1650009962),
(1 , 'Tef'      , 'Gojam',2,6780.567,1 , 1650002018 ,1650797454)




insert into infoadmin(firstname,lastname,phone ,email,created_at,password,imageurl,messages_count,created_by,lang) 
values
('Dawit','Teshager','','samuaeladnew@gmail.com', 1651306088 ,'$2a$10$lkeJioro3iuqWnEykSCNz.dvX3DKXz9G0GG7HQlE.P9PubTJrQoCG','',0,0,'amh'),
('Dawit','Teshager','','abebeteka@gmail.com'   , 1650279386 ,'$2a$10$OjghnRcNSoU7V4d7m50xp.EdNmpi.qRRRn1Uwe4Xr6cQhkALuSsEC','images/profile/NpKUf.jpg' ,0 ,0 , 'amh');
