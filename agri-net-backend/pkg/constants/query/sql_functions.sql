create or replace function allStudentsOfCategory(integer) returns integer as $$
DECLARE allstudents integer;
BEGIN
select count(*) as count into allstudents
FROM student
    INNER JOIN round ON student.round_id = round.id
    INNER JOIN category ON round.categoryid = category.id
where category.id = $1;
if not found then return 0;
end if;
return allstudents;
END;
$$ language plpgsql;
create or replace function activeStudentsOfCategory(integer) returns integer as $$
DECLARE allstudents integer;
BEGIN
select count(*) as count into allstudents
FROM student
    INNER JOIN round ON student.round_id = round.id
    INNER JOIN category ON round.categoryid = category.id
where category.id = $1
    and round.active = t;
if not found then return 0;
end if;
return allstudents;
END;
$$ language plpgsql;
create or replace function createAdmin(
        ifullname VARCHAR(250),
        iemail VARCHAR(250),
        ipassword text,
        isuperadmin boolean,
        iimgurl VARCHAR(100),
        years integer,
        months integer,
        days integer,
        hours integer,
        minutes integer,
        seconds integer
    ) returns admins as $$
declare theadmin admins;
dateid integer;
theadminid integer;
begin -- SAVEPOINT my_point;
insert into eth_date(year, month, day, hour, minute, second)
values(years, months, days, hours, minutes, seconds)
returning id into dateid;
if not found then ROLLBACK;
return null;
end if;
INSERT INTO admins(
        fullname,
        email,
        password,
        superadmin,
        imgurl,
        created_at
    )
VALUES(
        ifullname,
        iemail,
        ipassword,
        isuperadmin,
        iimgurl,
        dateid
    )
RETURNING id into theadminid;
if found then
select theadminid,
    ifullname,
    iemail,
    ipassword,
    iimgurl,
    isuperadmin,
    dateid into theadmin;
return theadmin;
else rollback;
return null;
end if;
end;
$$ language plpgsql;
-- DROP FUNCTION createadmin(character varying,character varying,text,boolean,character varying,integer,integer,integer,integer,integer,integer)
create or replace function createCategory(
        ititle varchar,
        ishort_title varchar,
        irounds_count integer,
        iimgurl varchar,
        ifee decimal,
        years integer,
        months integer,
        days integer,
        hours integer,
        minutes integer,
        seconds integer
    ) returns category as $$
declare thecategory category;
dateid integer;
thecategoryid integer;
begin -- SAVEPOINT my_point;
insert into eth_date(year, month, day, hour, minute, second)
values(years, months, days, hours, minutes, seconds)
returning id into dateid;
if not found then ROLLBACK;
return null;
end if;
INSERT INTO category(
        title,
        short_title,
        rounds_count,
        imgurl,
        fee,
        created_at
    )
VALUES(
        ititle,
        ishort_title,
        irounds_count,
        iimgurl,
        ifee,
        dateid
    )
RETURNING id into thecategoryid;
if found then
select thecategoryid,
    ititle,
    ishort_title,
    irounds_count,
    iimgurl,
    ifee,
    dateid into thecategory;
return thecategory;
else rollback;
return null;
end if;
end;
$$ language plpgsql;
-- DROP FUNCTION createcategory(character varying,character varying,integer,character varying,numeric,integer,integer,integer,integer,integer,integer)
create or replace function deleteAdmin(varchar(200), integer) returns boolean as $$
declare dateid integer;
begin
select created_at into dateid
from admins
where id = $2;
if not found then return false;
end if;
delete from admins
where id = $2;
delete from eth_date
where id = dateid;
return true;
end;
$$ language plpgsql;
create or replace function deleteAdminByEmail(iemail VARCHAR(250)) returns boolean as $$
declare dateid integer;
begin
select created_at into dateid
from admins
where email = iemail;
if not found then return false;
end if;
delete from admins
where email = iemail;
delete from eth_date
where id = dateid;
return true;
end;
$$ language plpgsql;
create or replace function createRound(
        icategoryid integer,
        itraining_hour integer,
        iround_no integer,
        istudents integer,
        iactive_amount decimal,
        iactive boolean,
        istart_date varchar,
        ilang char,
        iend_date varchar,
        ifee decimal,
        years integer,
        months integer,
        days integer,
        hours integer,
        minutes integer,
        seconds integer
    ) returns round as $$
