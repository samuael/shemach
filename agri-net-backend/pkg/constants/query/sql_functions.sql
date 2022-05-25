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


-- Create subscription and Unsubscription for 


create or replace function MerchantCreateNewSubscription( imerchantid integer, productid smallint ) returns integer as 
$$
    declare
        merchantid integer;
        statustrack integer;
    begin
        select id into merchantid from merchant where id=imerchantid;
        if not found then
            return -1;
        end if;
        select id into statustrack from product where id =productid;
        if not found then
            return -2;
        end if;
        with rows as ( update merchant set subscriptions = array_append(subscriptions, productid::smallint)
         where (not productid::smallint = any( subscriptions )) and id = imerchantid returning 1)
        select count(*) into statustrack from rows;
        if not found or statustrack = 0 then
            return -3;
        end if;
        return 0;
    end;
$$ language plpgsql;


create or replace function MerchantUnSubscribeToProduct(imerchantid integer,productid smallint) returns integer as 
$$
    declare
        merchantid integer;
        statustrack integer;
    begin
        select id into merchantid from merchant where id=imerchantid;
        if not found then
            return -1;
        end if;
        select id into statustrack from product where id =productid;
        if not found then
            return -2;
        end if;
        
        with rows as ( update merchant set subscriptions = array_remove(subscriptions, productid::smallint)
         where (productid::smallint = any( subscriptions )) and id = imerchantid returning 1)
        
        select count(*) into statustrack from rows;
        if (not found or statustrack = 0) then
            return -3;
        end if;
        return 0;
    end;
$$ language plpgsql;

-- Merchants Subscription and unsubscription method completed

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
        select id into theid from admin where id=userid or email=iemail;
        if found then
            return 3;
        end if;
        select id into theid from merchant where id=userid or email=iemail;
        if found then
            return 4;
        end if;
        select id into theid from agent where id=userid or email=iemail;
        if found then
            return 5;
        end if;

        return 0;
    end;
$$ language plpgsql;


-- getTheRoleOfUserByPhone
create or replace function getTheRoleOfUserByPhone( iphone varchar)  returns integer as 
$$
    declare
        theid varchar;
    begin
        select id into theid from superadmin where phone=iphone;
        if found then
            return 1;
        end if;
        select id into theid from infoadmin where phone=iphone;
        if found then
            return 2;
        end if;
        select id into theid from admin where phone=iphone;
        if found then
            return 3;
        end if;
        select id into theid from merchant where phone=iphone;
        if found then
            return 4;
        end if;
        select id into theid from agent where phone=iphone;
        if found then
            return 5;
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



create or replace function deleteinfoadminById(did integer) returns integer as 
$$
    declare
        statusCode integer;
        theid integer;
    begin
        select id into theid from infoadmin where id=did;
        if not found then
            return -1;
        end if;

        with rows as (
            delete from infoadmin where id=did returning id
        )
        select count(*) into theid from rows;
        if not found then
            return -2;
        else 
            return 0;
        end if;
    end;
$$ language plpgsql;

create or replace function deleteadminById(did integer) returns integer as 
$$
    declare
        statusCode integer;
        theid integer;
    begin
        select id into theid from admin where id=did;
        if not found then
            return -1;
        end if;

        with rows as (
            delete from admin where id=did returning id
        )
        select count(*) into theid from rows;
        if not found then
            return -2;
        else 
            return 0;
        end if;
    end;
$$ language plpgsql;


create or replace function insertEmailInConfirmation(iuserid integer, inewemail varchar, inewaccount boolean, ioldemail varchar) returns integer as 
$$
    declare
        duserid integer;
        theid integer;
    begin
        select id into duserid from users where id=iuserid;
        if not found then
            return -1;
        end if;
        insert into emailInConfirmation(userid, new_email,is_new_account, old_email)
            values(iuserid , inewemail , inewaccount, ioldemail ) returning id into theid;
        if not found then
            return -2;
        end if;
        return theid;
    end;
$$ language plpgsql;


