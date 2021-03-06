package main

import
//"encoding/json"

(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/jackc/pgx"
	httpstat "github.com/tcnksm/go-httpstat"
)

var dbx *pgx.Conn

var zahler int = 0

func main() {

	tr := &http.Transport{
		//	MaxIdleConns:       10,
		//	IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		DisableKeepAlives:  true,
	}
	client := &http.Client{Transport: tr,
		Timeout: 1 * time.Second}
	c := make(chan int)
	for i := 0; i < 20; i++ {
		go getter(client, c)
	}

	for i := 0; i < 1000; i++ {
		rows, err := dbx.Query(fmt.Sprintf("SELECT %s from %s.%s", os.Args[2], os.Getenv("PGSCHEMA"), os.Args[1]))
		checkErr(err)
		// iterate over each row
		for rows.Next() {
			var adsid int
			err = rows.Scan(&adsid)
			checkErr(err)
			c <- adsid
		}
	}

	close(c)
	//time.Sleep(1000 * time.Second)
}

func getter(client *http.Client, c chan int) {

	//stop := timer("get" + strconv.Itoa(i))
	//defer stop()
	for ga := range c {
		req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/test/service/%s/%d", os.Args[3], os.Args[1], ga), nil)

		checkErr(err)

		// Create a httpstat powered context
		var result httpstat.Result
		ctx := httpstat.WithHTTPStat(req.Context(), &result)
		req = req.WithContext(ctx)
		resp, err := client.Do(req)
		checkErr(err)

		//t := []dbstructs.Weburl{}
		t := make(map[string]interface{})
		checkErr(json.NewDecoder(resp.Body).Decode(&t))

		//_, err = ioutil.ReadAll(resp.Body)
		checkErr(err)
		resp.Body.Close()

		//fmt.Println("sync", <-c)
		zahler++
		log.Printf("Server processing: %d %d ms %d", zahler, int(result.ServerProcessing/time.Millisecond), ga)
		//		log.Printf("%d", ga)
		/*
			// Show the results
			log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
			log.Printf("TCP connection: %d ms", int(result.TCPConnection/time.Millisecond))
			log.Printf("TLS handshake: %d ms", int(result.TLSHandshake/time.Millisecond))
			log.Printf("Server processing: %d ms", int(result.ServerProcessing/time.Millisecond))
			log.Printf("Content transfer: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
		*/
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func timer(name string) func() {
	t := time.Now()
	log.Println(name, "start", t)
	return func() {
		d := time.Since(t)
		log.Println(name, "took", d)
	}
}

func init() {
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		logrus.Fatal("DB parse", err)
	}
	connConfig.LogLevel = pgx.LogLevelWarn

	if dbx, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB connect", err)
	}

}
