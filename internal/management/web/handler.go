package web

import (
	"bff/internal/management/adapter"
	"bff/internal/management/domain"
	"bff/internal/management/dto/management"
	"bff/internal/management/infra/containers"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	routes   adapter.HttpAdapterHandler
	services containers.ServicesContainer
}

func NewHandler(cs containers.ServicesContainer) *Handler {
	h := new(Handler)
	h.services = cs
	h.routes = adapter.HttpAdapterHandler{
		"GET /api/v1/health": h.Health,

		"GET /api/v1/services":                   h.GetServices,
		"GET /api/v1/services/{service_name}":    h.GetEndpoints,
		"DELETE /api/v1/services/{service_name}": h.DeleteCollection,
		"PATCH /api/v1/services/collection":      h.PatchServiceCollection, // TODO: adicionar prefixo /collection

		"PATCH /api/v1/{collection}":       h.PatchCollection,  // TODO: adicionar prefixo /collection
		"DELETE /api/v1/{collection_name}": h.DeleteCollection, // TODO: adicionar prefixo /collection
	}
	return h
}

func (h Handler) GetRoutes() adapter.HttpAdapterHandler {
	return h.routes
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h Handler) GetEndpoints(w http.ResponseWriter, r *http.Request) error {
	collectionName := r.PathValue("service_name")
	api, err := h.services.Management.GetEndpoints(r.Context(), collectionName)
	if err != nil {
		http.Error(w, "Error getting endpoints", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(api)
}

func (h Handler) GetServices(w http.ResponseWriter, r *http.Request) error {
	services, err := h.services.Management.GetServices(r.Context())
	if err != nil {
		http.Error(w, "Error getting endpoints", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(services)
}

func (h Handler) PatchServiceCollection(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var api management.API
	if err := json.NewDecoder(r.Body).Decode(&api); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return nil
	}

	if err := h.services.Management.RegistryApi(r.Context(), management.ToDomain(api)); err != nil {
		http.Error(w, "Error getting endpoints", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

const (
	environment = "environment"
	collection  = "collection"
)

func (h Handler) PatchCollection(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	collectionName := r.PathValue("collection")
	description := r.FormValue("description")

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return err
	}

	collectionFile, collectionFileHeader, err := r.FormFile(collection)
	if err != nil {
		http.Error(w, "Unable to read collectionFile", http.StatusBadRequest)
		return nil
	}
	defer collectionFile.Close()

	environmentFile, environmentFileHeader, err := r.FormFile(environment)
	if err != nil {
		http.Error(w, "Unable to read environmentFile", http.StatusBadRequest)
		return nil
	}
	defer environmentFile.Close()

	if !strings.Contains(collectionFileHeader.Header.Get("Content-Type"), "json") {
		http.Error(w, "Invalid collectionFile", http.StatusBadRequest)
		return nil
	}

	if !strings.Contains(environmentFileHeader.Header.Get("Content-Type"), "json") {
		http.Error(w, "Invalid environmentFile", http.StatusBadRequest)
		return nil
	}

	listFileName, err := h.services.PostmanCollection.CreateFile(ctx, collectionName, collectionFile, environmentFile)
	if err != nil {
		http.Error(w, "Unable to save collectionFile", http.StatusInternalServerError)
		return nil
	}

	defer h.services.PostmanCollection.UnlinkFiles(ctx, listFileName)

	collectionDto, environmentDto, err := h.services.PostmanCollection.GetInfos(ctx, listFileName)
	if err != nil {
		http.Error(w, "Invalid postman collection on read", http.StatusBadRequest)
		return nil
	}

	endpoints, err := collectionDto.ToDomain(environmentDto.ToDomain())
	if err != nil {
		return err
	}

	api := domain.NewAPI(collectionName, endpoints)

	if err := h.services.Management.RegistryApi(r.Context(), api); err != nil {
		return err
	}

	// Respond with success
	w.WriteHeader(http.StatusAccepted)
	return json.NewEncoder(w).Encode(map[string]string{
		"name":            collectionName,
		"description":     description,
		"collectionFile":  collectionFileHeader.Filename,
		"environmentFile": environmentFileHeader.Filename,
	})
}

func (h Handler) DeleteCollection(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	return nil
}