declare theround round;
dateid integer;
theroundid integer;
begin -- SAVEPOINT my_point;
insert into eth_date(year, month, day, hour, minute, second)
values(years, months, days, hours, minutes, seconds)
returning id into dateid;
if not found then ROLLBACK;
return null;
end if;
INSERT INTO round(
        categoryid,
        training_hour,
        round_no,
        students,
        active_amount,
        active,
        start_date,
        lang,
        end_date,
        fee,
        created_at
    )
VALUES(
        icategoryid,
        itraining_hour,
        iround_no,
        istudents,
        iactive_amount,
        iactive,
        istart_date,
        ilang,
        iend_date,
        ifee,
        dateid
    )
RETURNING id into theroundid;
if found then
select theroundid,
    icategoryid,
    itraining_hour,
    iround_no,
    istudents,
    iactive_amount,
    iactive,
    istart_date,
    ilang,
    iend_date,
    dateid,
    ifee into theround;
return theround;
else rollback;
return null;
end if;
end;
$$ language plpgsql;

-- -----------------------

create or replace function checkExistanceAndActivenessOrRound(integer) returns integer as $$
declare result boolean;
begin
select active
from round
where id = $1 into result;
if not found then return -1;
else if result then return 1;
else return 0;
end if;
end;
$$ language plpgsql;

-- ---------------------

-- -----------------------

create or replace function checkExistanceOfCategory(integer) returns integer as $$
declare 
    result integer;
begin
    select created_at
    from category
    where id = $1 into result;
    
    if not found then return 0;
    else return 1;
    end if; 
end;
$$ language plpgsql;

-- ---------------------

select registerstudent(
        'Mahder Adnew',
        cast ('m' as char),
        cast (12 as decimal),
        cast ('B.Sc.' as varchar),
        cast ('+251982870204' as varchar),
        cast (0.0 as decimal),
        cast (0 as integer),
        cast (21 as integer),
        cast (16 as integer),
        cast('' as varchar),
        cast(null as integer),
        cast (1991 as integer),
        cast (5 as integer),
        cast (14 as integer),
        cast (0 as integer),
        cast (0 as smallint),
        cast (0 as smallint),
        cast (2014 as integer),
        cast (7 as integer),
        cast (4 as integer),
        cast (0 as integer),
        cast (0 as smallint),
        cast (0 as smallint),
        cast ('assosa' as varchar),
        cast ('benishangul' as varchar),
        cast ('assosa' as varchar),
        cast ('assosa' as varchar),
        cast ('04' as varchar),
        cast ('amba ber' as varchar)
    );
create or replace function registerStudent(
        ifullname VARCHAR,
        isex char,
        iage decimal,
        iaccamic_status VARCHAR,
        iphone varchar,
        ipaid decimal,
        istatus integer,
        iregistered_by integer,
        iround_id integer,
        iimgurl VARCHAR,
        imark integer,
        -- 
        byear integer,
        bmonth integer,
        bday integer,
        bhour integer,
        bminute smallint,
        bsecond smallint,
        -- 
        cryear integer,
        crmonth integer,
        crday integer,
        crhour integer,
        crminute smallint,
        crsecond smallint,
        -- 
        icity VARCHAR,
        iregion VARCHAR,
        izone varchar,
        iworeda varchar,
        ikebele VARCHAR,
        iunique_address varchar
    ) returns student as $$
declare dirthdateid integer;
createdatid integer;
addressid integer;
studentid integer;
thestudent student;
theround round;
begin
select * into theround
from round
where id = iround_id;
if not found then return null;
end if;
insert into eth_date(year, month, day, hour, minute, second)
values(byear, bmonth, bday, bhour, bminute, bsecond)
returning id into dirthdateid;
if not found then ROLLBACK;
return null;
end if;
-- 
insert into eth_date(year, month, day, hour, minute, second)
values(
        cryear,
        crmonth,
        crday,
        crhour,
        crminute,
        crsecond
    )
returning id into createdatid;
if not found then ROLLBACK;
return null;
end if;
-- 
insert into addresses(
        city,
        region,
        zone,
        woreda,
        kebele,
        unique_address
    )
