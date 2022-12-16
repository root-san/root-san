package repository

type Repository interface {
	RoomRepository
	EventRepository
	MemberRepository
}