create or replace function commitEmailChange(  confirmNewEmail varchar)   returns integer as 
$$
    declare
        ec_id integer;
        counts integer;
        duserid integer;
    begin
        select id, userid into ec_id,duserid  from emailInConfirmation where new_email=confirmNewEmail;
        if not found then
            return -1;
        end if;
        -- 
        select id into duserid from users where id=duserid;
        if not found then
            return -2;
        end if;
        -- 
        with rows as (
            update users set email=confirmNewEmail where id=duserid returning id
        )
        select COUNT(*) into counts from rows;
        if counts = 0 then
            return -3;
        end if;
        raise notice 'the id is %',ec_id; 
        delete from emailInConfirmation where id=ec_id;
        return 0;
    end;
$$ language plpgsql;


-- This method returns two values.
create or replace function  createAdminInstance(ifirstname varchar,ilastname varchar,iphone varchar,iemail varchar,ipassword text,ilang char(3),icreated_by integer, 
ikebele varchar,iworeda varchar,icity varchar,iregion varchar, izone varchar,
iunique_name varchar ,ilatitude varchar,ilongitude varchar) returns integer as 
$$
    declare
        addressid integer;
        adminid integer;
        thesuperadminid integer;
        rec RECORD;
    begin
        select id into thesuperadminid from superadmin where id=icreated_by;
        if not found then
            return -1;
        end if;
        select address_id into addressid from address where kebele=ikebele and woreda=iworeda and  city=icity and  region = iregion and zone=izone and  unique_name=iunique_name and latitude=ilatitude and longitude=ilongitude;
        if not found and ikebele <> '' or iworeda<>'' or icity <>'' or iregion<>'' or izone<>'' or iunique_name<>'' or ilatitude <> '' or ilongitude <>'' then
                    insert into address(kebele,woreda,city,region,zone,unique_name,latitude,longitude) 
            values(ikebele,iworeda ,icity ,iregion,izone,iunique_name ,ilatitude,ilongitude) returning address_id into addressid;
            if not found then
                return -2;
            end if;
        end if;  

        insert into admin( firstname,lastname,phone,email,password,lang,created_by, address_id)
        values(ifirstname ,ilastname ,iphone ,iemail,ipassword ,ilang,icreated_by,addressid) returning id into adminid;
        if not found then
            rollback;
            return -3;
        end if;
        return adminid;
    end;
$$
language plpgsql;


-- createAgent
create or replace function createAgent( ifirstname varchar,ilastname varchar,iphone varchar,iemail varchar,ipassword text,ilang char(3), iregistered_by integer,
ikebele varchar,iworeda varchar,icity varchar,iregion varchar, izone varchar, iunique_name varchar ,ilatitude varchar,ilongitude varchar ) returns integer as 
$$
    declare
        addressid integer;
        agentid integer;
        theadminid integer;
    begin
        select id into theadminid from admin where id=iregistered_by;
        if not found then
            return -1;
        end if;
        select address_id into addressid from address where kebele=ikebele and woreda=iworeda and  city=icity and  region = iregion and zone=izone and  unique_name=iunique_name and latitude=ilatitude and longitude=ilongitude;
        if not found and ikebele <> '' or iworeda<>'' or icity <>'' or iregion<>'' or izone<>'' or iunique_name<>'' or ilatitude <> '' or ilongitude <>'' then
                    insert into address(kebele,woreda,city,region,zone,unique_name,latitude,longitude) 
            values(ikebele,iworeda ,icity ,iregion,izone,iunique_name ,ilatitude,ilongitude) returning address_id into addressid;
            if not found then
                return -2;
            end if;
        end if;  

        insert into agent( firstname,lastname,phone,email,password,lang,registered_by, field_address_ref)
        values(ifirstname ,ilastname ,iphone ,iemail,ipassword ,ilang,iregistered_by,addressid) returning id into agentid;
        if not found then
            rollback;
            return -3;
        end if;
        return agentid;
    end;
$$
language plpgsql;




