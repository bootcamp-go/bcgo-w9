package handlers

import (
	"app/internal/movies/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewMovieHandlers returns a new instance of MovieHandlers
func NewMovieHandlers(st storage.StorageMovie) *MovieHandlers {
	return &MovieHandlers{
		st: st,
	}
}

// MovieHandlers is a struct that represents a movie handlers
type MovieHandlers struct {
	st storage.StorageMovie
}

// GetMovies returns all movies
type MovieJSON struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Genre string `json:"genre"`
}
type ResponseBodyGetMovies struct {
	Message string		 `json:"message"`
	Data    []*MovieJSON `json:"data"`
}

// swaggo docs
// @Summary Get movies
// @Description Get all movies
// @Tags movies
// @Produce json
// @Success 200 {object} ResponseBodyGetMovies
// @Failure 500 {object} ResponseBodyGetMovies
// @Router /movies [get]
func (h *MovieHandlers) GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// ...

		// process
		// -> get movies
		m, err := h.st.Get()
		if err != nil {
			code := http.StatusInternalServerError
			body := ResponseBodyGetMovies{Message: "Internal error", Data: nil}

			c.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetMovies{Message: "Success", Data: make([]*MovieJSON, len(m))}
		for i := range m {
			body.Data[i] = &MovieJSON{		// serialize
				Title: m[i].Title,
				Year:  m[i].Year,
				Genre: m[i].Genre,
			}
		}

		c.JSON(code, body)
	}
}

// GetMovieByTitle returns a movie by title
type ResponseBodyGetMovieByTitle struct {
	Message string     `json:"message"`
	Data    *MovieJSON `json:"data"`
}

// swaggo docs
// @Summary Get movie by title
// @Description Get a movie by title
// @Tags movies
// @Produce json
// @Param title path string true "Movie title"
// @Success 200 {object} ResponseBodyGetMovieByTitle
// @Failure 404 {object} ResponseBodyGetMovieByTitle
// @Failure 500 {object} ResponseBodyGetMovieByTitle
// @Router /movies/{title} [get]
func (h *MovieHandlers) GetMovieByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// -> get title from path
		title := c.Param("title")

		// process
		// -> get movie by title
		m, err := h.st.GetByTitle(title)
		if err != nil {
			var code int; var body ResponseBodyGetMovieByTitle
			switch {
			case errors.Is(err, storage.ErrStorageMovieNotFound):
				code = http.StatusNotFound
				body = ResponseBodyGetMovieByTitle{Message: "Movie not found", Data: nil}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodyGetMovieByTitle{Message: "Internal error", Data: nil}
			}

			c.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetMovieByTitle{Message: "Success", Data: &MovieJSON{
			Title: m.Title,
			Year:  m.Year,
			Genre: m.Genre,
		}}

		c.JSON(code, body)
	}
}