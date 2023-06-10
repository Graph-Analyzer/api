package repository

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FullGraphItem struct {
	Nodes map[int64]FullGraphNodeItem
	Edges []FullGraphEdgeItem
}

type FullGraphNodeItem struct {
	Id       int64
	RouterId string
	Label    string
	Size     int64
}

type FullGraphEdgeItem struct {
	EdgeId       string
	From         int64
	FromRouterId string
	To           int64
	ToRouterId   string
	Weight       float64
}

func (r *graph) GetFullGraph() FullGraphItem {
	session := r.initReadSession()
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panic(err)
		}
	}(session)

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		resultNodes, err := getAllNodesFn(tx)
		if err != nil {
			return resultNodes, err
		}

		resultEdges, err := getAllEdgesFn(tx)
		if err != nil {
			return resultEdges, err
		}

		return FullGraphItem{
			Nodes: resultNodes.(map[int64]FullGraphNodeItem),
			Edges: resultEdges.([]FullGraphEdgeItem),
		}, err
	})
	if err != nil {
		log.Panic(err)
	}

	return result.(FullGraphItem)
}

func getAllNodesFn(tx neo4j.Transaction) (any, error) {
	query := `MATCH (n:Router)-[r:CONNECTS_TO]->(:Router)
		RETURN id(n), n.RouterID, n.Label, COUNT(r)`
	parameters := map[string]any{}

	records, err := tx.Run(query, parameters)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]FullGraphNodeItem)

	for records.Next() {
		record := records.Record()

		result[record.Values[0].(int64)] = FullGraphNodeItem{
			Id:       record.Values[0].(int64),
			RouterId: record.Values[1].(string),
			Label:    record.Values[2].(string),
			Size:     record.Values[3].(int64),
		}
	}

	if len(result) == 0 {
		return nil, err
	}

	return result, nil
}

func getAllEdgesFn(tx neo4j.Transaction) (any, error) {
	query := `MATCH (from:Router)-[r:CONNECTS_TO]->(to:Router)
		RETURN r.EdgeID, id(from), from.RouterID, id(to), to.RouterID, r.Weight`
	parameters := map[string]any{}

	records, err := tx.Run(query, parameters)
	if err != nil {
		return nil, err
	}

	result := make([]FullGraphEdgeItem, 0)

	for records.Next() {
		record := records.Record()

		result = append(result, FullGraphEdgeItem{
			EdgeId:       record.Values[0].(string),
			From:         record.Values[1].(int64),
			FromRouterId: record.Values[2].(string),
			To:           record.Values[3].(int64),
			ToRouterId:   record.Values[4].(string),
			Weight:       record.Values[5].(float64),
		})
	}

	return result, nil
}
