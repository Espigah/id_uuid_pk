package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/infra/database"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/infra/metrics"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/group"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/mark"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/student"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/subject"
	subjectTeacherGroup "github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/subject_teacher_group"
	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/school/database/teacher"
	"github.com/rs/xid"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const prometheusScrapeTiming = 15

func main() {

	reg := prometheus.NewRegistry()

	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	testInDocker(reg)
	// testSerial(reg)
	// testUUID(reg)
	// testPK(reg)

}

func testInDocker(reg *prometheus.Registry) {
	modeOption := mode.GetMode()
	test(modeOption, reg)
}

func testSerial(reg *prometheus.Registry) {
	os.Setenv("DATABASE_URL", "postgres://admin:admin@db_serial:5432/postgres")
	test(mode.Serial, reg)
}

func testUUID(reg *prometheus.Registry) {
	os.Setenv("DATABASE_URL", "postgres://admin:admin@db_uuid:5432/postgres")
	test(mode.UUID, reg)
}

func testPK(reg *prometheus.Registry) {
	os.Setenv("DATABASE_URL", "postgres://admin:admin@db_public_key:5432/postgres")
	test(mode.PublicKey, reg)
}

func test(modeOption mode.Mode, reg *prometheus.Registry) {
	modeString := mode.New(modeOption).String()
	customMetrics := metrics.New(reg, modeString)

	teacherFactory, subjectFactory, groupFactory, subjectTeacherGroupFactory, studentFactory, markFactory, dispose := createFactories(modeOption)
	defer dispose()

	log.Printf("Starting %s mode\n", modeString)

	responseTimeout := 60 * time.Minute

	deadline := time.Now().Add(responseTimeout)
	i := 0
	for time.Now().Before(deadline) {
		i++

		mainID := xid.New().String()

		// a cada +ou- 2 scrape do prometheus, reutiliza alguma label ja gerada
		// se gerar muitas labels, o prometheus pode travar
		id := strconv.Itoa(i % (prometheusScrapeTiming * 2))

		ctx := context.WithValue(context.Background(), "mainID", mainID)
		ctx = context.WithValue(ctx, "id", id)

		doWriteTest(ctx, subjectFactory, groupFactory, teacherFactory, studentFactory, subjectTeacherGroupFactory, markFactory, customMetrics)
		doReadTest(ctx, teacherFactory, customMetrics)

		customMetrics.Count()

	}
}

func doReadTest(ctx context.Context, teacherFactory func() teacher.Teacher, customMetrics *metrics.Metrics) {
	id := ctx.Value("id").(string)
	start := time.Now()
	teacherLatest, err := teacherFactory().GetLatest()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic(err)
	}
	teacherFactory().GetTeacherMark(teacherLatest)
	customMetrics.ReadDuration(id, time.Since(start).Seconds())
}

func doWriteTest(ctx context.Context, subjectFactory func() subject.Subject, groupFactory func() group.Group, teacherFactory func() teacher.Teacher, studentFactory func() student.Student, subjectTeacherGroupFactory func() subjectTeacherGroup.SubjectTeacherGroup, markFactory func() mark.Mark, customMetrics *metrics.Metrics) {
	id := ctx.Value("id").(string)
	start := time.Now()
	subjects := createSubjects(ctx, subjectFactory)
	groups := createGroups(ctx, groupFactory)
	teachers := createTeachers(ctx, teacherFactory)
	students := createStudents(ctx, studentFactory, groups)

	createSubjectTeacherGroup(ctx, subjectTeacherGroupFactory, subjects, groups, teachers)
	createMarksForStudents(ctx, markFactory, students, subjects)

	customMetrics.WriteDuration(id, time.Since(start).Seconds())
}

