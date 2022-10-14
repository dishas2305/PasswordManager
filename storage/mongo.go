package storage

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MONGO_DB *mongo.Database

const (
	readPreference           = "secondaryPreferred"
	connectionStringTemplate = "mongodb://%s:%s@%s/%s?retryWrites=false&tls=true&replicaSet=rs0&readpreference=%s"
)

func ConnectMongoDB() *mongo.Database {
	fmt.Println("---get env", os.Getenv("ENV"))

	var client *mongo.Client
	if os.Getenv("ENV") == "local" {
		_client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

		fmt.Println("_client", _client)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		client = _client
	} else {
		connectionURI := fmt.Sprintf(connectionStringTemplate, os.Getenv("DB_USER_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DOCDB_CLUSTER_ENDPOINT"), os.Getenv("DB_NAME"), readPreference)
		tlsConfig, err := getCustomTLSConfig(os.Getenv("RDS_DOCDB_CA_FILE_PATH"))
		if err != nil {
			log.Fatal(err)
		}
		_client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI).SetTLSConfig(tlsConfig))
		if err != nil {
			log.Fatal(err)
		}
		client = _client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		logger.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("MDB Connected!")
	}

	// Connect to the database
	db := client.Database(os.Getenv("MONGO_DB_NAME"))

	fmt.Println("db-------------------------------->", db)
	return db
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}
