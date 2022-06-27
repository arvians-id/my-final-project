package entity

type ModuleArticles struct {
	Id       int
	CourseId int
	Name     string
	Content  string
	Estimate int
}

type NextPreviousModuleArticles struct {
	Id         int
	CodeCourse string
}
