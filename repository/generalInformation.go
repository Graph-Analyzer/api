package repository

import (
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GeneralInformationItem struct {
	Name      string
	CreatedAt time.Time
}

func (r *graph) GetGeneralInformation() GeneralInformationItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	graph := r.GetLatestGraphName()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		return getGeneralInformationFn(tx, graph)
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(GeneralInformationItem)
}

func getGeneralInformationFn(tx neo4j.Transaction, graph GraphItem) (any, error) {
	query := `CALL gds.graph.list($graphName)
		YIELD graphName, creationTime
		RETURN graphName AS name, creationTime AS createdAt`
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

	return GeneralInformationItem{
		Name:      record.Values[0].(string),
		CreatedAt: record.Values[1].(time.Time),
	}, nil
}
