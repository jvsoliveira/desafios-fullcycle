package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jvsoliveira/desafios-fullcycle/desafio-go/adapter/factory"
	"github.com/jvsoliveira/desafios-fullcycle/desafio-go/adapter/repository/fixture"
	"github.com/jvsoliveira/desafios-fullcycle/desafio-go/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Processando nova transação...")
	db, err := sql.Open("sqlite3", "banco.db")
	if err != nil {
		log.Fatal(err)
	}
	db = fixture.CreateDb(db)

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	usecase := process_transaction.NewProcessTransaction(repository)

	input := process_transaction.TransactionDtoInput{
		ID:        "23",
		AccountID: "5753",
		Amount:    3600,
	}
	fmt.Printf("Transação : [ID AccountID Amount] -=> %v\n", input)

	err = usecase.Execute(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Transação registrada com sucesso em banco.db!")

}
