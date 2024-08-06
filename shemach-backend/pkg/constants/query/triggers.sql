


-- create or replace function deleteUnconfirmedUsers() returns trigger as 
-- $$
--     declare
--         theid integer;
--     begin
--         if OLD.is_new_account then
--             delete from admin where email=OLD.new_email returning id into theid;
--             if found then
--                 return NULL;
--             end if; 
--             delete from infoadmin where email=OLD.new_email returning id into theid;
--             if found then
--                 return NULL;
--             end if;
--             RETURN NULL;
--         end if;
--     end;
-- $$ language plpgsql;

-- create trigger deletingUnconfirmedAccounts after delete on emailInConfirmation for each 
-- ROW EXECUTE PROCEDURE  deleteUnconfirmedUsers(); 