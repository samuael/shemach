ALTER TABLE student ADD CONSTRAINT unique_phone_number UNIQUE (phone);
-- REFERENCE https://www.postgresql.org/docs/7.4/sql-altertable.html
grant all on FUNCTION createSpecialCaseForStudent to tamdes;
grant all on student to tamdes;
grant all on special_case to tamdes;
grant all on  special_case_id_seq  to tamdes;
grant all on table category to tamdes;
alter table payin add column uchars char(2) not null;

grant all on payin_id_seq to tamdes;

insert into payin(amount,recieved_by, payed_by, created_at, roundid, status, uchars) values(1200.0,21,19,25,16,0,'xf');

grant all on payin to tamdes;
select checktheexsistanceofpayment(cast ('xf' as char(2)),19,16, 2014,6,29,17,36,10);
select createPayinTransaction( cast('bf'as char(2)),8,21,16,2014,7,20,7,24,6,2000);

-- for creating a product instance information.
select createProduct( cast ( 'Tef' as varchar) , cast ( 'Gojam' as varchar),2,cast ( 0.0 as decimal),1,cast (ROUND(extract(epoch from now())) as integer), cast ( ROUND(extract(epoch from now())) as integer));

-- update subscriber set subscriptions=array_remove(subscriptions, 2::smallint);
select  * from createStore( cast ('Assosa Market' as varchar) , cast(10 as integer) , cast (5 as integer),cast('04' as varchar),cast( '05' as varchar),cast('assosa'as varchar),cast ('benishangul' as varchar), cast ('assosa' as varchar), cast('' as varchar),'4534534.53453','45435345.54354');