package manager
import "testing"
import "fmt"

var mm *MusicManager

func TestOps(t *testing.T) {
	println("This is a testing project")

	mm = NewMusicManager()
	if mm == nil{
		t.Error("NewMusicManager Failed")
	}
	if mm.Len() != 0{
		t.Error("MusicManager is not empty")
	}
	m0 := &MusicEntry{
		"1",
		"Hello",
		"PCT",
		"ShenZhen",
		"MP3",
	}
	mm.Add(m0)
	if mm.Len() != 1{
		t.Error("Add new file Failed")
	}
	fmt.Println(mm)

	m1 := &MusicEntry{
		"2",
		"World",
		"PCT",
		"ShenZhen",
		"MP3",
	}
	mm.Add(m1)
	fmt.Println(mm)

	_,m_f := mm.Find(m1.Name)
	if m_f == nil{
		t.Error("find file Failed")
	}else if  m_f.Name != m1.Name || m_f.Id != m1.Id || m_f.Artist != m1.Artist || m_f.Type != m1.Type || m_f.Source != m1.Source{
		t.Error("find error")
	}else{
		fmt.Println(m_f)
	}

	m_f2, err := mm.Get(0)
	if m_f2 == nil {
		t.Error("Get file Failed", err)
	}else{
		fmt.Println(m_f2)
	}

	mm.RemoveByName(m0.Name)

	fmt.Println(mm)
}


