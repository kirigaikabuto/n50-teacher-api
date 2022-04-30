package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	MakeLoginEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewUsersHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeLoginEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &LoginCommand{}
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(jsonData, &cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, resp)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(status)
	w.Write(response)
}
