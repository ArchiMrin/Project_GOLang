package services

import (
	"context"
	"fmt"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/entities"
	"github.com/ArchiMrin/Project_GOLang/eCOMM/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection //a data called Product carrying the context of collection instance
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct { //instantiating the product service
	return &ProductService{Product: collection}
}

func (p *ProductService) Insert(product *entities.Product) (string, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.Product.InsertOne(context.Background(), product) //p.Product gives the Collection & then we call the operation InsertOne

	if err != nil {
		return "", err
	} else {
		return "Record Inserted", nil
	}

}
func (p *ProductService) GetProducts() ([]*entities.Product, error) {
	result, err := p.Product.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		//do something
		fmt.Println(result)
		//build the array of products for the cursor that we received.
		var products []*entities.Product
		for result.Next(context.TODO()) {
			product := &entities.Product{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}

}
func (p *ProductService) GetProductById(id string) (*entities.Product, error) {

	filter := bson.M{"_id": id}

	var product entities.Product

	err := p.Product.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {

		if err == mongo.ErrNoDocuments {

			return nil, nil

		}

		return nil, err

	}

	return &product, nil

}
