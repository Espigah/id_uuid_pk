package teacher

import (
	"context"
	"errors"
)

type teacherPublicKey struct {
	teacher
}

func (t *teacherPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO teacher(name) VALUES($1) RETURNING teacher_id, public_key", input.Name).Scan(&model.ID, &model.PublicKey)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating teacher, name:"+input.Name))
		return nil, err

	}

	model.PublicKey = input.PublicKey
	model.Name = input.Name
	return &model, nil
}

func (t *teacherPublicKey) GetLatest() (*Model, error) {
	var model Model

	query := `
		select 
			t.teacher_id
			, t.public_key
			, t."name"  
		from teacher t 
		where t.CTID = (
			select max(CTID) from teacher t 
		)	
	`

	err := t.db.QueryRow(context.Background(), query).Scan(&model.ID, &model.PublicKey, &model.Name)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error get teacher mark"))
		return nil, err
	}

	return &model, nil
}

func (t *teacherPublicKey) GetTeacherMark(input *Model) (interface{}, error) {

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
		where t.public_key = ($1)
		group by 1,2
		order by 3 desc
	`
	err := t.db.QueryRow(context.Background(), query, input.PublicKey).Scan(&model.Name, &model.Subject, &model.Mark)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error get teacher mark"))
		return nil, err
	}

	return nil, nil
}
