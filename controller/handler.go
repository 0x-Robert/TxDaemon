package controller

import (
	"context"
	"log"
	"net/http"

	"block-go-web/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

type User2 struct {
	Name  string
	Email string
	Age   int
}

type Block struct {
	BlockNumber string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// bks := make([]model.Block, 0)
	// for rows.Next() {
	// 	bk := model.Block{}
	// 	err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	bks = append(bks, bk)
	// }
	// Bk := AllBlock()

	// BlockNumber := AllBlock()
	// BlockNumber := model.Block{BlockNumber: AllBlock()}
	// user := User2{Name: "changbeom", Email: "changbeom@naver.com", Age: 23}
	block := Block{BlockNumber: AllBlock()}
	// config.TPL.ExecuteTemplate(w, "index.gohtml", Bks)
	config.TPL.ExecuteTemplate(w, "index.gohtml", block)
}

func AllBlock() string {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(header.Number.String()) // 5671744
	return header.Number.String()
}