-- createMerchant
create or replace function createMerchant( ifirstname varchar,ilastname varchar,iphone varchar,iemail varchar,ipassword text,ilang char(3), iregistered_by integer,
ikebele varchar,iworeda varchar,icity varchar,iregion varchar, izone varchar, iunique_name varchar ,ilatitude varchar,ilongitude varchar ) returns integer as 
$$
    declare
        merchantid integer;
        addressid integer;
        theadminid integer;
    begin
        select id into theadminid from admin where id=iregistered_by;
        if not found then
            return -1;
        end if;
        select address_id into addressid from address where kebele=ikebele and woreda=iworeda and  city=icity and  region = iregion and zone=izone and  unique_name=iunique_name and latitude=ilatitude and longitude=ilongitude;
        if not found and ikebele <> '' or iworeda<>'' or icity <>'' or iregion<>'' or izone<>'' or iunique_name<>'' or ilatitude <> '' or ilongitude <>'' then
                    insert into address(kebele,woreda,city,region,zone,unique_name,latitude,longitude) 
            values(ikebele,iworeda ,icity ,iregion,izone,iunique_name ,ilatitude,ilongitude) returning address_id into addressid;
            if not found then
                return -2;
            end if;
        end if;  

        insert into merchant( firstname,lastname,phone,email,password,lang,registerd_by, address_ref)
        values(ifirstname ,ilastname ,iphone ,iemail,ipassword ,ilang,iregistered_by,addressid) returning id into merchantid;
        if not found then
            rollback;
            return -3;
        end if;
        return merchantid;
    end;
$$
language plpgsql;


-- get all temporary commodity exchange participants and increment the trial by one every second.
create or replace function getTempoCXPByPhone( iphone varchar ) returns tempo_cxp as 
$$
    declare
        tempo tempo_cxp;
    begin
        select * into tempo from tempo_cxp where phone=iphone;
        if not found then
            return null;
        end if;
        if tempo.trials >=51 then
            delete from tempo_cxp where phone=iphone;
            return null;
        end if;
        update tempo_cxp as t set trials=t.trials + 1 where phone=iphone;
        return tempo;
    end; 
$$ language plpgsql;



create or replace function deleteExpredCXPAccount(  iphones varchar[] ) returns integer as 
$$
    declare
        theid integer;
        vari varchar;
        cunt integer;
    begin
        cunt :=0;
        foreach vari in  array iphones loop
            with rows as (
                delete from merchant where phone=vari returning id
            )
            select count(*) into theid from rows;
            if found and theid>0 then
                delete from tempo_cxp where phone= vari;
                cunt = cunt+theid;
                continue;
            end if;

            with rows as (
                delete from agent where phone=vari returning id
            ) 
            select count(*) into theid from rows;
            if found and theid >0 then
                delete from tempo_cxp where phone= vari;
                cunt = cunt+theid;
            end if;
        end loop;
        return cunt;
    end;
$$ language plpgsql;



create or replace function deleteUnconfirmedAdmins(  ids integer[] ) returns integer as 
$$
    declare
        idd integer;
        theid integer;
        cunt integer;
    begin
        cunt := 0;
        foreach idd in array ids loop
            delete from admin where id=idd returning id into theid;
            if found then
                cunt = cunt+theid;
                continue;
            end if; 
            delete from infoadmin where id=idd returning id into theid;
            if found then
                cunt = cunt+theid;
            end if;
            RETURN cunt;
        end loop;
    end;
$$ language plpgsql;




-- dictionary



create or replace function createDictionary( ilang char(3), itext text , itranslation text )  returns integer as 
$$
    declare 
        theid integer;
    begin
        select id into theid from dictionary where sentence=itext and lang=ilang and translation=itranslation;
        if found and theid >0 then
            return theid;
        end if;

        insert into dictionary(lang, sentence,translation) values(ilang, itext, itranslation) returning id into theid;
        if found then 
            return theid;
        end if;
        return -1;
    end;
$$ language plpgsql;



create or replace function createStore(  name varchar(100) , iowner_id integer , icreated_by integer,ikebele varchar(100),iworeda varchar(100),icity varchar(100),iregion varchar(100), izone varchar(100), iunique_name varchar(100) ,ilatitude varchar(40),ilongitude varchar(40)) returns integer as 
$$
    declare 
        addressid integer;
        theadminid integer;
        merchantid integer;
        thestoreid integer;
    begin
        select id into theadminid from admin where id=icreated_by;
        if not found then
            return -1;
        end if;
        select address_id into addressid from address where kebele=ikebele and woreda=iworeda and  city=icity and  region = iregion and zone=izone and  unique_name=iunique_name and latitude=ilatitude and longitude=ilongitude;
        if not found and ikebele <> '' or iworeda<>'' or icity <>'' or iregion<>'' or izone<>'' or iunique_name<>'' or ilatitude <> '' or ilongitude <>'' then
                    insert into address(kebele,woreda,city,region,zone,unique_name,latitude,longitude) 
            values(ikebele,iworeda ,icity ,iregion,izone,iunique_name ,ilatitude,ilongitude) returning address_id into addressid;
            if not found then
                return -2;
            end if;
        end if;
        select id into merchantid from merchant where id=iowner_id;
        if not found then
            return -3;
        end if; 

        insert into store (
            owner_id,address_id,store_name,created_by
        ) values(iowner_id,addressid,name,icreated_by) returning store_id into thestoreid;
        if not found then 
            return -4;
        end if;
        return thestoreid;
    end;
