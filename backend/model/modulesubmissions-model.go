package model

type GetModuleSubmissions struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}

type CreateModuleSubmissions struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}

type UpdateModuleSubmissions struct {
	Id       int    `json:"id"`
	ModuleId int    `json:"module_id"`
	File     string `json:"file"`
	Type     string `json:"type"`
	MaxSize  int    `json:"max_size"`
}
