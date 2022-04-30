package groups

import (
	"encoding/json"
	"fmt"
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
	MakeGetUserGroupByGroupIdEndpoint() gin.HandlerFunc
	MakeGetUserGroupByUserIdEndpoint() gin.HandlerFunc
	MakeDeleteUserGroupByIdEndpoint() gin.HandlerFunc
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
		fmt.Println(c.Get("user_id"))
		fmt.Println(c.Get("user_type"))
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetGroupByIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetGroupByIdCommand{}
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrGroupIdNotProvided))
			return
		}
		cmd.Id = id
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

func (h *httpEndpoints) MakeGetUserGroupByGroupIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetUserGroupByGroupId{}
		groupId := c.Request.URL.Query().Get("group_id")
		if groupId == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrGroupIdNotProvided))
			return
		}
		cmd.GroupId = groupId
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetUserGroupByUserIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetUserGroupByUserId{}
		userId := c.Request.URL.Query().Get("user_id")
		if userId == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrUserIdNotProvided))
			return
		}
		cmd.UserId = userId
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeDeleteUserGroupByIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &DeleteUserGroupById{}
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrUserGroupIdNotProvided))
			return
		}
		cmd.Id = id
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
