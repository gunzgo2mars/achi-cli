package model

type MicroserviceLayout struct {
	LID     int                        `json:"lid"`
	DirName string                     `json:"dir_name"`
	SubDir  []SubDirMicroserviceLayout `json:"sub_dir"`
}

type SubDirMicroserviceLayout struct {
	LISD    int    `json:"lisd"`
	DirName string `json:"dir_name"`
}

type MicroserviceRootFile struct {
	RFID     int    `json:"rfid"`
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}
