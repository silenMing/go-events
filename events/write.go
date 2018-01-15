package events

import(
	io "io/ioutil"
	json "encoding/json"
	// "path/filepath"
)

type readFile struct{

}

// func SelfPath() string{
// 	path, _ := filepath.Abs(os.Args[0])
// 	return path
// }

func NewJsonStruct() *readFile{
	return &readFile{}
}

func (self *readFile) Load (filename string, v interface{}) {

	data, err := io.ReadFile(filename)
	
	if err != nil{
		return 
	}
	
	datajson := []byte(data)
	
	err = json.Unmarshal(datajson, v)
	
	if err != nil{
		return
	}
	
}
	
	

		
		
	
	

