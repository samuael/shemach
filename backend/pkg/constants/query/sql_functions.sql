-- registerUserForVarificationWithPhoneNumber checks the availability of the number in the database and insert into temporary registration field if not found.
create or replace function registerUserForVarificationWithPhoneNumber(phoneNumber varchar, shortCode varchar, trials integer) returns integer as 
$$
    declare
    -- -1 if number found in already registered users
    -- -2 if insertion is not successful
    -- -{created_at} if number found in pending code confirmation
    -- {created_at} otherwise
    createdat integer;
    begin
        select created_at into createdat from users where phone=phoneNumber;
        if found then
            return -1;
        end if;
        select created_at into createdat from tempo_registration_info where phone=phoneNumber ;
        if found then
            return -1 * createdat;
        end if;
        insert into tempo_registration_info(phone,code,trials) values(phoneNumber,shortCode,trials) returning created_at into createdAt;
        if not found then
            return -2;
        end if;
        return createdat;
    end;
$$ language plpgsql;

-- checkVerificationCode used to check a user from the phone number waiting to be confirmed but has an expiration date.
-- then forward the confirmation information to verified number where a user with the complete information will awaited to complete the registration.
create or replace function checkVerificationCode( phoneNumber varchar, shortCode varchar) returns bigint as
$$
    declare 
    -- 0 for success
    -- -1 for incorrect
    -- -2 for not found
    -- -3 for trial exceeded
    -- -4 for internal operation error
    -- -5 for verified phone creation error
        verifiedPhoneID bigint;
        varificationData tempo_registration_info;
    begin
        select * into varificationData from tempo_registration_info where phone =phoneNumber;
        if not found then
            return -2;
        end if;

        if varificationData.trials < 0 then
            return -3;
        end if;

        if varificationData.code != shortCode then
            update tempo_registration_info set trials = varificationData.trials - 1 where phone = phoneNumber;
            if found and varificationData.trials - 1 <0 then
                return -3;
            elseif found then
                return -1;
            else
                return -4;
            end if;
        else
            delete from tempo_registration_info where id = varificationData.id;
            if not found then
                return -4;
            end if;

            insert into verified_phones(phone) values(varificationData.phone) returning id into verifiedPhoneID;
            if not found then
                return -5;
            end if;
            return verifiedPhoneID;
        end if;
    end;
$$ language plpgsql;


-- registerForgotPasswordShortCodePhoneNumber checks if the phone number is undergoing a forgot password procedure
create or replace function registerForgotPasswordShortCodePhoneNumber(phoneNumber varchar, shortCode varchar, phone_used smallint, email_used smallint, telegram_used smallint,  maxTrials integer) returns statusandmsg as 
$$
    declare
    -- 1 if trial exceeded try again later
    -- 2 update was not succesful
    -- 3 insertion was not succesful
    -- {created_at} if number found in pending code confirmation
    createdat integer;
    emailUseCount smallint;
    phoneUseCount smallint;
    telegramUseCount smallint;
    trialCount smallint;
    theShortCode varchar;
    begin
        select created_at,email_sent, phone_sent, telegram_sent, trials, code into createdat, emailUseCount, phoneUseCount, telegramUseCount, trialCount, theShortCode from forgot_password_shortcode where phone=phoneNumber;
        if found then
            if trialCount <=0 then
                return Row(1, theShortCode)::statusandmsg;
            end if;
            update forgot_password_shortcode set email_sent= emailUseCount +email_used  , phone_sent = phoneUseCount +phone_used , telegram_sent = telegramUseCount+telegram_used, trials = trials -1, last_sent= extract(epoch from now()) where phone= phoneNumber;
            if not found then
                -- when it says too many request, we don;t delete the row so that the next time it tries to get a short code it will keep sending him a too many request error message
                return Row(2, theShortCode)::statusandmsg;
            end if;
            return Row(createdat, theShortCode)::statusandmsg;
        end if;
        insert into forgot_password_shortcode(phone,code,trials,email_sent, phone_sent, telegram_sent) values(phoneNumber,shortCode,maxTrials,email_used, phone_used,telegram_used) returning created_at into createdAt;
        if not found then
            return Row(3, shortCode)::statusandmsg;
        end if;
        return Row(createdat, shortCode)::statusandmsg;
    end;
$$ language plpgsql;


-- confirmForgotPasswordShortcode confirms the phone and new short code and returns status code
-- -1 : account information not found
-- -2 : trial already exceeded
-- -3 : shortcode doesn;t match and trial exceeded
-- -4 : shortcode does not match
-- 0 : success
create or replace function confirmForgotPasswordShortcode(phoneNumber varchar, shortCode varchar) returns smallint as
$$
    declare
        remainTrial smallint;
        theCode varchar;
    begin
        update forgot_password_shortcode set trials = trials -1 where phone=phoneNumber returning code, trials into theCode, remainTrial;
        if not found then
            return -1;
        end if;
        if remainTrial < 0 then
            return -2;
        end if;
        if theCode != shortCode and remainTrial = 0 then
            return -3;
        end if;
        if theCode != shortCode then
            return -4;
        end if;
        return 0;
    end
$$ language plpgsql;

-- checkandregisteruserinformation checks if the verified phone number is found in the verified phone number informations and if found, it will register a new user with the specified detailed information
create or replace function checkandregisteruserinformation(verifiedAccID bigint, phoneNUmber varchar, demail varchar, telegramID varchar, dfirstname varchar, dlastname varchar, dpassword varchar, drole integer) returns bigint as
$$
    declare
        -- -1 for not found    
        -- -2 for intrenal database error
        -- -3 for user creation error
        statusCode bigint;
    begin
        delete from verified_phones where id= verifiedAccID and phone=phoneNUmber returning id into statusCode;
        if not found then
            return -1;
        end if;

        insert into users(firstname,lastname,phone,email,telegram,password,role) values(dfirstname, dlastname, phoneNumber, demail, telegramID, dpassword, drole) returning id into statusCode;
        if found then
            return statusCode;
        end if;
        
        insert into verified_phones(id, phone) values(verifiedAccID, phoneNUmber);
        return -2;
    end;
$$ language plpgsql;


CREATE OR REPLACE FUNCTION generate_unique_number()
RETURNS BIGINT AS $$
BEGIN
    RETURN extract(epoch FROM now())::BIGINT * 1000 + (random() * 1000)::BIGINT;
END;
$$ LANGUAGE plpgsql;