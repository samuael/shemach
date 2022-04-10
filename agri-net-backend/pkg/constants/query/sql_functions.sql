create or replace function checkTheExistanceOfSubscriberByPhone( iphone varchar )  returns integer as 
$$
    declare 
        subscriberid integer;
    begin 
       
        select id into subscriberid from tempo_subscriber where phone=iphone;
        if found then 
            return 2;
        end if;
        
        select id into subscriberid from tempo_subscribers_login  where phone=iphone;
        if found then
            return 4;
        end if;
        select id into subscriberid from subscriber where phone=iphone;
        if found then
            return 1;
        end if;
        return 3;
    end;
$$ language plpgsql;


create or replace function selectTempoSubscriberWithPhoneAndUpdatedTrials(iphone varchar) returns tempo_subscriber as 
$$
    declare 
        sampleTempo tempo_subscriber;
    begin
        select * into sampleTempo from tempo_subscriber where phone=iphone;
        if not found then
            return null;
        else
            if sampleTempo.trials >=50 then
                delete from tempo_subscriber where id=sampleTempo.id;
                return null;
            else
                update tempo_subscriber as ts set trials= ts.trials +1 where id=sampleTempo.id;
                sampleTempo.trials= sampleTempo.trials+1;
                return sampleTempo;
            end if;
        end if;
        return NULL;
    end;
$$ language plpgsql;



create or replace function selectTempoLoginSubscriberWithPhoneAndUpdatedTrials(iphone varchar()) returns tempo_subscriber as 
$$
    declare 
        sampleTempo tempo_subscriber_login;
    begin
        select id,fullname,Phone,lang,role,confirmation,unix ,trials into sampleTempo from tempo_subscriber where phone=iphone;
        if not found then
            return null;
        else
            if sampleTempo.trials >=50 then
                delete from tempo_subscriber where id=sampleTempo.id;
                return null;
            else
                update tempo_subscriber as ts set ts.trials=trials +1 where id=sampleTempo.id;
                return sampleTempo; 
            end if;
        end if;
        return NULL;
    end;
$$ language plpgsql;

-- No Trigger is allowed with a select statement.

-- CREATE TRIGGER incrementRegistrationTempoRequestTrial 
-- AFTER SELECT ON tempo_subscriber FOR EACH 
-- ROW EXECUTE PROCEDURE incrementOrDeleteTempoSubscriberInfo();