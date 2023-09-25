package teacher

import (
	"context"
	"errors"
)

type teacherSerial struct {
	teacher
}

func (t *teacherSerial) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO teacher(name) VALUES($1) RETURNING (teacher_id)", input.Name).Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating teacherSerial, name:"+input.Name))
		return nil, err

	}

	model.Name = input.Name
	return &model, nil
}

func (t *teacherSerial) GetLatest() (*Model, error) {
	var model Model

	query := `
		select 
			t.teacher_id
			, t."name"  
		from teacher t 
		where t.CTID = (
			select max(CTID) from teacher t 
		)	
	`

	err := t.db.QueryRow(context.Background(), query).Scan(&model.ID, &model.Name)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error get teacher mark"))
		return nil, err
	}

	return &model, nil
}

func (t *teacherSerial) GetTeacherMark(input *Model) (interface{}, error) {

	var model TeacherMark

	query := `
		select  
			t."name"
			, s.title as subject
			, sum(m.mark) as mark 
		from teacher t
		join subject_teacher_group stg on stg.teacher_id = t.teacher_id 
		join subject s on s.subject_id = stg.subject_id  
		join "group" g on g.group_id = stg.group_id 
		join student st on st.group_id = g.group_id 
		join mark m on m.student_id = st.student_id and m.subject_id = s.subject_id 
		where t.teacher_id = ($1)
		group by 1,2
		order by 3 desc
	`
	err := t.db.QueryRow(context.Background(), query, input.ID).Scan(&model.Name, &model.Subject, &model.Mark)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error get teacher mark"))
		return nil, err
	}

	return nil, nil
}
