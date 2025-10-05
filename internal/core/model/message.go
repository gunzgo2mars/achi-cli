package model

type FlagDetail struct {
	Name        string
	Description string
}

type MicroserviceLayout struct {
	LID     int
	DirName string
	SubDir  []SubDirMicroserviceLayout
}

type SubDirMicroserviceLayout struct {
	LID     int
	DirName string
}