values(
        icity,
        iregion,
        izone,
        iworeda,
        ikebele,
        iunique_address
    )
returning id into addressid;
if not found then ROLLBACK;
return null;
end if;
INSERT INTO student(
        fullname,
        sex,
        age,
        birth_date,
        accamic_status,
        address,
        phone,
        paid,
        status,
        registered_by,
        round_id,
        imgurl,
        mark,
        registered_at
    )
VALUES(
        ifullname,
        isex,
        iage,
        dirthdateid,
        iaccamic_status,
        addressid,
        iphone,
        ipaid,
        istatus,
        iregistered_by,
        iround_id,
        iimgurl,
        imark,
        createdatid
    )
RETURNING id into studentid;
if found then
select studentid,
    ifullname,
    isex,
    iage,
    dirthdateid,
    iaccamic_status,
    addressid,
    iphone,
    ipaid,
    istatus,
    iregistered_by,
    iround_id,
    iimgurl,
    imark,
    createdatid into thestudent;
return thestudent;
else rollback;
return null;
end if;
end;
$$ language plpgsql;
-- creating a trigger to update the rounds table while creating a student instacne.
CREATE OR REPLACE FUNCTION increment_rounds_students_count() RETURNS TRIGGER AS $round$ 
BEGIN
    update round as r
        set students = r.students + 1
        where id = new.round_id;
        RETURN NEW;
END;
$round$ language plpgsql;

CREATE TRIGGER trigger_increment_rounds_student_quantity
AFTER
INSERT ON student FOR EACH ROW EXECUTE PROCEDURE increment_rounds_students_count();

-- drop trigger trigger_increment_rounds_student_quantity on studet;




-- This function below creates a special case information and 
-- update the special case mark of the student to the id of the special case.

create or replace function createSpecialCaseForStudent(  ireason varchar , amount decimal  , complete boolean  , studentid integer  )  returns special_case as 
$$
    declare
        specialcaseid integer;
        thecase special_case;
        updateCount integer;
    begin 
        insert into special_case( reason, covered_amount, complete_fee) values(  ireason, amount, complete) returning id into specialcaseid;
        if not found then 
            return null;
        else 
            with rows as ( update student set mark=specialcaseid where id=studentid returning 1 )

            select count(*) into  updateCount from rows;

            if cast ( updateCount as integer ) = cast ( 1 as integer) then
                select specialcaseid , ireason, amount , complete into thecase;
                return thecase;
            else
                rollback;
                return null;
            end if;
        end if;
    end;
$$ language plpgsql;


-- This trigger below is created to make the update on the category fee propagate to the rounds of the category.
create or replace function updateRoundFee() returns trigger as 
$$
    begin
        update round as r
            set fee = new.fee
            where categoryid = new.id;
        RETURN NEW;
    end;
$$
 language plpgsql;

CREATE TRIGGER trigger_round_payment_update
    AFTER
    UPDATE ON category FOR EACH ROW EXECUTE PROCEDURE updateRoundFee();






-- check the existance of date time and unique chars and student id combination 
-- exist in the payments. Using a function 
--  THis returns 
-- -3 if the student is not a student from the specified round.
-- -2 if the round with this id doesn't exist.
-- -1 if the student does not exist
--  0 if the payment instance with this information does not exit  
--  1 if there is.
create or replace function checkTheExsistanceOfPayment( 
    duchars char, 
    studentid integer, 
    droundid integer,
    dyear integer, 
    dmonth integer, 
    dday integer, 
    dhour integer, 
    dminute integer, 
    dsecond integer) returns integer as 
$$
    declare
        thedateid integer;
        thepayinid integer;
        thestudentroundid integer;
        theroundid integer;
    begin

        select round_id from student where id=studentid into thestudentroundid;
        if not found then
            return -1;
        end if; 
        if thestudentroundid <> droundid then
            return -3;
        end if;

        select id from round where id=droundid into theroundid;
        if not found then
            return -2;
        end if;

        select id from eth_date where year=dyear and month=dmonth and 
        day=dday and hour=dhour and minute=dminute and second=dsecond into thedateid;
        if not found then
            return 0;
        end if;
        -- 
        select 
            payin.id from payin INNER JOIN eth_date ON
        eth_date.id=payin.created_at
        WHERE payin.payed_by=studentid and 
        payin.created_at=thedateid and payin.uchars=duchars and payin.roundid = theroundid into thepayinid;
        if not found then
            return 0;
        else
            return 1;
        end if;
    end;
