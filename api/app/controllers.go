// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// MemosController is the controller interface for the Memos actions.
type MemosController interface {
	goa.Muxer
	Create(*CreateMemosContext) error
	Delete(*DeleteMemosContext) error
	List(*ListMemosContext) error
	Show(*ShowMemosContext) error
	Update(*UpdateMemosContext) error
}

// MountMemosController "mounts" a Memos resource controller on the given service.
func MountMemosController(service *goa.Service, ctrl MemosController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/memos", ctrl.MuxHandler("preflight", handleMemosOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/memos/:id", ctrl.MuxHandler("preflight", handleMemosOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateMemosContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MemoPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleMemosOrigin(h)
	service.Mux.Handle("POST", "/memos", ctrl.MuxHandler("create", h, unmarshalCreateMemosPayload))
	service.LogInfo("mount", "ctrl", "Memos", "action", "Create", "route", "POST /memos")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteMemosContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleMemosOrigin(h)
	service.Mux.Handle("DELETE", "/memos/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Memos", "action", "Delete", "route", "DELETE /memos/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListMemosContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleMemosOrigin(h)
	service.Mux.Handle("GET", "/memos", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Memos", "action", "List", "route", "GET /memos")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowMemosContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleMemosOrigin(h)
	service.Mux.Handle("GET", "/memos/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Memos", "action", "Show", "route", "GET /memos/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateMemosContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MemoPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleMemosOrigin(h)
	service.Mux.Handle("PUT", "/memos/:id", ctrl.MuxHandler("update", h, unmarshalUpdateMemosPayload))
	service.LogInfo("mount", "ctrl", "Memos", "action", "Update", "route", "PUT /memos/:id")
}

// handleMemosOrigin applies the CORS response headers corresponding to the origin.
func handleMemosOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateMemosPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateMemosPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &memoPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateMemosPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateMemosPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &memoPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swaggerui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swaggerui/*filepath", "swaggerui/dist")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swaggerui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swaggerui/dist", "route", "GET /swaggerui/*filepath")

	h = ctrl.FileHandler("/swaggerui/", "swaggerui/dist/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swaggerui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swaggerui/dist/index.html", "route", "GET /swaggerui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
