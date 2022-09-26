package productHandlers

import (
	"encoding/json"
	"net/http"
	"project/pkg/models"
	"project/pkg/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	b := CreateProduct.CreateProduct()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
