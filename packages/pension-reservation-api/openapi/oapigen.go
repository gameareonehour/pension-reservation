// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"github.com/gofiber/fiber/v2"
)

// GetReleaseNotesResponse defines model for GetReleaseNotesResponse.
type GetReleaseNotesResponse struct {
	Items []struct {
		CreatedAt string `json:"created_at"`
		Id        int    `json:"id"`
		Text      string `json:"text"`
		UpdatedAt string `json:"updated_at"`
	} `json:"items"`
}

// PostVacancyRoomsSearchResponse defines model for PostVacancyRoomsSearchResponse.
type PostVacancyRoomsSearchResponse struct {
	Items []struct {
		Dayfee   int    `json:"dayfee"`
		Id       int    `json:"id"`
		ImageUrl string `json:"image_url"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	} `json:"items"`
}

// PostVacancyRoomsSearchJSONBody defines parameters for PostVacancyRoomsSearch.
type PostVacancyRoomsSearchJSONBody struct {
	ReservationDate *string `json:"reservation_date,omitempty"`
}

// PostVacancyRoomsSearchJSONRequestBody defines body for PostVacancyRoomsSearch for application/json ContentType.
type PostVacancyRoomsSearchJSONRequestBody PostVacancyRoomsSearchJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 最新リリースノート取得API
	// (GET /release-notes)
	GetReleaseNotes(c *fiber.Ctx) error
	// 空室検索API
	// (POST /vacancy-rooms/search)
	PostVacancyRoomsSearch(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// GetReleaseNotes operation middleware
func (siw *ServerInterfaceWrapper) GetReleaseNotes(c *fiber.Ctx) error {

	return siw.Handler.GetReleaseNotes(c)
}

// PostVacancyRoomsSearch operation middleware
func (siw *ServerInterfaceWrapper) PostVacancyRoomsSearch(c *fiber.Ctx) error {

	return siw.Handler.PostVacancyRoomsSearch(c)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Get(options.BaseURL+"/release-notes", wrapper.GetReleaseNotes)

	router.Post(options.BaseURL+"/vacancy-rooms/search", wrapper.PostVacancyRoomsSearch)

}
