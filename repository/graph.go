package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GraphRepository interface {
	GetAverageClusteringCoefficient() AverageClusteringCoefficientItem
	GetConnected() ConnectedItem
	GetDegreeDistribution() DegreeDistributionItem
	GetDensity() DensityItem
	GetFullGraph() FullGraphItem
	GetGeneralInformation() GeneralInformationItem
	GetLatestGraphName() GraphItem
}

type graph struct {
	driver neo4j.Driver
}

func NewGraphRepository(driver neo4j.Driver) GraphRepository {
	return &graph{
		driver: driver,
	}
}

type GraphItem struct {
	Name string
}

func (r *graph) GetLatestGraphName() GraphItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	result, err := session.ReadTransaction(getLatestGraphNameFn)
	if err != nil {
		log.Panic(err)
	}

	return result.(GraphItem)
}

func (r *graph) initReadSession() neo4j.Session {
	session := r.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})

	return session
}

func getLatestGraphNameFn(tx neo4j.Transaction) (any, error) {
	query := `CALL gds.graph.list()
		YIELD graphName, creationTime
		RETURN graphName
		ORDER BY creationTime DESC
		LIMIT 1`
	parameters := map[string]any{}

	records, err := tx.Run(query, parameters)
	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return GraphItem{
		Name: record.Values[0].(string),
	}, nil
}
