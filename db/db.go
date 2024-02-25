package db

import (
	"api/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func DBConnect() {
	var err error
	DB, err = sql.Open("mysql", "root:@Arjun897@tcp(127.0.0.1)/ARJUN")
	if err != nil {
		fmt.Errorf("error while connecting db , err = %v", err)
		panic(err)
	}
	fmt.Println("Connection is Successful")
}

func GetMovies() []types.Movies {
	query := `SELECT NAME,LANGUAGE,BUDGET,COLLECTION FROM MOVIES`
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("Error while querying [%v], err = %v", query, err)
		panic(err)
	}
	var movies []types.Movies
	for rows.Next() {
		movie := types.Movies{}
		rows.Scan(&movie.Name,
			&movie.Language,
			&movie.Budget,
			&movie.Collection)
		movies = append(movies, movie)
	}
	return movies

}

func GetMovieByName(name string) (bool, types.Movies) {
	query := `SELECT NAME,LANGUAGE,BUDGET,COLLECTION FROM MOVIES WHERE NAME=?`
	rows, err := DB.Query(query, name)
	if err != nil {
		fmt.Errorf("Error while querying [%v], err = %v", query, err)
		panic(err)
	}
	defer rows.Close()
	movie := types.Movies{}
	present := false
	for rows.Next() {
		err := rows.Scan(&movie.Name,
			&movie.Language,
			&movie.Budget,
			&movie.Collection)
		if err != nil {
			panic(err)
		}
		present = true
	}
	return present, movie
}

func UpdateMovie(name string, movie types.Movies) bool {
	present, _ := GetMovieByName(name)
	if !present {
		return false
	}
	query := `UPDATE MOVIES SET NAME=?,LANGUAGE=?,BUDGET=?,COLLECTION=? WHERE NAME=?`
	statement, err := DB.Prepare(query)
	if err != nil {
		fmt.Errorf("Error while querying [%v], err = %v", query, err)
		panic(err)
	}
	_, err = statement.Exec(movie.Name, movie.Language, movie.Budget, movie.Collection, name)
	if err != nil {
		panic(err)
	}
	return true
}

func DeleteMovie(name string) bool {
	present, _ := GetMovieByName(name)
	if !present {
		return false
	}
	query := `DELETE FROM MOVIES WHERE NAME=?`
	statement, err := DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error while querying [%v], err = %v", query, err)
		panic(err)
	}
	_, err = statement.Exec(name)
	if err != nil {
		fmt.Printf("Error while executing query [%v], err = %v", query, err)
		panic(err)
	}
	return present
}

func InsertMovie(movie types.Movies) {
	query := `INSERT INTO MOVIES(NAME,LANGUAGE,BUDGET,COLLECTION) VALUES(?,?,?,?)`
	statement, err := DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error while querying [%v], err = %v", query, err)
		panic(err)
	}
	_, err = statement.Exec(movie.Name, movie.Language, movie.Budget, movie.Collection)
	if err != nil {
		fmt.Printf("Error while executing query [%v], err = %v", query, err)
		panic(err)
	}
}
