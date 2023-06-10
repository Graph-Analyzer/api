package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type DensityItem struct {
	Value float64
}

func (r *graph) GetDensity() DensityItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	graph := r.GetLatestGraphName()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		return getDensityFn(tx, graph)
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(DensityItem)
}

func getDensityFn(tx neo4j.Transaction, graph GraphItem) (any, error) {
	query := `CALL gds.graph.list($graphName)
		YIELD density
		RETURN density`
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

	return DensityItem{
		Value: record.Values[0].(float64),
	}, nil
}