$$ language plpgsql;



-- crop_id integer,
create or replace function createProductPost( itype_id integer,idescription varchar(500),
        inegotiable boolean,iremaining_quantity integer,iselling_price decimal,
        iaddress_id integer,istore_id integer,iagent_id integer,istore_owned boolean) returns integer as
$$
    declare
        addressid integer;
        cropid integer;
        userid integer;
    begin
        select address_id into addressid from address where address_id = iaddress_id;
        if not found then
            return -1;
        end if;
        if istore_owned then
            select store_id into userid from store where store_id=istore_id;
            if not found then 
                return -2;
            end if;
        else 
            select id into userid from agent where id=iagent_id;
            if not found then
                return -3;
            end if;
        end if;
        insert into crop(
            type_id,description,negotiable,remaining_quantity,selling_price,address_id,store_id,agent_id,store_owned
        ) values (
            itype_id,idescription,inegotiable,iremaining_quantity,iselling_price,iaddress_id,istore_id,iagent_id,istore_owned
        ) returning crop_id into cropid;
        if not found then
            return -4;
        end if;
        return cropid;
    end;
$$ language plpgsql;




create or replace function declineTransaction(transactionid integer,userid integer) returns integer as 
$$
    declare 
        dstate integer;
        dcount integer;
    begin
        select state from transaction where requester_id=userid or seller_id=userid into dstate;
        if not found then
            return -1;
        end if;
        if dstate = 12 then
            return -2;
        end if;
        with rows as (
            delete from transaction where transaction_id=transactionid returning transaction_id
        )
        select count(*) into dcount from rows;
        if found and dcount >0 then
            return 0;
        end if;
        return -3;
    end;
$$ language plpgsql;



create or replace function acceptTransactionAmendmentRequest(merchantid integer, requestid integer) returns integer as 
$$
    declare
        transactionreq transaction_changes;
        transactionid integer;
        buyerid integer;
    begin
        select * from transaction_changes into transactionreq where transaction_changes_id=requestid;
        if not found then
            return -4;
        end if;
        select id into merchantid from merchant where id=merchantid;
        if not found then
            return -1;
        end if;
        select requester_id into buyerid from transaction where transaction_id=transactionreq.transaction_id;
        if not found then
            return -2;
        end if;
        if buyerid <> merchantid then
            return -3;
        end if;
        update transaction set price=transactionreq.price, quantity=transactionreq.qty, state=3 where transaction_id=transactionreq.transaction_id returning transaction_id into transactionid;
        if not found then 
            return -5;
        end if;
        delete from transaction_changes where transaction_changes_id=transactionreq.transaction_changes_id;
        update transaction set state=3 where transaction_id=transactionreq.transaction_id;
        return 0;
    end;
$$ language plpgsql;

create or replace function amendTransaction(buyerid integer, transactionrequestid integer,dprice decimal, dquantity integer , descs varchar(500)) returns integer as 
$$
    declare
        transactionid integer;
        transactionreq transaction_changes;
        buyerid integer;
        change_description boolean;
    begin
        select * from transaction_changes into transactionreq where transaction_changes_id=transactionrequestid;
        if not found then
            return -1;
        end if;
        select requester_id into buyerid from transaction where transaction_id=transactionreq.transaction_id;
        if not found then
            return -2;
        end if;
        if descs <> '' and transactionreq.description <> descs then
            change_description= true;
        else 
            change_description= false;
        end if;
        if change_description then
            update transaction set price=dprice, quantity=dquantity, description=descs,state=3 
            where transaction_id=transactionreq.transaction_id returning transaction_id into transactionid;
        else
            update transaction set price=dprice, quantity=dquantity, state=3 
            where transaction_id=transactionreq.transaction_id returning transaction_id into transactionid;
        end if;

        if not found then 
            return -3;
        end if;
        delete from transaction_changes where transaction_changes_id=transactionreq.transaction_changes_id;
        update transaction set state=3 where transaction_id=transactionreq.transaction_id;
        return 0;
    end;
