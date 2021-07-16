//----------------------------------------------//
//Product controller

package controller

import (
	"GOProject/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) GetAllProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//id := p.ByName("id")

	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound) // 404
	// 	return
	// }

	// oid := bson.ObjectIdHex(id)

	products := []model.Product{}

	if err := uc.session.DB("go-web-dev-db").C("products").Find(nil).All(&products); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	product := model.Product{}

	if err := uc.session.DB("go-web-dev-db").C("products").FindId(oid).One(&product); err != nil {
		w.WriteHeader(404)
		return
	}

	fmt.Println(product.ProductName)

	uj, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	product := model.Product{}

	json.NewDecoder(r.Body).Decode(&product)

	product.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("products").Insert(product)

	uj, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("products").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
