package counter

//默认
const OTHER = 0

//餐饮
const CANYIN = 1

//通行
const TONGXING = 2

//娱乐
const YULE = 3

var TagList = map[int]string{
	CANYIN:   "餐饮",
	TONGXING: "通行",
	YULE:     "娱乐",
	OTHER:    "其他",
}

var TagConfig = map[string]int{
	`(.*?)(饭|餐|食|饮)(.*?)`:         CANYIN,
	`(.*?)(路|车|机票|地铁|通行|交通)(.*?)`: TONGXING,
	`(.*?)(电影|娱乐|唱)(.*?)`:         YULE,
	`(.*?)(其他)(.*?)`:                       OTHER,
}
