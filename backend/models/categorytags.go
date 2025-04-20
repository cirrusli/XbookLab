package models

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 书籍类目
var AllBookCategories = []Category{
	{1, "小说"},
	{2, "散文"},
	{3, "诗歌"},
	{4, "计算机"},
	{5, "工程"},
	{6, "医学"},
	{7, "中国史"},
	{8, "世界史"},
	{9, "考古"},
	{10, "绘画"},
	{11, "音乐"},
	{12, "建筑"},
	{13, "中国哲学"},
	{14, "西方哲学"},
	{15, "伦理学"},
	{16, "其他"},
}
