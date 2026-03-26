package seeder

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func Seed(ctx context.Context, queries *sqlc.Queries) error {
	log.Println("Starting seeding...")

	if err := seedSubjects(ctx, queries); err != nil {
		return err
	}

	if err := seedClasses(ctx, queries); err != nil {
		return err
	}

	if err := seedStudents(ctx, queries); err != nil {
		return err
	}

	log.Println("✅ Seeding completed")
	return nil
}

func seedSubjects(ctx context.Context, q *sqlc.Queries) error {
	subjects := []string{
		"Mathematics",
		"Science",
		"English",
		"History",
	}

	for _, name := range subjects {
		_, err := q.CreateSubject(ctx, sqlc.CreateSubjectParams{
			ID:   utils.NewPGUUID(),
			Name: name,
		})

		if err != nil {
			log.Println("subject insert skipped:", err)
		}
	}

	log.Println(" Subjects seeded")
	return nil
}

func seedClasses(ctx context.Context, q *sqlc.Queries) error {
	classes := []struct {
		name    string
		section string
	}{
		{"Grade 1", "A"},
		{"Grade 1", "B"},
		{"Grade 10", "A"},
	}

	for _, c := range classes {
		_, err := q.CreateClass(ctx, sqlc.CreateClassParams{
			ID:      utils.NewPGUUID(),
			Name:    c.name,
			Section: pgtype.Text{String: c.section, Valid: true},
		})

		if err != nil {
			log.Println("class insert skipped:", err)
		}
	}

	log.Println("🏫 Classes seeded")
	return nil
}

func seedStudents(ctx context.Context, q *sqlc.Queries) error {
	type studentData struct {
		email      string
		password   string
		fullName   string
		rollNumber string
	}

	students := []studentData{
		{"john@example.com", "password123", "John Doe", "R001"},
		{"jane@example.com", "password123", "Jane Smith", "R002"},
	}

	class, err := q.GetAnyClass(ctx)
	if err != nil {
		return err
	}

	for _, s := range students {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = q.CreateUser(ctx, sqlc.CreateUserParams{
			ID:       utils.NewPGUUID(),
			Email:    s.email,
			Password: string(hashedPassword),
			Role:     "student",
		})

		if err != nil {
			log.Println("user insert skipped:", err)
		}

		user, err := q.GetUserByEmail(ctx, s.email)
		if err != nil {
			return err
		}

		_, err = q.CreateStudent(ctx, sqlc.CreateStudentParams{
			ID:         utils.NewPGUUID(),
			UserID:     user.ID,
			FullName:   s.fullName,
			RollNumber: s.rollNumber,
			ClassID:    class.ID,
		})

		if err != nil {
			log.Println("student insert skipped:", err)
		}
	}

	log.Println("👨‍🎓 Students seeded")
	return nil
}