$$ language plpgsql;



-- This method checks for the cost of a round the student is learning is and the amount of money paid by the student 
-- to return the amount of money that has to be paid by the student.
create or replace function getRemainingPaymentOfStudentForRound( studentid integer) returns decimal as 
$$
    declare 
        studentPaidAmount decimal;
        roundid integer;
        roundCost decimal;
    begin
        select paid,round_id from student where id=studentid into studentPaidAmount, roundid;
        if not found then
            return -1;
        end if; 

        select fee from round where id=roundid into roundCost;
        if not found then
            return -2;
        end if;
        
        if studentPaidAmount > roundCost then 
            return 0;
        else 
            return roundCost-studentPaidAmount;
        end if;
    end; 
$$ language plpgsql;

-- Create Payment transaction for the student
create or replace function createPayinTransaction(
    duchars char, 
    studentid integer,
    secretaryid integer,
    droundid integer,
    dyear integer,
    dmonth integer,
    dday integer,
    dhour integer,
    dminute integer,
    dsecond integer,
    damount decimal
    ) returns payin as 
$$
    declare 
        thedateid integer;
        thestudentroundid integer;
        theroundid integer;
        thepaymentid integer;
        thepayment payin;
        remainingPayment decimal;
        isTheAdminaSuperAdmin boolean;

        messageStatus integer;
    begin
        select round_id from student where id=studentid into thestudentroundid;
        if not found then
            RAISE exception SQLSTATE '77777' USING MESSAGE = 'student with this id does not exist';
            return null;
        end if; 
        if thestudentroundid <> droundid then
            return null;
        end if;

        select id from round where id=droundid into theroundid;
        if not found then
            RAISE exception SQLSTATE '77777' USING MESSAGE = 'round with this id does not exist';
            return null;
        end if;
        -- checking the amount of money the student is expected to pay and if the amount is surplus the raise an exception with
        -- some information.
        -- RAISE EXCEPTION SQLSTATE '90001' USING MESSAGE = 'my own error';
        select * from getRemainingPaymentOfStudentForRound(studentid) into remainingPayment;

        select superadmin from admins where id=secretaryid into isTheAdminaSuperAdmin;
        if not found then 
            RAISE exception SQLSTATE '77777' USING MESSAGE = 'admin with this id does not exist';
            return null;
        end if;

        if remainingPayment < 0 then
            RAISE exception SQLSTATE '77777' USING MESSAGE = 'students information is incomplete';
            return null;
        end if;
        if remainingPayment =0 then
             RAISE exception SQLSTATE '77777' USING MESSAGE = 'the student has completed the payment';
             return null;
        end if;

        if remainingPayment < damount then
            RAISE exception SQLSTATE '77777' USING MESSAGE = 'the amount to be paid is more than what is needed';
            return null;
        end if;

        select id from eth_date where year=dyear and month=dmonth and 
        day=dday and hour=dhour and minute=dminute and second=dsecond into thedateid;
        if not found then
            insert into eth_date(year , month,day,hour,minute,second) values(dyear,dmonth,dday,dhour,dminute,dsecond) returning id into thedateid;
            if not found then
                return null;
            end if;
        end if;
        if isTheAdminaSuperAdmin then 
            messageStatus=3;
        else 
            messageStatus=0;
        end if;
        if isTheAdminaSuperAdmin then
            insert into payin(amount,recieved_by,payed_by,created_at,roundid,status,uchars)
            values( damount, secretaryid, studentid, thedateid, droundid,messageStatus,duchars) returning id into thepaymentid;
            
            if not found then
                ROLLBACK;
                return null;
            else         
                select thepaymentid, damount, secretaryid, studentid, thedateid, droundid,messageStatus,duchars into thepayment;
            end if;
        end if;
        return thepayment;

    end;
$$ language plpgsql; 

-- the trigger function for updating the student's paid amount information and the rounds active account information
-- to increament the student payment information.
create or replace function update_the_student_payment_information() returns trigger as 
$$
    begin
        update round as r set active_amount = r.active_amount + new.amount where id=new.roundid;
        update student as s set paid = s.paid + new.amount where id= new.payed_by;
        return NEW;
    end;
