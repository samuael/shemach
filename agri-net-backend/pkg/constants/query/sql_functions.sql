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


-- This function returns 
-- 0 if the user does exist in the superadmins 
-- 1 if found in the user and 
-- 2 if no user instance was found with this email address.
create or replace function checkTheExistanceOfUser(iemail varchar) returns integer as 
$$
    declare 
        theuserid integer;
    begin
        select id into theuserid from superadmin where email = iemail;
        if found then
            return 0;
        end if;
        select id into theuserid from users where email=iemail;
        if found then
            return 1;
        end if;
        return 2;
    end;
$$ language plpgsql;



create or replace function createProduct(
    iname varchar , 
    iproduction_area varchar, 
    iunit_id integer,
    icurrent_price decimal,
    icreated_by integer, 
    icreated_at integer, 
    ilast_updated_time integer) returns integer as 
$$
    declare
        theproduct product;
        theid integer;
        theadmin superadmin;
    begin
        select * from product into theproduct where name=iname and production_area=iproduction_area  and unit_id=iunit_id;
        if found then
            return -1;
        end if;
        insert into product(
            name,
            production_area,
            unit_id,
            current_price,
            created_by,
            created_at,
            last_updated_time)values(
                iname,
                iproduction_area,
                iunit_id,
                icurrent_price,
                icreated_by,
                icreated_at,
                ilast_updated_time) returning id into theid;
            if not found then
                return -2;
            end if;

            select * into theadmin from superadmin where id=icreated_by;
            if not found then
                return -3;
            end if;

            return theid;
    end;
$$ language plpgsql;

create or replace function checkTheExistanceOfProduct( iname varchar(200), iproduction_area varchar(200),iunit_id integer) returns boolean as 
$$
    declare 
        productid integer;
    begin
        select id into productid from product where name=iname and production_area=iproduction_area and unit_id=iunit_id;
        if found then
            return true;
        end if;
        return false; 
    end;
$$ language plpgsql;



create or replace function createNewSubscription( isubscriberid integer, productid smallint ) returns integer as 
$$
    declare
        subscriberid integer;
        statustrack integer;
    begin
        select id into subscriberid from subscriber where id=isubscriberid;
        if not found then
            return -1;
        end if;
        select id into statustrack from product where id =productid;
        if not found then
            return -2;
        end if;
        with rows as ( update subscriber set subscriptions = array_append(subscriptions, productid::smallint)
         where (not productid::smallint = any( subscriptions )) and id = isubscriberid returning 1)
        select count(*) into statustrack from rows;
        if not found or statustrack = 0 then
            return -3;
        end if;
        return 0;
    end;
$$ language plpgsql;


create or replace function UnSubscribeToProduct(isubscriberid integer,productid smallint) returns integer as 
$$
    declare
        subscriberid integer;
        statustrack integer;
    begin
        select id into subscriberid from subscriber where id=isubscriberid;
        if not found then
            return -1;
        end if;
        select id into statustrack from product where id =productid;
        if not found then
            return -2;
        end if;
        
        with rows as ( update subscriber set subscriptions = array_remove(subscriptions, productid::smallint)
         where (productid::smallint = any( subscriptions )) and id = isubscriberid returning 1)
        
        select count(*) into statustrack from rows;
        
        if (not found or statustrack = 0) then
            return -3;
        end if;

        return 0;
    end;
$$ language plpgsql;




create or replace function getTheRoleOfUserByIdOrEmail( userid integer , iemail varchar)  returns integer as 
$$
    declare
        theid varchar;
    begin
        select id into theid from superadmin where id=userid or email=iemail;
        if found then
            return 1;
        end if;
        select id into theid from infoadmin where id=userid or email=iemail;
        if found then
            return 2;
        end if;
        return 0;
    end;
$$ language plpgsql;


create or replace function updateProductPrice( productid integer , new_price decimal ) returns integer as 
$$
    declare
        cost integer;
        updated_count integer;
    begin
        select current_price into cost from product where id=productid;
        if not found then 
            return -1;
        end if;
        if cost = new_price then 
            return -2;
        end if;
        with rowss as (
            update product set current_price=new_price , last_updated_time=ROUND(extract( epoch  from now())) where id=productid returning id
        )
        select COUNT(*) into updated_count from rowss;
        if not found then
            return -3;
        end if;
        return 0;
    end;
$$ language plpgsql;