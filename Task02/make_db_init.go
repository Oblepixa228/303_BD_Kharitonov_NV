package main

import (
	"fmt"
	"os"
)

func main() {
	sqlFilename := "db_init.sql"

	sqlContent := `-- Удаляем таблицы, если они уже существуют
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS users;

-- Создание таблицы movies
CREATE TABLE movies (
    id INTEGER PRIMARY KEY,
    title TEXT,
    year INTEGER,
    genres TEXT
);

-- Создание таблицы ratings
CREATE TABLE ratings (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    movie_id INTEGER,
    rating REAL,
    timestamp INTEGER
);

-- Создание таблицы tags
CREATE TABLE tags (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    movie_id INTEGER,
    tag TEXT,
    timestamp INTEGER
);

-- Создание таблицы users
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    email TEXT,
    gender TEXT,
    register_date TEXT,
    occupation TEXT
);
`

	err := os.WriteFile(sqlFilename, []byte(sqlContent), 0644)
	if err != nil {
		fmt.Printf("Ошибка записи файла: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Файл db_init.sql успешно сгенерирован!")
}