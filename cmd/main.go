package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/arimatakao/test-task-evo/internal/apiserver"
	_ "github.com/lib/pq"
)

var (
	user     = "postgres"
	password = "1234"
	dbname   = "testtask"
)

func pareseCsvToDB(filename string) error {
	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	insertScrpit := `INSERT INTO public.material(
		"TransactionId", "RequestId", "TerminalId", "PartnerObjectId", "AmountTotal", "AmountOriginal", "CommissionPS", "CommissionClient", "CommissionProvider", "DateInput", "DatePost", "Status", "PaymentType", "PaymentNumber", "ServiceId", "Service", "PayeeId", "PayeeName", "PayeeBankMfo", "PayeeBankAccount", "PaymentNarrative")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);`
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	isFirstRecord := true
	for _, record := range records {
		if isFirstRecord {
			isFirstRecord = false
			continue
		}
		_, err = db.Exec(insertScrpit, record[0], record[1], record[2], record[3], record[4], record[5], record[6],
			record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14], record[15],
			record[16], record[17], record[18], record[19], record[20])
		if err != nil {
			return err
		}
		fmt.Println("Added new record to DB", record[0])
	}
	fmt.Println("End of parsing")
	return nil
}

func main() {

	if len(os.Args) > 1 {
		if err := pareseCsvToDB(os.Args[1]); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	s := &apiserver.APIServer{}
	s.SetDB(db)
	if err = s.Start(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
