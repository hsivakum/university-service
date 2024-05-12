-- Enums
CREATE TYPE GENDER AS ENUM ('MALE', 'FEMALE', 'OTHER');

CREATE TYPE STATE AS ENUM (
    'AL', 'AK', 'AZ', 'AR', 'CA', 'CO', 'CT', 'DE', 'FL', 'GA',
    'HI', 'ID', 'IL', 'IN', 'IA', 'KS', 'KY', 'LA', 'ME', 'MD',
    'MA', 'MI', 'MN', 'MS', 'MO', 'MT', 'NE', 'NV', 'NH', 'NJ',
    'NM', 'NY', 'NC', 'ND', 'OH', 'OK', 'OR', 'PA', 'RI', 'SC',
    'SD', 'TN', 'TX', 'UT', 'VT', 'VA', 'WA', 'WV', 'WI', 'WY'
    );


CREATE TYPE SEMESTER AS ENUM ('Fall', 'Winter', 'Spring', 'Summer');

CREATE TYPE PROGRAM_LEVEL AS ENUM ('UG', 'PG', 'DOCTORATE', 'ASSOCIATE', 'EXECUTIVE');

CREATE TYPE RACE AS ENUM (
    'White', 'Black or African American', 'American Indian or Alaska Native',
    'Asian', 'Native Hawaiian or Other Pacific Islander', 'Two or more races', 'Other'
    );

CREATE TYPE GRADE AS ENUM ('A', 'A-', 'B+', 'B', 'B-');

CREATE TYPE "DAY" AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');

CREATE TYPE DESIGNATION_LEVEL AS ENUM (
    'President',
    'Vice president',
    'Dean',
    'Tenure Professor',
    'Professor',
    'Associate Professor',
    'Instructor 1',
    'Instructor 2',
    'Instructor 3'
    );


CREATE TABLE address
(
    id     SERIAL PRIMARY KEY,
    street VARCHAR(255),
    city   VARCHAR(100),
    state  State,
    zip    INTEGER
);

CREATE TABLE employee
(
    emp_id     SERIAL PRIMARY KEY,
    name       VARCHAR(255),
    phone      VARCHAR(20),
    email      VARCHAR(255),
    gender     GENDER,
    address_id INTEGER REFERENCES address (id)
);

CREATE TABLE college
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255),
    president_id INTEGER REFERENCES employee (emp_id)
);

CREATE TABLE department
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255),
    dean_id INTEGER REFERENCES employee (emp_id)
);

CREATE TABLE college_department
(
    college_id    INTEGER REFERENCES college (id),
    department_id INTEGER REFERENCES department (id),
    PRIMARY KEY (college_id, department_id)
);

CREATE TABLE faculty
(
    emp_id            INTEGER PRIMARY KEY REFERENCES employee (emp_id),
    department_id     INTEGER REFERENCES department (id),
    designation_level DESIGNATION_LEVEL
);

CREATE TABLE program
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255),
    level           PROGRAM_LEVEL,
    department_id   INTEGER REFERENCES department (id),
    duration_months INTEGER
);

CREATE TABLE student
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(255),
    email               VARCHAR(255),
    enrolled_program_id INTEGER REFERENCES program (id),
    start_year          INTEGER,
    semester            SEMESTER,
    gpa                 INTEGER,
    active              BOOLEAN,
    is_citizen          BOOLEAN,
    race                RACE,
    dob                 DATE
);

CREATE TABLE courses
(
    id                     SERIAL PRIMARY KEY,
    number                 INTEGER,
    name                   VARCHAR(255),
    credits                INTEGER,
    pre_req_id             INTEGER,
    co_requisite_id        INTEGER,
    seat                   INTEGER,
    waiting_list_seat      INTEGER,
    no_of_classes_per_week INTEGER
);

CREATE TABLE course_offered
(
    id            SERIAL PRIMARY KEY,
    course_id     INTEGER REFERENCES courses (id),
    department_id INTEGER REFERENCES department (id),
    taught_by     INTEGER REFERENCES faculty (emp_id),
    semester      SEMESTER,
    year          INTEGER,
    minutes       INTEGER,
    hours         INTEGER
);

CREATE TABLE course_registration
(
    student_id        INTEGER REFERENCES student (id),
    course_offered_id INTEGER REFERENCES course_offered (id),
    PRIMARY KEY (student_id, course_offered_id)
);

CREATE TABLE student_grades
(
    course_offered_id INTEGER REFERENCES course_offered (id),
    student_id        INTEGER REFERENCES student (id),
    grade             GRADE,
    PRIMARY KEY (course_offered_id, student_id)
);

CREATE TABLE course_schedule
(
    course_offered_id INTEGER REFERENCES course_offered (id),
    time_slot_from    TIMESTAMP,
    time_slot_to      TIMESTAMP,
    day               "DAY",
    PRIMARY KEY (course_offered_id, day)
);