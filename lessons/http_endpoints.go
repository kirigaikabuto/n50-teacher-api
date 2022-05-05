package lessons

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpEndpoints interface {
	MakeCreateLessonEndpoint() gin.HandlerFunc
	MakeGetLessonByIdEndpoint() gin.HandlerFunc
	MakeUpdateLessonEndpoint() gin.HandlerFunc
	MakeListLessonByGroupSubjectIdEndpoint() gin.HandlerFunc
	MakeDeleteLessonEndpoint() gin.HandlerFunc

	MakeUploadFileEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewLessonHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateLessonEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &CreateLessonCommand{}
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

func (h *httpEndpoints) MakeGetLessonByIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &GetLessonByIdCommand{}
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
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrLessonIdNotProvided))
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

func (h *httpEndpoints) MakeUpdateLessonEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &UpdateLessonCommand{}
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
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrLessonIdNotProvided))
			return
		}
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
		cmd.Id = id
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusCreated, resp)
	}
}

func (h *httpEndpoints) MakeListLessonByGroupSubjectIdEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &ListLessonByGroupSubjectIdCommand{}
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
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrGroupSubjectIdNotProvided))
			return
		}
		cmd.GroupSubjectId = id
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		respondJSON(c.Writer, http.StatusOK, common.Resp{Objects: resp})
	}
}

func (h *httpEndpoints) MakeDeleteLessonEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := &DeleteLessonCommand{}
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
		id := c.Request.URL.Query().Get("id")
		if id == "" {
			respondJSON(c.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(ErrLessonIdNotProvided))
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

func (h *httpEndpoints) MakeUploadFileEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &UploadFileCommand{}
		fileId := context.Request.URL.Query().Get("id")
		if fileId == "" {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrFileIdNotProvided))
			return
		}
		cmd.Id = fileId
		buf := bytes.NewBuffer(nil)
		file, header, err := context.Request.FormFile("file")
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		fileInfo := strings.Split(header.Filename, ".")
		cmd.Name = fileInfo[0]
		cmd.Type = fileInfo[1]
		_, err = io.Copy(buf, file)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = file.Close()
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		cmd.File = buf
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
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
