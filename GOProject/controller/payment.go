//Payment controller
package controller

import (
	"GOProject/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) GetPayment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u.Name)

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(cart.TotalPrice)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) PostPayment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	fmt.Println(id)
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	cartproducts := model.CartProduct{}
	json.NewDecoder(r.Body).Decode(&cartproducts)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u.Name)

	cart := model.Cart{}
	if err := uc.session.DB("go-web-dev-db").C("carts").Find(bson.M{"uname": u.Name}).One(&cart); err != nil {
		fmt.Println("Error we arer in")
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	productNotFound, flag := uc.CheckProductsBeforeOrder(cart.CartProducts)
	if flag == 0 {
		fmt.Println("Quantities insufficient:", productNotFound)
	} else {
		fmt.Println("Redirecting to Place Orders")
		//w.Header().Set("Content-Type", "application/json")
		//r.Method == "POST"
		http.Redirect(w, r, "/user/"+id+"/order", http.StatusMovedPermanently)

		//http.Redirect(w, r, "/user/:id/order", http.StatusSeeOther)

		//uc.PlaceOrder()
	}

}

func (uc UserController) CheckProductsBeforeOrder(slice []model.CartProduct) (string, int) {
	flag := 1
	var productNotFound string
	for i, _ := range slice {
		product := model.Product{}
		if err := uc.session.DB("go-web-dev-db").C("products").Find(bson.M{"pname": slice[i].ProductName}).One(&product); err != nil {
			fmt.Println("Error we arer in")
			fmt.Println(err)
			//w.WriteHeader(404)
		}
		if product.ProductQty >= slice[i].ProductQty {
		} else {
			flag = 0 //change here
			productNotFound = product.ProductName
			return productNotFound, flag
		}

	}
	return "", flag
}
