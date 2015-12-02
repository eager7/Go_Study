package manager

import "errors"

type MusicEntry struct{
	Id 		string
	Name 	string
	Artist	string
	Source  string
	Type 	string
}

type MusicManager struct{
	musics []MusicEntry
}

func NewMusicManager() *MusicManager{
	return &MusicManager{make([]MusicEntry,0)}
}

func (m *MusicManager)Len()int {
	return len(m.musics)
}

func (m *MusicManager)Get(index int)(music *MusicEntry, err error){
	if index < 0 || index > len(m.musics){
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager)Find(name string)(index int, music *MusicEntry){
	if len(m.musics) == 0{
		return 0,nil
	}
	for id,music := range m.musics{
		if music.Name == name{
			return id,&music
		}
	}
	return 0,nil
}

func (m *MusicManager)Add(music *MusicEntry){
	m.musics = append(m.musics, *music)
}

func (m *MusicManager)Remove(index int)(*MusicEntry){
	if index < 0 || index > len(m.musics){
		return nil
	}
	removedMusic := &m.musics[index]
	if index == 0{
		m.musics = m.musics[1:]
	}else if index < len(m.musics)-1 {
		m.musics = append(m.musics[index-1:],m.musics[:index]...)
	}
	return removedMusic
}

func (m *MusicManager)RemoveByName(name string){
	if 0 == len(m.musics){
		return
	}
	id,music := m.Find(name)
	if music == nil{
		println("Can,t find this file")
		return
	}
	m.Remove(id)
}









