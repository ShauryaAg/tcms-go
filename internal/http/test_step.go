package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	utils "github.com/ShauryaAg/tcms-go/pkg/http"
)

type TestStepHandler struct {
	Handler
}

func (h *TestStepHandler) CreateTestStep(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	testcase, err := h.Repository.Create(r.Context(), data)
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(testcase)
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
