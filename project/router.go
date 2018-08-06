package project

import (
	"encoding/json"
	"net/http"

	"github.com/hdnha11/go-api-example/entity"
	"github.com/julienschmidt/httprouter"
)

// InitRouters init project routers
func InitRouters(router *httprouter.Router, service *Service) {
	router.GET("/api/v1/projects", index(service))
	router.GET("/api/v1/projects/:id", get(service))
	router.POST("/api/v1/projects", save(service))
	router.PUT("/api/v1/projects/:id", update(service))
	router.DELETE("/api/v1/projects/:id", remove(service))
}

func index(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		query := req.URL.Query().Get("name")

		var ps []*entity.Project
		var err error

		if query == "" {
			ps, err = service.FindAll()
		} else {
			ps, err = service.FindByName(query)
		}

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(ps); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func get(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id, err := entity.StringToID(params.ByName("id"))

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		p, err := service.Find(id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}

func save(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var p *entity.Project

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := service.Save(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func update(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var p *entity.Project
		id, err := entity.StringToID(params.ByName("id"))

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := service.Update(entity.ID(id), p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func remove(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id, err := entity.StringToID(params.ByName("id"))

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		if err := service.Delete(id); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
	}
}
