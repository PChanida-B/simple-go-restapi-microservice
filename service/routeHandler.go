package service

import (
	"fmt"
	"net/http"
	"strconv"
)

type Context interface {
	BindJSON(interface{}) error
	JSON(int, interface{})
	Status(int)
	Param(string) string
	GetHeader(string) string
	Next()
	AbortWithStatus(int)
}

type storer interface {
	Create(interface{}) error
	Read(interface{}) error
	ReadIndex(interface{}, int) error
	Update(interface{}, int) error
	Delete(interface{}, int) error
}

type Handler struct {
	store storer
}

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateHandler(c Context) {
	fmt.Println("CreateHandler Comeing")
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	if err := h.store.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		UserName: req.UserName,
		ID:       req.ID,
	})
}

func (h *Handler) ReadHandler(c Context) {
	fmt.Println("ReadHandler Comeing")
	var req Request
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	if err := h.store.ReadIndex(&req, id); err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *Handler) ReadAllHandler(c Context) {
	var req []Request

	if err := h.store.Read(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *Handler) UpdateHandler(c Context) {
	fmt.Println("UpdateHandler Comeing")
	var req Request
	idParam := c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	if err := h.store.Update(&req, id); err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		UserName: req.UserName,
		ID:       uint(id),
	})
}

func (h *Handler) DeleteHandler(c Context) {

	fmt.Println("DeleteHandler Comeing")
	var req Request
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	if err := h.store.Delete(&req, id); err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
