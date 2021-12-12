package model

var (
	signalSupervisor *Supervisor = &Supervisor{rooms: make(map[string]*Room)}
	supervisor *Supervisor = &Supervisor{rooms: make(map[string]*Room)}
)

// A supervisor manage battle room.
// A supervisor is singleton.
type Supervisor struct {
	rooms map[string]*Room
}
// type Supervisor struct {
// 	rooms []Hoge
// }

// type TaggedRoom struct {
// 	ID	 string
// 	room *Room
// }

func GetSupervisor() *Supervisor {
	return supervisor
}

func GetSignalSupervisor() *Supervisor {
	return signalSupervisor
}


func (spv *Supervisor) HasRoom(name string) bool {
	 _, exist := spv.rooms[name]
	 return exist
}

func (spv *Supervisor) GetRooms() map[string]*Room {
	return spv.rooms
}

func (spv *Supervisor) GetRoom(name string) *Room {
	return spv.rooms[name]
}

func (spv *Supervisor) SetRoom(name string, room *Room) {
	spv.rooms[name] = room
	return
}

func (spv *Supervisor) NewRoom(name string) *Room {
	room := NewRoom(name)
	spv.SetRoom(room.Name, room)
	return room
}