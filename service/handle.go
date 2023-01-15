package service

type Context interface {
	BindJSON() (Request, error)
	JSON(int, interface{})
	Status(int)
}

type storer interface {
	Create(Request) error
	Read(Request) error
	Update(Request) error
	Delete(Request) error
}

type Handler struct {
	store storer
}

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CreateHandler(c Context) {}

func (h *Handler) ReadHandler(c Context) {}

func (h *Handler) UpdateHandler(c Context) {}

func (h *Handler) DeleteHandler(c Context) {}
