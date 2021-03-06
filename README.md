# restfest
rest services with postgresql database, boiler code is generated with go generate

- Works with Postgresql-Database
- only standard-Libraries - no SQLX
- works now with pgx-Driver version 3 (yet unreleased )
- generates Rest-Services for every Table in PG-Schema with go generate

- Uses Tables, Views, Types and Functions in Postgres as source for struct-generation
- with pgx now support of nested structs and arrays

you can use array_agg and nested types in Postgres which are supported in the generated structs

with the extendable type interface of pgx can the nested and aggregated data in Postges be scanned into the go-structs and then displayed as json or used programatically


Typesafe database-Code with generated structues,sql and scanners:

scanning with individual generated structs:

<!--  -->
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
<!--  -->

generic scanning with interface wrapped structs and generated SQL:

<!--  -->
     func rowScanner(tab string, rows *sql.Rows, len int) (stru []interface{}, err error) {
	t := make([]interface{}, 0)
	fun := gener.ScannerFunMap[tab]

	for anz := 0; rows.Next(); anz++ {
		arr, ts := fun()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		if len == 1 {
			stru = []interface{}{ts}
			return
		}
		t = append(t, ts)
	}
	stru = []interface{}{&t}
	return
    }
<!--  -->


The resulting code is type-safe:
- go generate
- go run main/*

A build mismatch against a master-database produces errors at compile-time not only at runtime !

Generic code adapts to the new schema after every go generate, individual code is checked at compile time

There is special meaning in some column names which is used in update and insert sql, also primary keys are assumed
to be generated by the database:
_cr_date,
_upd_date,
_upd_uid,
_cr_uid,

The generation is done with templates in the generator file.

It contains also the most extensive set of scanners, valuers, marshallers and unmarshallers for postgresql.

The null-handling of strings can be changed between sql.Nullstring and converting null to empty string.


Flexible REST-Services with functions:

```
CREATE OR REPLACE FUNCTION get_weburl(
p_id integer,
	p_start integer,
  p_end integer,
  p_len integer,
	p_lala character varying)
    RETURNS setof weburl
    LANGUAGE 'sql'
    COST 100.0
    ROWS 1000.0
AS $function$

SELECT * FROM weburl WHERE id >= coalesce(p_id,id) and zusatz between coalesce(p_start,zusatz) and coalesce(p_end,zusatz) and url like coalesce(p_lala,'')||'%'
order by id
limit coalesce(p_len,100);

$function$;
```



After go generate you can use the functions for Rest-Services with variables:

http://localhost:8080/test/service/get_weburl_173826?p_start=80&p_end=100&p_len=100&p_lala=www&p_id=7000

Unset variables are null and can be defaulted in your function.
The combination of views and functions allows flexible data selections
