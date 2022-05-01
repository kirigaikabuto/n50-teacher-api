package subjects

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreateSubjectEndpoint() gin.HandlerFunc
	MakeListSubjectsEndpoint() gin.HandlerFunc
	MakeGetSubjectByIdEndpoint() gin.HandlerFunc

	MakeCreateTeacherSubjectEndpoint() gin.HandlerFunc
	MakeListTeacherSubjectsEndpoint() gin.HandlerFunc
	MakeGetTeacherSubjectByIdEndpoint() gin.HandlerFunc
	MakeGetTeacherSubjectsByTeacherIdEndpoint() gin.HandlerFunc
	MakeGetTeacherSubjectsBySubjectIdEndpoint() gin.HandlerFunc

	MakeCreateGroupSubjectEndpoint() gin.HandlerFunc
	MakeListGroupSubjectsEndpoint() gin.HandlerFunc
	MakeGetGroupSubjectsByIdEndpoint() gin.HandlerFunc
	MakeGetGroupSubjectByIdTeacherSubEndpoint() gin.HandlerFunc
	MakeGetGroupSubjectByGroupIdEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewSubjectsHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateSubjectEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &CreateSubjectCommand{}
		currentUserId, ok := c.Get("user_id")
		if !ok {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrNoUserIdInToken))
			return
		}
		currentUserType, ok := c.Get("user_type")
		if !ok {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrNoUserTypeInToken))
			return
		}
		cmd.CurrentUserId = currentUserId.(string)
		cmd.CurrentUserType = currentUserType.(string)
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
		respondJSON(c.Writer, http.StatusCreated, resp)
	}
}

func (h *httpEndpoints) MakeListSubjectsEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetSubjectByIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeCreateTeacherSubjectEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeListTeacherSubjectsEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetTeacherSubjectByIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetTeacherSubjectsByTeacherIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetTeacherSubjectsBySubjectIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeCreateGroupSubjectEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeListGroupSubjectsEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetGroupSubjectsByIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetGroupSubjectByIdTeacherSubEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (h *httpEndpoints) MakeGetGroupSubjectByGroupIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {

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
