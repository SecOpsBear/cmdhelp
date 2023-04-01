package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// OpenDatabaseConn opens database connection and ping's the database and returns db connection
func OpenDatabaseConn(dbfile string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CheckTableExist check if the table "helpcommnds" exist in the master database
func CheckTableExist() (bool, error) {
	var count int
	query := `SELECT count(*) FROM sqlite_master WHERE name ='helpcommands' and type='table';`
	if err := db.QueryRow(query).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("thers is no helpcommands table in the database")
		}
	}
	if count == 0 {
		return false, errors.New("thers is no helpcommands table in the database")
	}
	return true, nil
}

// CreateTable creates helpcommands table in the database
func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS helpcommands (
		"idCommand" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"tool" TEXT,
		"gist" TEXT,
		"toolCommand" TEXT

	);`

	// CheckTableExist()
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
		// return nil, err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Successfully Create helpcommands table.")
}

// InsertNote inserts tool command with gist in the helpcommands table
func InsertCommand(tool string, gist string, toolCommand string) {
	insertNoteSQL := `INSERT INTO helpcommands (tool, gist, toolCommand) 
	VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(strings.ToLower(tool), gist, toolCommand)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully add the Command")
}

// DisplayAllNotes prints all the rows in the helpcommands table
func DisplayAllNotes() []ToolsData {
	row, err := db.Query("SELECT * FROM helpcommands ORDER BY tool")

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()
	var helpCmds []ToolsData
	for row.Next() {
		var td ToolsData
		row.Scan(&td.IdCommand, &td.Tool, &td.Gist, &td.ToolCommand)
		helpCmds = append(helpCmds, td)
	}

	return helpCmds
}

// DisplayCatNotes returns a slice of ToolsData after performing search based on the string provided
func SearchForTool(toolSearch string) []ToolsData {
	toolSearch = strings.TrimSpace(toolSearch)
	toolSearch = strings.ToLower(toolSearch)
	row, err := db.Query("SELECT * FROM helpcommands Where (tool = ?) ORDER BY tool", toolSearch)

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	var searchData []ToolsData
	for row.Next() {

		var td ToolsData
		row.Scan(&td.IdCommand, &td.Tool, &td.Gist, &td.ToolCommand)
		searchData = append(searchData, td)
	}

	return searchData
}

// DisplayCatNotes returns a slice of ToolsData after performing search based on the string provided
func SearchForToolLike(toolSearch string) []ToolsData {
	toolSearch = strings.TrimSpace(toolSearch)
	toolSearch = strings.ToLower(toolSearch)
	qStatement := "SELECT * FROM helpcommands Where (tool like '%" + toolSearch + "%') ORDER BY tool"
	row, err := db.Query(qStatement)

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	var searchData []ToolsData
	for row.Next() {

		var td ToolsData
		row.Scan(&td.IdCommand, &td.Tool, &td.Gist, &td.ToolCommand)
		searchData = append(searchData, td)
	}

	return searchData
}

// ExtractToolsList returns a slice containg the list of tools
func ExtractToolsList() []string {
	row, err := db.Query("select tool from helpcommands ORDER BY tool")

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()
	var listtools []string

	for row.Next() {
		var tool string
		row.Scan(&tool)
		listtools = append(listtools, tool)
	}
	uniqueList := Unique(listtools)
	return uniqueList
}

// DisplaySelectedCommand retuns slice to ToolsData based on search string
func DisplaySelectedCommand(toolSearch string) *[]ToolsData {
	var toolsData []ToolsData
	row, err := db.Query("SELECT * FROM helpcommands Where (tool = ?) ORDER BY tool", toolSearch)

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var rowData ToolsData
		row.Scan(&rowData.IdCommand, &rowData.Tool, &rowData.Gist, &rowData.ToolCommand)
		toolsData = append(toolsData, rowData)
	}

	return &toolsData
}

// DeleteCommand deletes the row from the helpcommands table based on the row id provided
func DeleteCommand(index int) {
	stmt, err := db.Prepare("DELETE FROM helpcommands Where (idCommand = ?)")

	if err != nil {
		log.Fatalln(err)
	}

	defer stmt.Close()
	res, err := stmt.Exec(index)
	if err != nil {
		log.Fatalln(err)
	}
	af, err := res.RowsAffected()

	if err != nil {
		log.Fatalln(err)
	}

	if af == 1 {
		fmt.Println("Succesfully Deleted")
	}
}
