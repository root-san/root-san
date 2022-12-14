package repository

type RoomRepository interface {
	// CreateRoom
	CreateRoom(args *RoomArgs) error
}

type RoomArgs struct {
	// RoomId
	Id string
	// RoomName
	Name string
}
