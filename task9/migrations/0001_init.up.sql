CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    parent_id INT REFERENCES groups(id)
);

CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    birth_year INT,
    group_id INT REFERENCES groups(id)
);

-- тестовые данные
INSERT INTO groups(name, parent_id) VALUES
('Root', NULL),
('Team A', 1),
('Team B', 1),
('Subteam A1', 2);

INSERT INTO people(first_name,last_name,birth_year,group_id) VALUES
('Ivan','Ivanov',1990,2),
('Petr','Petrov',1985,3),
('Anna','Sidorova',2000,4);