package library

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//取得歌曲数
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//根据索引找音乐
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

//根据歌名找音乐
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

//添加音乐
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//删除音乐
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	} else if index == len(m.musics)-1 && index != 0 {
		m.musics = m.musics[:index]
	} else {
		m.musics = make([]MusicEntry, 0)
	}

	return removedMusic
}
