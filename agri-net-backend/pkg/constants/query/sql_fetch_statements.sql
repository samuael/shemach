-- selecting all students of a category using inner join statements.
select 
student.id,
student.fullname,
student.sex,
student.age,
student.birth_date,
student.accamic_status,
student.address,
student.phone,
student.paid,
student.status,
student.registered_by,
student.round_id,
student.imgurl,
student.mark,
student.registered_at from student INNER JOIN round ON round.id=student.round_id
INNER JOIN category ON category.id = round.categoryid WHERE category.id= 3;


