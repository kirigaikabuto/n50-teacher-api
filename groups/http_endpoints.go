package groups

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreateGroupEndpoint() gin.HandlerFunc
	MakeListGroupEndpoint() gin.HandlerFunc
	MakeGetGroupByIdEndpoint() gin.HandlerFunc
	MakeCreateUserGroupEndpoint() gin.HandlerFunc
	MakeGetGroupByGroupIdEndpoint() gin.HandlerFunc
	MakeGetGroupByUserIdEndpoint() gin.HandlerFunc
	MakeDeleteGroupByIdEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewUserGroupHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateGroupEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &CreateGroupCommand{}
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

func (h *httpEndpoints) MakeListGroupEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &ListGroupCommand{}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetGroupByIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetGroupByIdCommand{}
		groupId := c.Request.URL.Query().Get("id")
		if groupId == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrGroupIdNotProvided))
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

func (h *httpEndpoints) MakeCreateUserGroupEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &CreateUserGroupCommand{}
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

func (h *httpEndpoints) MakeGetGroupByGroupIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetUserGroupByGroupId{}
		groupId := c.Request.URL.Query().Get("group_id")
		if groupId == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrGroupIdNotProvided))
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

func (h *httpEndpoints) MakeGetGroupByUserIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetUserGroupByUserId{}
		userId := c.Request.URL.Query().Get("user_id")
		if userId == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrUserIdNotProvided))
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

func (h *httpEndpoints) MakeDeleteGroupByIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &DeleteUserGroupById{}
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrUserGroupIdNotProvided))
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
