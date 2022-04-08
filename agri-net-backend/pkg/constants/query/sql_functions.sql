create or replace function checkTheExistanceOfSubscriberByPhone( iphone varchar )  returns integer as 
$$
    declare 
        subscriberid integer;
    begin 
        select id into subscriberid from subscriber where phone=iphone;
        if found then
            return 1;
        end if; 
        select id into subscriberid from tempo_subscriber where phone=iphone;
        if found then 
            return 2;
        end if;
        return 3;
    end;
$$ language plpgsql;