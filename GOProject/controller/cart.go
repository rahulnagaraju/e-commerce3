//-------------------------------------//
//Cart Controller

package controller

import (
	"encoding/json"
	"fmt"

	//"C:/Users/Dell/Desktop/GOProject/model"
	"GOProject/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func (uc UserController) GetAllCarts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//id := p.ByName("id")

	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound) // 404
	// 	return
	// }

	// oid := bson.ObjectIdHex(id)

	carts := []model.Cart{}

	if err := uc.session.DB("go-web-dev-db").C("carts").Find(nil).All(&carts); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(carts)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	cart := model.Cart{}
	json.NewDecoder(r.Body).Decode(&cart)
	cart.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("carts").Insert(cart)

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetCartUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u_id := p.ByName("id")

	if !bson.IsObjectIdHex(u_id) {
		fmt.Println("Inside get cart user bson error")
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(u_id)

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

	uj, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u_id := p.ByName("id")

	if !bson.IsObjectIdHex(u_id) {
		fmt.Println("Inside get cart user bson error")
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(u_id)

	u := model.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u.Name)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("carts").Remove(bson.M{"uname": u.Name}); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
