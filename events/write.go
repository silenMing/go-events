package events

import(
	io "io/ioutil"
	json "encoding/json"
)

type readFile struct{

}

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
	

		
		
	
	

