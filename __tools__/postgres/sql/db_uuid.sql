CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.teacher (
	teacher_id uuid DEFAULT uuid_generate_v4 (),
	"name" text NOT NULL,
	CONSTRAINT teacher_pk PRIMARY KEY (teacher_id)
);

CREATE TABLE public.subject (
	subject_id uuid DEFAULT uuid_generate_v4 (),
	title text NOT NULL,
	CONSTRAINT subject_pk PRIMARY KEY (subject_id)
);


CREATE TABLE public."group" (
	group_id uuid DEFAULT uuid_generate_v4 (),
	"name" text NOT NULL,
	CONSTRAINT group_pk PRIMARY KEY (group_id)
);

CREATE TABLE public.student (
	student_id uuid DEFAULT uuid_generate_v4 (),
	"name" text NOT NULL,
	group_id uuid NULL,
	CONSTRAINT student_pk PRIMARY KEY (student_id),
	CONSTRAINT student_fk FOREIGN KEY (group_id) REFERENCES public."group"(group_id)
);

CREATE TABLE public.mark (
	mark_id uuid DEFAULT uuid_generate_v4 (),
	student_id uuid NULL,
	subject_id uuid NULL,
	mark int4 NULL,
	CONSTRAINT mark_pk PRIMARY KEY (mark_id),
	CONSTRAINT mark_fk_student FOREIGN KEY (student_id) REFERENCES public.student(student_id),
	CONSTRAINT mark_fk_subject FOREIGN KEY (subject_id) REFERENCES public.subject(subject_id)
);


CREATE TABLE public.subject_teacher_group (
	subject_teacher_group_id uuid DEFAULT uuid_generate_v4 (),
	subject_id uuid NOT NULL,
	teacher_id uuid NOT NULL,
	group_id uuid NOT NULL,
	CONSTRAINT subject_teacher_group_pk PRIMARY KEY (subject_teacher_group_id),
	CONSTRAINT subject_teacher_group_fk FOREIGN KEY (group_id) REFERENCES public."group"(group_id),
	CONSTRAINT subject_teacher_group_fk_subject FOREIGN KEY (subject_id) REFERENCES public.subject(subject_id),
	CONSTRAINT subject_teacher_group_fk_teacher FOREIGN KEY (teacher_id) REFERENCES public.teacher(teacher_id)
);