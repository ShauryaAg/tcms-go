package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

type TestCaseHandler struct {
	Repository interfaces.Repository
}

func (h *TestCaseHandler) CreateTestCase(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	testcase, err := h.Repository.Create(r.Context(), data)
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(testcase)
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
