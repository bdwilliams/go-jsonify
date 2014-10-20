package jsonify

import (
	"encoding/json"
	"strconv"
	"strings"
	"database/sql"
)

func Jsonify(rows *sql.Rows) ([]string) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	c := 0
	results := make(map[string]interface{})
	data := []string{}

	for rows.Next() {
		if c > 0 {
			data = append(data, ",")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, value := range values {
			switch value.(type) {
				case nil:
					results[columns[i]] = nil

				case []byte:
					s := string(value.([]byte))
					x, err := strconv.Atoi(s)

					if err != nil {
						results[columns[i]] = s
					} else {
						results[columns[i]] = x
					}


				default:
					results[columns[i]] = value
			}
		}

		b, _ := json.Marshal(results)
		data = append(data, strings.TrimSpace(string(b)))
		c++
	}

	return data
}
