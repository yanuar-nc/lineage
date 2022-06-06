package omdb

type OMDBService interface {
	FindAll(p Params) (*FindAllResponse, error)
	Detail(ID string) (*DetailResponse, error)
}
