package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AverageClusteringCoefficientItem struct {
	Value float64
}

func (r *graph) GetAverageClusteringCoefficient() AverageClusteringCoefficientItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	graph := r.GetLatestGraphName()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		return getAverageClusteringCoefficientFn(tx, graph)
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(AverageClusteringCoefficientItem)
}

func getAverageClusteringCoefficientFn(tx neo4j.Transaction, graph GraphItem) (any, error) {
	query := `CALL gds.localClusteringCoefficient.stats($graphName)
		YIELD averageClusteringCoefficient
		RETURN averageClusteringCoefficient`
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

	return AverageClusteringCoefficientItem{
		Value: record.Values[0].(float64),
	}, nil
}
