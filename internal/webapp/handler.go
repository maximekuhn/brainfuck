package webapp

import (
	"log/slog"
	"net/http"

	"github.com/maximekuhn/brainfuck/internal/webapp/ui"
)

type Handler struct {
	logger  *slog.Logger
	service *service
}

func NewHandler(l *slog.Logger, s *service) *Handler {
	return &Handler{
		logger:  l,
		service: s,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.get(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.post(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("GET endpoint invoked")
	if err := ui.Index().Render(r.Context(), w); err != nil {
		h.logger.Error(
			"failed to render Index",
			slog.String("errMsg", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	const (
		codeMaxLen  = 1_000
		inputMaxLen = 1_000
	)

	h.logger.Info("POST endpoint invoked")

	if err := r.ParseForm(); err != nil {
		h.logger.Error(
			"failed to parse form",
			slog.String("errMsg", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	code := r.Form.Get("bf-code")
	input := r.Form.Get("bf-input")

	if len(code) > codeMaxLen {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(input) > inputMaxLen {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.service.runCode(code, input)
	if err != nil {
		// for now, we will consider the user has sent an invalid brainfuck code
		h.logger.Error(
			"failed to run code",
			slog.String("errMsg", err.Error()),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ui.RunOutput(output).Render(r.Context(), w); err != nil {
		h.logger.Error(
			"failed to render RunOutput",
			slog.String("errMsg", err.Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
