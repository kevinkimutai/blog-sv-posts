package ports

type MessagePort interface {
	GetMovieRating(movieID string)
}
