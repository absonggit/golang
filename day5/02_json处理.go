package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 成员变量名首字母必须大写
type IT struct {
	Company  string   // `json:"-"` // 此字段不会输出到屏幕
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"` // 转换类型
	Price    float64  `json:",string"` // 需要输入小写字母，可以这样二次编码
}

func main() {
	// 通过结构体生成json
	s := IT{"idcast", []string{"go", "c++", "python"}, true, 123.123}
	j1, err := json.Marshal(s)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("json1 = ", string(j1))

	// 通过map生成json
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"go", "c++", "python"}
	m["isok"] = true
	m["price"] = 123.123

	result, err := json.MarshalIndent(m, "", " ") // 格式化编码
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("json2 = ", string(result))

	// json解析成结构体
	json1 := string(j1)
	var tmp IT
	err = json.Unmarshal([]byte(json1), &tmp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", tmp)
	fmt.Println(tmp.Subjects)

	// json解析成map
	m1 := make(map[string]interface{}, 4)
	err = json.Unmarshal([]byte(result), &m1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", m1)
	fmt.Println(m1["subjects"])

	// 类型断言
	fmt.Printf("类型\t键\t值\n")
	for key, value := range m1 {
		switch data := value.(type) {
		case string, bool, float64, int, byte, complex128, []byte, []interface{}, nil, []string:
			fmt.Printf("%T\t%v\t%v\n", data, key, value)
		default:
			fmt.Printf("其他\t%v\t%v\n", key, value)
		}
	}
	// 使用反射获取值、类型
	fmt.Println(reflect.TypeOf(m1["subjects"]), reflect.ValueOf(m1["subjects"]))

	// 修改map中interface的值
	m1["subjects"] = []string{"aa", "bb", "cc"}
	fmt.Println(m1["subjects"])

}
