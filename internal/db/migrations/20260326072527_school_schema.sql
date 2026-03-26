-- +goose Up

-- USERS (add role)
ALTER TABLE users 
ADD COLUMN role TEXT NOT NULL DEFAULT 'student';

--------------------------------------------------
-- CLASSES
--------------------------------------------------

CREATE TABLE classes (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    section TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT unique_class UNIQUE (name, section)
);

--------------------------------------------------
-- SUBJECTS
--------------------------------------------------

CREATE TABLE subjects (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

--------------------------------------------------
-- STUDENTS
--------------------------------------------------

CREATE TABLE students (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    full_name TEXT NOT NULL,
    roll_number TEXT NOT NULL UNIQUE,
    class_id UUID REFERENCES classes(id),
    date_of_birth DATE,
    created_at TIMESTAMP DEFAULT NOW()
);

--------------------------------------------------
-- FACULTY
--------------------------------------------------

CREATE TABLE faculty (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    full_name TEXT NOT NULL,
    department TEXT,
    designation TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

--------------------------------------------------
-- ENROLLMENTS (optional future use)
--------------------------------------------------

CREATE TABLE enrollments (
    id UUID PRIMARY KEY,
    student_id UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT unique_enrollment UNIQUE (student_id, class_id)
);

--------------------------------------------------
-- TEACHING ASSIGNMENTS
--------------------------------------------------

CREATE TABLE teaching_assignments (
    id UUID PRIMARY KEY,
    faculty_id UUID NOT NULL REFERENCES faculty(id) ON DELETE CASCADE,
    subject_id UUID NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT unique_teaching UNIQUE (faculty_id, subject_id, class_id)
);

--------------------------------------------------
-- +goose Down

DROP TABLE teaching_assignments;
DROP TABLE enrollments;
DROP TABLE faculty;
DROP TABLE students;
DROP TABLE subjects;
DROP TABLE classes;
