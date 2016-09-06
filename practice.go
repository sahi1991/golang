package main

import (
	"fmt"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func main(){
db,err := sql.Open("conn", "server=localhost;user id=SAHITHYA-PC\\SAHITHYA;password=pasword;")
if err  != nil {
fmt.Println("  Error open db:", err.Error())
}

defer db.Close()
}
