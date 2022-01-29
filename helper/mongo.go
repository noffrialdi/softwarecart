package helper

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbMongo struct {
	Conn *mongo.Database
	Err  error
}

func (_d *dbMongo) Init() {
	_d.Conn, _d.Err = _d.Connect()
}

// Connect Execute Connect to Mongo server
func (_d *dbMongo) Connect() (*mongo.Database, error) {
	connStr, dbname := _d.ConnStr()
	var clientOptions *options.ClientOptions

	// Set client options
	connStr = connStr + "&ssl=false"
	clientOptions = options.Client().ApplyURI(connStr).SetRetryWrites(false)

	// Connect to MongoDB
	dbTimeout := 10 * time.Second // Default

	ctx, _ := context.WithTimeout(context.Background(), dbTimeout)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB: ", connStr)

	return client.Database(dbname), nil
}

// ConnStr Generate connection string for mongo from Env vars
func (*dbMongo) ConnStr() (string, string) {

	// if username not empty / production
	var (
		// host     = mongoHost + ":" + mongoPort
		// user     = mongoUser
		dbname = "sirclo"
		// password = url.QueryEscape(mongoPas)
		// authdb   = "buddyku"
	)

	connStr := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	return connStr, dbname
}

func ConnectDB(tablename string) *mongo.Collection {
	return Mongo.Conn.Collection(tablename)
}

func ContextTimeout(t int) (c context.Context, cf context.CancelFunc) {

	if t == 0 {
		t, _ = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	}

	c, cf = context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	return
}

var Mongo = &dbMongo{}
