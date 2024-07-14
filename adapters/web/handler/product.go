package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andrerampanelli/hexagonal-arch/adapters/dto"
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service interfaces.ProductServiceInterface) {
	r.Handle("/product", n.With(
		negroni.Wrap(listProducts(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}", n.With(
		negroni.Wrap(updateProduct(service)),
	)).Methods("PATCH", "OPTIONS")
}

func listProducts(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		products, err := service.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error listing products"))
			return
		}

		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding products"))
			return
		}
	})
}

func getProduct(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError("Product not found"))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding product"))
			return
		}
	})
}

func createProduct(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dto dto.CreateProductDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error decoding product"))
			return
		}

		product, err := service.Create(dto.Name, dto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error creating product"))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding product"))
			return
		}
	})
}

func enableProduct(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError("Product not found"))
			return
		}

		product, err = service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error enabling product"))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding product"))
			return
		}
	})
}

func disableProduct(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError("Product not found"))
			return
		}

		product, err = service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error disabling product"))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding product"))
			return
		}
	})
}

func updateProduct(service interfaces.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError("Product not found"))
			return
		}

		var updateDto dto.UpdateProductDto
		err = json.NewDecoder(r.Body).Decode(&updateDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error decoding product"))
			return
		}

		productDto := dto.NewProduct()
		productDto.ID = product.GetId()
		productDto.Name = updateDto.Name
		productDto.Price = updateDto.Price
		productDto.Status = updateDto.Status
		product, err = productDto.Bind(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error binding product"))
			return
		}

		product, err = service.Save(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error modifying product"))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Error encoding product"))
			return
		}
	})
}
