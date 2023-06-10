package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectedItem struct {
	ComponentCount int64
}

func (r *graph) GetConnected() ConnectedItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	graph := r.GetLatestGraphName()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		return getConnectedFn(tx, graph)
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(ConnectedItem)
}

func getConnectedFn(tx neo4j.Transaction, graph GraphItem) (any, error) {
	query := `CALL gds.wcc.stats($graphName)
		YIELD componentCount
		RETURN componentCount`
	parameters := map[string]any{
		"graphName": graph.Name,
	}

	records, err := tx.Run(query, parameters)
	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return ConnectedItem{
		ComponentCount: record.Values[0].(int64),
	}, nil
}
