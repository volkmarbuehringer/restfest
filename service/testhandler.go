package service

import (
	"net/http"
	"restfest/db"
)

func testHandler(w http.ResponseWriter, r *http.Request) {

	//	vars := mux.Vars(r)
	//	var cursor gener.Gutschein

	rows, err := db.DB.Query("select * from gutschein")
	if err != nil {
		senderErr(w, err)
		return
	}

	defer rows.Close()
	for anz := 0; rows.Next(); anz++ {
		/*arr, ts := gener.ScannerTGutschein()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		if ts.Auf_partner.Valid {
			fmt.Printf("%d %s \n", ts.Auf_id.Int64, ts.Auf_partner.String)
		} else {
			fmt.Printf("%d null \n", ts.Auf_id.Int64)
		}*/
		//fmt.Printf("%v %T \n", ts, ts)
	}
}
