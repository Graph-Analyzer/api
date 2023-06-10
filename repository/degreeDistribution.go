package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type DegreeDistributionItem struct {
	Values []DegreeDistributionValueItem
}

type DegreeDistributionValueItem struct {
	Degree int64
	Amount int64
}

func (r *graph) GetDegreeDistribution() DegreeDistributionItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	graph := r.GetLatestGraphName()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		result, err := getDegreeDistributionFn(tx, graph)
		if err != nil {
			return result, err
		}

		return DegreeDistributionItem{
			Values: result.([]DegreeDistributionValueItem),
		}, err
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(DegreeDistributionItem)
}

func getDegreeDistributionFn(tx neo4j.Transaction, graph GraphItem) (any, error) {
	query := `CALL gds.degree.stream($graphName)
		YIELD score
		RETURN score AS degree, COUNT(*) as amount
		ORDER BY degree ASC`
	parameters := map[string]any{
		"graphName": graph.Name,
	}

	records, err := tx.Run(query, parameters)
	if err != nil {
		return nil, err
	}

	result := make([]DegreeDistributionValueItem, 0)

	for records.Next() {
		record := records.Record()

		result = append(result, DegreeDistributionValueItem{
			Degree: int64(record.Values[0].(float64)),
			Amount: record.Values[1].(int64),
		})
	}

	return result, nil
}
