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

// ContainsItem
// RESTITUISCE TRUE SE ITEM (NOME, SKU) È PRESENTE NEL DB, FALSO ALTRIMENTI/**
func (r *RepoDriver) ContainsItem(nome string, sku int) bool {
	sessione := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, err := sessione.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := `MATCH (m:Item) WHERE m.nome = $wantedNome AND m.sku = $wantedSku RETURN m`
		res, err := tx.Run(query, map[string]interface{}{"wantedNome": nome, "wantedSku": sku})
		if err != nil {
			log.Fatal(err)
		}
		res.Next()
		if res.Record() == nil {
			return false, nil
		}
		props := res.Record().Values[0].(neo4j.Node).Props
		if props["nome"].(string) != nome {
			return false, nil
		}
		if int(props["sku"].(int64)) != sku {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return result.(bool)
}

/*
*
Se non è presente inserisce il nodo Item all'interno del dd, altrimenti non fa nulla.
Al termine di ogni inserimento avvenuto con successo verifica anche che sia stato creato uno e un solo nood
*/
func (r *RepoDriver) InsertItem(nome string, sku int) bool {
	session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(session)
	//if r.ContainsItem(nome, sku) {
	//	return false
	//}
	ok, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		//ora che ho sku come chiave primaria posso direttamente fare il merge
		query := `MERGE (i:Item{nome: $newNome, sku: $newSku})`
		_, err := tx.Run(query, map[string]interface{}{"newNome": nome, "newSku": sku})
		if err != nil {
			log.Fatal(err)
		}
		return true, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return ok.(bool)
}

func (r *RepoDriver) GetItems() []core.Item {

	session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Fatal(err)
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
