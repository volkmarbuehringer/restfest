# restfest
rest services with postgresql database, boiler code is generated with go generate

- Works with Postgresql-Database
- only standard-Libraries - no SQLX
- generates Rest-Services for every Table in PG-Schema with go generate

Typesafe database-Code with generated structues,sql and scanners:

scanning with individual generated structs:

	rows, err := db.DB.Query("select * from gutschein")
	if err != nil {
		senderErr(w, err)
		return
	}
	
	defer rows.Close()
	for anz := 0; rows.Next(); anz++ {
		arr, ts := gener.ScannerTGutschein()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		if ts.Auf_partner.Valid {
			fmt.Printf("%d %s \n", ts.Auf_id.Int64, ts.Auf_partner.String)
		} else {
			fmt.Printf("%d null \n", ts.Auf_id.Int64)
		}
		//fmt.Printf("%v %T \n", ts, ts)
	}
	
	
generic scanning with interface wrapped structs and generated SQL:


The resulting code is type-safe:
go generate
go run/main*

A build mismatch againt a master-database produces errors at compile-time not only runtime !

Generic code adapts to the new schema after every go generate, individual code is checked at compile time
