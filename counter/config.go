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
	OTHER:    "餐饮",
	CANYIN:   "通行",
	TONGXING: "娱乐",
	YULE:     "其他",
}

var TagConfig = map[string]int{
	`(.*?)(饭|餐|食|饮)(.*?)`:      CANYIN,
	`(.*?)(路|车|机票|地铁|通行|交通)(.*?)`: TONGXING,
	`(.*?)(电影|娱乐|唱)(.*?)`:      YULE,
	`(.*?)`:                    OTHER,
}
