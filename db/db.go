package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Used to import the side effects of a package. Allows for SQlit3 Driver to be known
	"github.com/rs/zerolog/log"
)

type fooStruct struct {
	id string
}


// Good enough table SCHEMA for now

/*
rank_id is a comma separated list of the foreign keys associated with ranking table

num_ranks is the number of times a specific song has been ranked.
Useful in case a special day is made where previous songs can be ranked again 
(ex. Christmas songs on christmas, where instead of three it's every submitters favorite christmas song)
*/
const createMusicTable string = `CREATE TABLE IF NOT EXISTS music (
	id INTEGER NOT NULL PRIMARY KEY,
	insert_date DATETIME NOT NULL,
	songURL TEXT,
	genre TEXT,
	subgenre TEXT,
	description TEXT,
	submitterID INTEGER FOREIGN KEY,
	rank_id TEXT
	num_ranks INTEGER
)`
/*
Ranking can include multiple songs that have been ranked, thus it will be
a string that has specific format easy for tokenization, with each token
being a reference to a music ID
*/

const createRankingTable string = `CREATE TABLE IF NOT EXISTS ranking (
	id INTEGER NOT NULL PRIMARY KEY,
	date_ranked DATETIME NOT NULL,
	ranking TEXT
)`


type MusicDB struct {
	db *sql.DB
}


func CreateSQLiteStorage() *MusicDB{
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	db.Exec(createMusicTable)
	db.Exec(createRankingTable)
	mdb := MusicDB{db}
	return &mdb
}

func (mdb MusicDB) createNewCurator(musicEntry MusicEntry){

}

func (mdb MusicDB) insertNewSong(musicEntry MusicEntry){

}

func (mdb MusicDB) getTodaysRanking(){

}

func (mdb MusicDB) updateTodaysRanking(){
	
}

func (mdb MusicDB) getTodaysMusic(){

}

func (mdb MusicDB) getCalendarsMusic(){

}



func NewSQLiteStorageFoo() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS foo(
		fid INTEGER
	)`)
	fd, err2 := db.Exec("INSERT INTO foo VALUES(?);", 'Q')
	if err2 != nil {
		log.Error().Msg(err2.Error())
	}
	fd.LastInsertId()
	var ex1 *sql.Row = db.QueryRow("SELECT * FROM foo WHERE fid=?", 'Q')
	var result fooStruct
	errr := ex1.Scan(&result.id)
	fmt.Println(errr)
	fmt.Println(result)
}
