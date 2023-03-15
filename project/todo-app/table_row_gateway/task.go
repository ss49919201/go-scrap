package tablerowgateway

type TaskRowGateway interface {
	Insert() error
	Update() error
	Delete() error
}
type TaskFinder interface {
	First(id int) (TaskRowGateway, error)
	FindByUserID(userID int) ([]TaskRowGateway, error)
}
