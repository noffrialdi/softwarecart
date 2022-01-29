package main

import (
	"fmt"
	"log"
	"softwarecart/helper"
	"softwarecart/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type responMessage struct {
	valid   bool
	message string
}

func main() {
	// TambahProduk("apel", 1)
	//  ShowData()
	//  DeleteData("apel")
	fmt.Println("Software library cart")
}

func TambahProduk(p string, q int) (responMessage, error) {
	helper.Mongo.Init()

	var produkname = p
	var kuantitas = q
	var tempInsert models.Shopingcart

	tempInsert.Produkname = produkname
	tempInsert.Kuantitas = kuantitas
	tempInsert.Created_at = time.Now()

	filter := bson.M{"produkname": produkname}

	var temRes responMessage
	log.Println(filter)
	resultData, _ := models.FindOne(filter)
	if resultData.Produkname != "" {
		kuantity_update := resultData.Kuantitas + tempInsert.Kuantitas

		updateContent := bson.M{"$set": bson.M{
			"kuantitas": kuantity_update,
		}}

		models.UpdateOne(filter, updateContent)

		temRes.valid = true
		temRes.message = fmt.Sprintf("Produk : %s, Quantity %d", produkname, kuantity_update)
	} else {
		models.InserData(tempInsert)
		temRes.valid = true
		temRes.message = fmt.Sprintf("Produk : %s, Quantity %d", produkname, q)
	}
	fmt.Println(temRes)
	return temRes, nil
}

func ShowData() ([]models.Shopingcart, error) {
	helper.Mongo.Init()
	filter := bson.M{}
	returnData, error := models.FindCart(filter, 0, 0)
	fmt.Println(returnData)
	return returnData, error
}

func DeleteData(p string) (responMessage, error) {
	filter := bson.M{"produk": p}
	get_respon := models.DeleteOne(filter)
	var respon responMessage
	if get_respon {
		respon.valid = true
		respon.message = fmt.Sprintf("%s Berhasil didelete.", p)
		return respon, nil
	}
	x := fmt.Errorf("failed or not found produkname %s", p)
	return respon, x
}
