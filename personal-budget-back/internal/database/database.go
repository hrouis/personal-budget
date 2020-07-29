package database

import (
	"context"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

//Connection database connection informtions
type Connection struct {
	client        driver.Client
	collectionMap map[string]driver.Collection
	database      driver.Database
}

//Config for the database.
type Config struct {
	URL      string
	Port     string
	User     string
	Password string
	DbName   string
}

//NewConnection creates a new database client
func NewConnection(ctx context.Context, config Config) *Connection {
	connection := new(Connection)
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.URL},
	})
	if err != nil {
		// Handle error
	}
	connection.client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.User, config.Password),
	})
	if err != nil {
		// Handle error
	}
	connection.database, err = connection.client.CreateDatabase(ctx, config.DbName, nil)
	if err != nil {
		// Handle error
	}
	connection.collectionMap = make(map[string]driver.Collection)
	return connection
}

//CreateCollection creates the collection.
func (connection *Connection) CreateCollection(ctx context.Context, collectionName string) *Connection {
	options := &driver.CreateCollectionOptions{
		//TODO add options here.
	}
	col, _ := connection.database.CreateCollection(ctx, collectionName, options)
	connection.collectionMap[collectionName] = col
	return connection
}

//ReadDocument fetches document from the database
func (connection *Connection) ReadDocument(ctx context.Context, collectionName string, documentID string) interface{} {
	document := new(interface{})
	connection.collectionMap[collectionName].ReadDocument(ctx, documentID, &document)
	return document
}

//UpdateDocument replace the document in the database
func (connection *Connection) UpdateDocument(ctx context.Context, collectionName string, documentID string, document interface{}) {
	connection.collectionMap[collectionName].ReplaceDocument(ctx, documentID, document)
}

//CreateDocument to a collection in the database.
func (connection *Connection) CreateDocument(ctx context.Context, collectionName string, document interface{}) {
	connection.collectionMap[collectionName].CreateDocument(ctx, document)
}

//RemoveDocument from a collection in the database.
func (connection *Connection) RemoveDocument(ctx context.Context, collectionName string, documentID string) {
	connection.collectionMap[collectionName].RemoveDocument(ctx, documentID)
}