$$ language plpgsql;

-- CREATING A TRIGGER WHICH IS TO BE RAISED AFTER AN INSERT ON THE payin TABLE.
create TRIGGER update_the_student_payment_information_trigger AFTER INSERT ON payin 
    FOR EACH ROW EXECUTE PROCEDURE update_the_student_payment_information();
-- 



-- GetPayinInstanceUsingItsInformation 
create or replace function GetPayinInstanceUsingItsInformation( 
    duchars char, 
    studentid integer,
    dyear integer, 
    dmonth integer, 
    dday integer, 
    dhour integer, 
    dminute integer, 
    dsecond integer) returns payin as 
$$
    declare
        thedateid integer;
        thepayin payin;
    begin
        select id from eth_date where year=dyear and month=dmonth and 
        day=dday and hour=dhour and minute=dminute and second=dsecond limit 1 into thedateid;
        if not found then
            return null;
        end if;
        -- 
        select 
            payin.id,
            payin.amount,
            payin.recieved_by,
            payin.payed_by,
            payin.created_at,
            payin.roundid,
            payin.status,
            payin.uchars
            from payin INNER JOIN eth_date ON
        eth_date.id=payin.created_at
        WHERE payin.payed_by=studentid and 
        payin.created_at=thedateid and payin.uchars=duchars into thepayin;
        if not found then
            return null;
        else
            return thepayin;
        end if;
    end;
$$ language plpgsql;
-- ------------------------------------------------------------------------------------




-- ------------Deleting date IDS after delete on the Payin table --------
--  --------------- The Trigger function --------------------------------

create or replace function deletePayinDates() returns trigger as 
$$
    declare 
        theid integer;
    begin
        select id into theid from payin where created_at=old.created_at;
        if found then 
            return NULL;
        end if;

        select id into theid from admins where created_at=old.created_at;
        if found then
            return NULL;
        end if;

        select id into theid from category where created_at=old.created_at;
        if found then
            return NULL;
        end if;

        select id into theid from payout where created_at=old.created_at;
        if found then
            return NULL;
        end if;

        select id into theid from round where created_at=old.created_at;
        if found then
            return NULL;
        end if;

        select id into theid from student where registered_at=old.created_at or birth_date=old.created_at;
        if found then 
            return NULL;
        end if;
        delete from eth_date where id=old.created_at;
        return NULL;
    end;
$$ language plpgsql;

-- ------------- Creating a trigger to use the above function and instantiate after deletion on any table -----
create trigger delete_orphan_eth_date_after_deletion_on_payin AFTER DELETE ON payin 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();
create trigger delete_orphan_eth_date_after_deletion_on_admins AFTER DELETE ON admins 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();
create trigger delete_orphan_eth_date_after_deletion_on_category AFTER DELETE ON category 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();
create trigger delete_orphan_eth_date_after_deletion_on_payout AFTER DELETE ON payout 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();
create trigger delete_orphan_eth_date_after_deletion_on_round AFTER DELETE ON round 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();
create trigger delete_orphan_eth_date_after_deletion_on_student AFTER DELETE ON student 
    FOR EACH ROW EXECUTE PROCEDURE deletePayinDates();

-- ---- a function to update the payment status
-- return value 
--  -3 -- for invalid database query
--  -2 -- for invalid database query status parameter
--  -1 -- payin instance with this id doesn't exist
--   0 -- payment status is already rejected;
--   1 -- payment status is already accepted
--   2 -- payment status updated
create or replace function updatePayInStatus( payinid integer , pstatus integer ) returns integer as 
$$
    declare
        paymentstatus integer;
    begin
        if pstatus > 3  or  pstatus < 0 then 
            return -2;
        end if;

        select status into paymentstatus from payin where id=payinid;
        if not found then 
            return  -1;
        end if;

        if paymentstatus >= pstatus then 
            if paymentstatus = 3 then 
                return 1;
            end if;
            if paymentstatus =2 then 
                return 0;
            end if;
        end if;

        update payin set status=pstatus where id=payinid returning id into payinid;
        if not found then
            return -3;
        end if;

        return 2;
    end;
$$ language plpgsql;