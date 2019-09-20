package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/samygp/dummy-rest-api/config"
	log "github.com/sirupsen/logrus"
)

// Handler is the main struct to reference when calling
// the handling methods
type Handler struct {
	guard chan string
}

//NewHandler generates a new handler
func NewHandler() *Handler {
	return &Handler{
		guard: make(chan string, config.Config.Server.MaxRequests),
	}
}

func (h *Handler) updateGuard(calledFrom string) {
	if config.Config.Server.MaxRequests > 0 {
		log.Debugf("Adding to guard: %s", calledFrom)
		h.guard <- calledFrom
	} else {
		log.Debugf(calledFrom)
	}
}

func (h *Handler) clearAndRecover(ctx context.Context, writer http.ResponseWriter, recoveredFrom string) {
	if r := recover(); r != nil {
		log.Errorf("Recovered in %s: %s", recoveredFrom, r)
		RespondError(ctx, writer, fmt.Errorf("Error on operation %s: %s", recoveredFrom, r), http.StatusInternalServerError)
	}
	if config.Config.Server.MaxRequests > 0 {
		log.Debugf("Popped from guard: %s", <-h.guard)
	}
}

//Index method to use as a heartbeat
func (h *Handler) Index(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	Respond(ctx, writer, "OK", http.StatusOK)
	return nil
}

// handleGET receives a GET request and returns OK 200
func (h *Handler) handleGET(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	h.updateGuard("handleGET")
	defer h.clearAndRecover(ctx, writer, "handleGET")
	Respond(ctx, writer, "OK", http.StatusOK)
	return nil
}

// handlePOST receives a POST request and returns OK 200
func (h *Handler) handlePOST(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	h.updateGuard("handlePOST")
	defer request.Body.Close()
	defer h.clearAndRecover(ctx, writer, "handlePOST")
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Errorf("Error reading body: %v", err)
		RespondError(ctx, writer, err, http.StatusBadRequest)
	} else {
		log.Debugf("BODY: %s", string(body))
		Respond(ctx, writer, "OK Post", http.StatusOK)
	}
	time.Sleep(time.Duration(1000) * time.Millisecond)

	return nil
}

// handlePUT receives a PUT request and returns OK 200
func (h *Handler) handlePUT(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	h.updateGuard("handlePUT")
	defer h.clearAndRecover(ctx, writer, "handlePUT")
	Respond(ctx, writer, "OK", http.StatusOK)
	return nil
}

// handleDELETE receives a DELETE request and returns OK 200
func (h *Handler) handleDELETE(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	h.updateGuard("handleDELETE")
	defer h.clearAndRecover(ctx, writer, "handleDELETE")
	Respond(ctx, writer, "OK", http.StatusOK)
	return nil
}
