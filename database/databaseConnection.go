package database
import(
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)
func DBinstance() *mongo.Client{
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	mongoDb:=os.Getenv("MONGODB_URL")
	client,err:=mongo.NewClient(options.Client().ApplyURI(mongoDb))
	if err!= nil{
		log.Fatal(err)
	}
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err=client.Connect(ctx)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB:")
	return client
}
var Client *mongo.client=DBinstance()
func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	var collection *mongo.Collection=client.Database("cluster0").Collection(collectionName)
	return collection
}