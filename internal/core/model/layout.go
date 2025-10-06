package model

type MicroserviceLayout struct {
	LID     int
	DirName string
	SubDir  []SubDirMicroserviceLayout
}

type SubDirMicroserviceLayout struct {
	LID     int
	DirName string
}
