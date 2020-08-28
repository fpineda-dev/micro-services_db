package product

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"main.go/src/entities"

	"github.com/gorilla/mux"
	"main.go/src/config"

	"main.go/src/models"
)

func FindAll(response http.ResponseWriter, request *http.Request) {

	//Manejar cierre una vez que se ejecuta la busqueda
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}

	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()

}

func Search(response http.ResponseWriter, request *http.Request) {

	//Manejar cierre una vez que se ejecuta la busqueda
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	vars := mux.Vars(request)
	Keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search(Keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}

	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
}

func SearchPrices(response http.ResponseWriter, request *http.Request) {

	//Manejar cierre una vez que se ejecuta la busqueda
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	vars := mux.Vars(request)
	smin := vars["min"]
	smax := vars["max"]
	min, _ := strconv.ParseFloat(smin, 64)
	max, _ := strconv.ParseFloat(smax, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchPrices(min, max)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}

	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
}

func Create(response http.ResponseWriter, request *http.Request) {
	//Manejar cierre una vez que se ejecuta la busqueda
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}

	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		RowsAffected, err2 := productModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"RowsAffected": RowsAffected,
			})
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*Funcion para manejar los tiempos de cancelacion*/
func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
