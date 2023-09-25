CREATE TABLE public.teacher (
	teacher_id serial4 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT teacher_pk PRIMARY KEY (teacher_id)
);

CREATE TABLE public.subject (
	subject_id serial4 NOT NULL,
	title text NOT NULL,
	CONSTRAINT subject_pk PRIMARY KEY (subject_id)
);


CREATE TABLE public."group" (
	group_id serial4 NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT group_pk PRIMARY KEY (group_id)
);

CREATE TABLE public.student (
	student_id serial4 NOT NULL,
	"name" text NOT NULL,
	group_id int4 NULL,
	CONSTRAINT student_pk PRIMARY KEY (student_id),
	CONSTRAINT student_fk FOREIGN KEY (group_id) REFERENCES public."group"(group_id)
);

CREATE TABLE public.mark (
	mark_id serial4 NOT NULL,
	student_id int4 NULL,
	subject_id int4 NULL,
	mark int4 NULL,
	CONSTRAINT mark_pk PRIMARY KEY (mark_id),
	CONSTRAINT mark_fk_student FOREIGN KEY (student_id) REFERENCES public.student(student_id),
	CONSTRAINT mark_fk_subject FOREIGN KEY (subject_id) REFERENCES public.subject(subject_id)
);


CREATE TABLE public.subject_teacher_group (
	subject_teacher_group_id serial4 NOT NULL,
	subject_id int4 NOT NULL,
	teacher_id int4 NOT NULL,
	group_id int4 NOT NULL,
	CONSTRAINT subject_teacher_group_pk PRIMARY KEY (subject_teacher_group_id),
	CONSTRAINT subject_teacher_group_fk FOREIGN KEY (group_id) REFERENCES public."group"(group_id),
	CONSTRAINT subject_teacher_group_fk_subject FOREIGN KEY (subject_id) REFERENCES public.subject(subject_id),
	CONSTRAINT subject_teacher_group_fk_teacher FOREIGN KEY (teacher_id) REFERENCES public.teacher(teacher_id)
);