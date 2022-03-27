package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"os"
)

func FunctionWithDefer() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	//defer can be anywhere
	//but it will get executed
	//once the entire function is done
	defer f.Close()
}

func DefersRolePlayInDBHandle() {
	//DoSomeInserts(ctx, db, columnValue1, columnValue2 string)
}
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO Foo (val) values $1", value1)
	if err != nil {
		return err
	}
	//use tx to do more DB inserts here
	return nil
}

func DefersRolePlayInFileHandle() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		data := make([]byte, 2048)
		count, err := f.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data[:count])
	}
	defer closer()
}
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return file, func() { file.Close() }, nil
}
