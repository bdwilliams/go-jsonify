go-jsonify
==========

Example Usage:

	import (  
		"database/sql"  
		_ "github.com/go-sql-driver/mysql"  
		"github.com/bdwilliams/go-jsonify/jsonify"  
		"fmt"  
	)
	
	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	if err != nil {
		panic(err.Error())
	}
	
	rows, err := con.Query("select id, something_else from table")
	if err != nil {
		panic(err.Error())
	}
	
	defer rows.Close()
	defer con.Close()
	
	fmt.Println(jsonify.Jsonify(rows))
