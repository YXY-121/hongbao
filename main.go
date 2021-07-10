package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
	_ "github.com/urfave/cli/v2"
	"hongbao/apis"
	"math/rand"
	"net/http"
	"strings"
)

type s struct {
	a int `json:"a"`
	b string `json:"b"`
}
func disserilizabale(GetValue interface{},vo interface{})  interface{}{
	str:=GetValue.([]byte)
	json.Unmarshal(str,&vo)
	return vo
}
func main()  {
	r:=apis.APIRouter()
	r.Run()


	//mysql.GetRedis()
	//e:=mysql.EnvelopeDo{
	//	EnvelopeId: 1,
	//}
	//
	////jsonData, _ := json.Marshal(e)
	//jsonData, _ := json.MarshalIndent(e, "", "   ")
	//fmt.Println(string(jsonData))
	//a,_:=mysql.RedisDB.Do("set","Aa1",string(jsonData))
	//b,_:=mysql.RedisDB.Do("get","Aa1")
	//fmt.Println(a)
	//stu1:=disserilizabale(b,mysql.EnvelopeDo{})
	//do := stu1.(mysql.EnvelopeDo)
	//
	//
	//fmt.Println(do.EnvelopeId)




}

func request()  {
	req, err := http.NewRequest("GET", "https://apache.claz.org/zookeeper/zookeeper-3.7.0/apache-zookeeper-3.7.0-bin.tar.gz", nil)
	if err != nil {
		fmt.Println("right")
	}
	rangeStart := 2000
	rangeStop := 3000
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeStop))

	res, err := http.DefaultClient.Do(req)
	fmt.Println(res)

}
//func  downloadPartial(strURL, filename string, rangeStart, rangeEnd, i int) {
//	if rangeStart >= rangeEnd {
//		return
//	}
//
//	req, err := http.NewRequest("GET", strURL, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	flags := os.O_CREATE | os.O_WRONLY
//	partFile, err := os.OpenFile(d.getPartFilename(filename, i), flags, 0666)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer partFile.Close()
//
//	buf := make([]byte, 32*1024)
//	_, err = io.CopyBuffer(partFile, resp.Body, buf)
//	if err != nil {
//		if err == io.EOF {
//			return
//		}
//		log.Fatal(err)
//	}
//}

// getPartDir 部分文件存放的目录
func  getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

// getPartFilename 构造部分文件的名字
func  getPartFilename(filename string, partNum int) string {
	partDir :="s"
	 	//d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partDir, filename, partNum)
}




func testFlag()  {
	n:=flag.String("n","yanxy","名字")
	o:=flag.Int("o",13,"年龄")
	sex:=flag.Bool("sex",true,"是不是女的")
	flag.Parse()    //解析命令行参数
	flag.PrintDefaults()    //输出帮助信息
	fmt.Println(*n, *o, *sex)//go run main.go -n "熊率" -o 30 -sex false

}
func RandomCreateHongbao(totalAmount float64, quantity int) []decimal.Decimal{
	now := decimal.NewFromFloat(totalAmount)
	hongbaoList := make([]decimal.Decimal, quantity)
	len:= quantity
	for i := 0; i < len-1; i++ {
		max := decimal.NewFromFloat((totalAmount / float64(quantity)) * 2)
		ran := decimal.NewFromFloat(rand.Float64())
		value := max.Mul(ran)
		fmt.Println(value)
		if value.LessThan(decimal.NewFromFloat(0.01)) {
			value = decimal.NewFromFloat(0.01)

		} else {
			value = value.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100))

		}
		hongbaoList = append(hongbaoList, value)
		quantity--
		now=now.Sub(value)

	}

	hongbaoList = append(hongbaoList, now.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100)))
	return hongbaoList
}
