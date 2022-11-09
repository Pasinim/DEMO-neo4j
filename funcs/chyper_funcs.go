package funcs

import (
	"DEMO-neo4j/core"
	"DEMO-neo4j/utility"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type RepoDriver struct {
	drv neo4j.Driver
}

func New() *RepoDriver {
	db := utility.InitDriver()
	drv := RepoDriver{drv: *db}
	return &drv
}

func (r *RepoDriver) GetItems() []core.Item {

	session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		res := make([]core.Item, 0)
		query := `MATCH (m:Item) RETURN m`
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

/**
* Restituisce l'item con sku passato come argomento
 */
//func (r *RepoDriver) GetItemFromSku (sku int) (core.Item, error){
//	//todo
//	return  nil, nil
//}