func createFactories(modeOption mode.Mode) (func() teacher.Teacher, func() subject.Subject, func() group.Group, func() subjectTeacherGroup.SubjectTeacherGroup, func() student.Student, func() mark.Mark, func()) {
	database := database.New()
	dispose := func() {
		database.Close(context.Background())
	}
	teacherFactory, err := teacher.NewFactory(
		teacher.Input{
			Mode: modeOption,
			DB:   database,
		})

	if err != nil {
		panic(err)
	}

	subjectFactory, err := subject.NewFactory(
		subject.Input{
			Mode: modeOption,
			DB:   database,
		})

	if err != nil {
		panic(err)
	}

	groupFactory, err := group.NewFactory(
		group.Input{
			Type: modeOption,
			DB:   database,
		})

	if err != nil {
		panic(err)
	}

	subjectTeacherGroupFactory, err := subjectTeacherGroup.NewFactory(
		subjectTeacherGroup.Input{
			Mode: modeOption,
			DB:   database,
		},
	)

	if err != nil {
		panic(err)
	}

	studentFactory, err := student.NewFactory(
		student.Input{
			Mode: modeOption,
			DB:   database,
		},
	)

	if err != nil {
		panic(err)
	}

	markFactory, err := mark.NewFactory(
		mark.Input{
			Mode: modeOption,
			DB:   database,
		},
	)

	if err != nil {
		panic(err)
	}
	return teacherFactory, subjectFactory, groupFactory, subjectTeacherGroupFactory, studentFactory, markFactory, dispose
}

func createSubjects(ctx context.Context, subjectFactory func() subject.Subject) []*subject.Model {
	subjects := []*subject.Model{}
	mainID := ctx.Value("mainID").(string)

	for i := 0; i < 10; i++ {
		subjectModel, err := subjectFactory().Create(
			subject.CreateInput{
				Title: "subject-" + mainID + "-" + strconv.Itoa(i),
			},
		)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			panic(err)
		}

		subjects = append(subjects, subjectModel)
	}

	return subjects
}

func createGroups(ctx context.Context, groupFactory func() group.Group) []*group.Model {
	groups := []*group.Model{}
	mainID := ctx.Value("mainID").(string)

	for i := 0; i < 10; i++ {
		groupModel, err := groupFactory().Create(
			group.CreateInput{
				Name: "group-" + mainID + "-" + strconv.Itoa(i),
			},
		)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			panic(err)
		}

		groups = append(groups, groupModel)
	}

	return groups
}

func createTeachers(ctx context.Context, teacherFactory func() teacher.Teacher) []*teacher.Model {
	teachers := []*teacher.Model{}
	mainID := ctx.Value("mainID").(string)

	for i := 0; i < 10; i++ {
		teacherModel, err := teacherFactory().Create(
			teacher.CreateInput{
				Name:      "teacher-" + mainID + "-" + strconv.Itoa(i),
				PublicKey: xid.New(),
			},
		)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			panic(err)
		}

		teachers = append(teachers, teacherModel)
	}

	return teachers
}

func createStudents(ctx context.Context, studentFactory func() student.Student, groups []*group.Model) []*student.Model {
	students := []*student.Model{}
	mainID := ctx.Value("mainID").(string)

	for i := 0; i < 10; i++ {
		for _, g := range groups {
			studentModel, err := studentFactory().Create(
				student.CreateInput{
					Name:      "student-" + mainID + "-" + strconv.Itoa(i),
					GroupID:   g.ID,
					GroupUUID: g.UUID,
				},
			)

			if err != nil {
				fmt.Printf("error: %v\n", err)
				panic(err)
			}

			students = append(students, studentModel)
		}
	}

	return students
}

func createSubjectTeacherGroup(ctx context.Context, subjectTeacherGroupFactory func() subjectTeacherGroup.SubjectTeacherGroup, subjects []*subject.Model, groups []*group.Model, teachers []*teacher.Model) {

	for _, s := range subjects {
		for _, t := range teachers {
			for _, g := range groups {
				_, err := subjectTeacherGroupFactory().Create(
					subjectTeacherGroup.CreateInput{
						SubjectID:   s.ID,
						TeacherID:   t.ID,
						GroupID:     g.ID,
						SubjectUUID: s.UUID,
						TeacherUUID: t.UUID,
						GroupUUID:   g.UUID,
					},
				)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}
		}
	}
}

func createMarksForStudents(ctx context.Context, markFactory func() mark.Mark, students []*student.Model, subjects []*subject.Model) {
	for i, stu := range students {
		for j, sub := range subjects {
			value := i * j // "random" value
			_, err := markFactory().Create(
				mark.CreateInput{
					StudentID:   stu.ID,
					SubjectID:   sub.ID,
					Mark:        int64(value),
					StudentUUID: stu.UUID,
					SubjectUUID: sub.UUID,
				},
			)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}
	}
}
