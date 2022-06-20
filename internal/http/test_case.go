package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	utils "github.com/ShauryaAg/tcms-go/pkg/http"
	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	repointerfaces "github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

type TestCaseHandler struct {
	Router     httpinterfaces.Router
	Repository repointerfaces.Repository
}

func (h *TestCaseHandler) CreateTestCase(w http.ResponseWriter, r *http.Request) {
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

func (h *TestCaseHandler) GetTestCase(w http.ResponseWriter, r *http.Request) {
	objectId := h.Router.URLParam(*r, "ID")
	model, err := h.Repository.FindById(r.Context(), objectId)
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(model)
	if err != nil {
		utils.Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