$$ language plpgsql;

-- createGuaranteeRequest
create or replace function createKebdRequest( sellerid integer , iamount decimal, ideadline integer, descs varchar(500), trid integer ) returns integer as 
$$
    declare
        transactioninst transaction;
        thecxp integer;
        kebdinfoid integer;
        kebd kebd_transaction_info;
    begin
        select * from transaction into transactioninst where transaction_id= trid;
        if not found then
            return -1;
        end if;
        if transactioninst.state >= 5 then
            return -2;
        end if;
        if transactioninst.seller_id <> sellerid then
            return -6;
        end if;
        select id into thecxp from merchant where id=sellerid;
        if not found then
            select id into thecxp from agent where id = sellerid;
            if not found then
                return -4;
            end if;
        end if;
        select * into kebd from kebd_transaction_info where transaction_id=trid;
        if found and kebd.deadline = ideadline and kebd.description=descs 
        and kebd.kebd_amount=iamount then
            return -3;
        elseif found then
                update kebd_transaction_info set transaction_id=trid,state=4,kebd_amount=iamount,deadline=ideadline,description=descs where kebd_transaction_info_id=kebd.kebd_transaction_info_id returning kebd_transaction_info_id into kebdinfoid;
                if found then
                    return kebd.kebd_transaction_info_id;
                end if;
        end if;
        insert into kebd_transaction_info(transaction_id,state,kebd_amount,deadline,description)
        values( trid,4, iamount,ideadline,descs) returning kebd_transaction_info_id into kebdinfoid;
        if not found then
            return -5;
        end if;
        update transaction set state=4 where transaction_id= trid;
        return kebdinfoid;
    end;
$$ language plpgsql;

create or replace function createNewTrasactionAmendmentRequest( istate integer , trid integer,descs varchar(500) , prices decimal , quant integer) returns integer as 
$$
    declare 
        requestid integer;
        dtransaction transaction;
    begin
        select * into dtransaction from transaction where transaction_id =trid;
        if not found then
            return -1;
        end if;
    insert into transaction_changes( state,transaction_id,
    description,price,qty) values(istate,trid,descs,prices,quant) returning transaction_changes_id into requestid;
    if not found then
        return -2;
    end if;
    update transaction set state=2 where transaction_id=dtransaction.transaction_id returning transaction_id into trid;
    if not found then
        rollback;
        return -3;
    end if;
    return requestid;
    end;
$$ language plpgsql;

create or replace function createGuaranteeRequest( buyerid integer , iamount decimal, descs varchar(500), trid integer ) returns integer as 
$$
    declare
        transactioninst transaction;
        thecxp integer;
        guaranteeinfoid integer;
        guarantee transaction_guarantee_info;
    begin
        select * from transaction into transactioninst where transaction_id= trid;
        if not found then
            return -1;
        end if;
        if transactioninst.state >= 9 then
            return -2;
        end if;
        if transactioninst.requester_id <> buyerid then
            return -6;
        end if;
        select id into thecxp from merchant where id=buyerid;
        if not found then
                return -4;
        end if;
        select * into guarantee from transaction_guarantee_info where transaction_id=trid;
        if found and guarantee.description=descs 
        and guarantee.amount=iamount then
            return -3;
        elseif found then
                update transaction_guarantee_info set  state=7,amount=iamount,description=descs where transaction_guarantee_info_id=guarantee.transaction_guarantee_info_id returning transaction_guarantee_info_id into guaranteeinfoid;
                if found then
                    return guarantee.transaction_guarantee_info_id;
                end if;
        end if;
        insert into transaction_guarantee_info(transaction_id,state,amount,description)
        values( trid,7, iamount,descs) returning transaction_guarantee_info_id into guaranteeinfoid;
        if not found then
            return -5;
        end if;
        update transaction set state=7 where transaction_id= trid;
        return guaranteeinfoid;
    end;
$$ language plpgsql;