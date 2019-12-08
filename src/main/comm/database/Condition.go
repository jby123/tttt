package database

//针对字符串 定义个 查询过滤条件
const (
	Equals    int = 0
	Like      int = 1
	LeftLike  int = 2
	RightLike int = 3
)

type Condition struct {
	Operate int
	Value   string
}
