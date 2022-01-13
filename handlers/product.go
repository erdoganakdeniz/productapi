package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"productapi/databases"
	"productapi/models"
)
var ctx context.Context
var collection *mongo.Collection
var database *mongo.Database
func InsertProduct(w http.ResponseWriter, r *http.Request) {
	database,collection=databases.DB()
	w.Header().Set("Content-Type","application/json")
	log.Println("Serving :", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Error :", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	d,err:=ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Error : ",http.StatusInternalServerError)
		return
	}
	var product models.Product
	err=json.Unmarshal(d,&product)
	if err != nil {
		log.Println(err)
		http.Error(w,"Error : ",http.StatusBadRequest)
		return
	}
	result,err:=collection.InsertOne(ctx,&product)
	if err != nil {
		log.Println(err)
	}
	res,_:=json.Marshal(result)
	w.Write(res)



}
