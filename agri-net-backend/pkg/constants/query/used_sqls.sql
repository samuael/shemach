grant all on all tables in schema public to tamdes;
grant all on all sequences in schema public to tamdes;




-- getting the current unix second into big integer
select ROUND(extract(epoch from now()));