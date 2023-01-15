package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PChanida-B/simple-go-restapi-microservice/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	return &Router{gin.Default()}
}

func (r *Router) ListenAndServe() func() {
	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		<-ctx.Done()
		stop()
		fmt.Println("shutting down gracefully, press Ctrl+C again to force")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(timeoutCtx); err != nil {
			fmt.Println(err)
		}
	}
}

/* Convert Function App Format to Gin Format */
type HandlerFunc func(service.Context)

func (r *Router) GET(relativePath string, handler HandlerFunc) {
	r.Engine.GET(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func (r *Router) POST(relativePath string, handler HandlerFunc) {
	r.Engine.POST(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func (r *Router) PUT(relativePath string, handler HandlerFunc) {
	r.Engine.PUT(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func (r *Router) DELETE(relativePath string, handler HandlerFunc) {
	r.Engine.DELETE(relativePath, func(c *gin.Context) {
		handler(&Context{c})
	})
}

/* Convert Function App Format to Gin Format */
type Context struct {
	*gin.Context
}

func (c *Context) BindJSON() (r service.Request, err error) {
	err = c.ShouldBindJSON(&r)
	return
}

func (c *Context) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}

func (c *Context) Status(statusCode int) {
	c.Context.Status(statusCode)
}
