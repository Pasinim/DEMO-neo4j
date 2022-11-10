package funcs

import (
	"DEMO-neo4j/core"
	"DEMO-neo4j/utility"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type RepoDriver struct {
	drv neo4j.Driver
}

func New() *RepoDriver {
	db := utility.InitDriver()
	drv := RepoDriver{drv: db}
	return &drv
}

func (r *RepoDriver) GetItems() []core.Item {

	session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		res := make([]core.Item, 0)
		query := `MATCH (m:Item) RETURN m ORDER BY m.sku`
		result, err := tx.Run(query, nil)
		if err != nil {
			log.Fatal(err)
		}

		for result.Next() {
			props := result.Record().Values[0].(neo4j.Node).Props
			i := core.Item{
				Name: props["nome"].(string),
				Sku:  int(props["sku"].(int64)),
			}
			res = append(res, i)
		}
		return res, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return res.([]core.Item)
}

// GetItemFromSku
/* Restituisce l'item con sku passato come argomento*/
func (r *RepoDriver) GetItemFromSku(sku int) (core.Item, error) {
	session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)
	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (m:Item) WHERE m.sku = $wantedSku RETURN m"
		result, err := tx.Run(query, map[string]interface{}{"wantedSku": sku})
		if err != nil {
			log.Fatal(err)
		}
		if result.Next() {
			props := result.Record().Values[0].(neo4j.Node).Props
			retIt := core.Item{
				Name: props["nome"].(string),
				Sku:  sku,
			}
			return retIt, nil
		}
		return nil, nil
	})

	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		fmt.Print("RES NIL\n")
		return *new(core.Item), nil
	}
	return res.(core.Item), nil
}
