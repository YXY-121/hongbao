package envelope

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"hongbao/apis/dto"
	"hongbao/store"
	"hongbao/store/mysql"
	"math/rand"
	"time"
)

type EnvelopeService struct {
	envelopeDao store.OperateEnvelope
}
type EnvelopeDTO struct {
	id int

}
//总红包
type EnvelopeVo struct {
	EnvelopeId int//红包id
	SenderId int //发送者
	blessing string //祝福
	Number int //个数
	TotalAmount decimal.Decimal //总金额
	RemainedAmount decimal.Decimal//剩的 如果剩的>0，就去拿红包小项、
}
//红包项
type EnvelopeItemVo struct {
	EnvelopeId uuid.UUID//所属大红包id
	EnvelopeItemId uuid.UUID//红包item
	AccepterId int//接收者
	ItemAmount decimal.Decimal//单个红包金额
}

//分配红包红包类型生成响应的红包数量、以及money\生成红包
func (e *EnvelopeService) DivideHongbao(evo *vo.EnvelopeVo) {
	//设置过期时间
	expiredTime := evo.CreateTime.Add(24 * time.Hour)
	updateTime := evo.CreateTime
	amounts := make([]decimal.Decimal, evo.Number)

	if evo.Status == "1" { //普通红包用一个红包id搞定
		//单个金额
		singleAmount := evo.TotalAmount.Div(decimal.NewFromInt(int64(evo.Number)))
		for i:=0;i<evo.Number;i++ {
			amounts=append(amounts,singleAmount)
		}

	} else {
	//	amounts = RandomCreateHongbao(evo.TotalAmount, evo.Number)
	}
	//生成红包小项

//	EnvelopeItemList:=CreateEnvelopeItem(evo.EnvelopeId, amounts)
	//fmt.Println(EnvelopeItemList)
	//红包小项插入缓存、数据库
	//将当前的值插入缓存
	evoj,_:=json.Marshal(evo)
	edo:=new(mysql.EnvelopeDo)
	json.Unmarshal(evoj,edo)
	edo.ExpiredTime=expiredTime
	edo.UpdateTime=updateTime
	fmt.Println(edo)
	fmt.Println(evoj)
	fmt.Println(evo)


	//红包大项插入缓存、数据库
}

//生成红包小项
func CreateEnvelopeItem(EnvelopeId uuid.UUID,amounts []decimal.Decimal) []EnvelopeItemVo {
	EnvelopeItemList:=make([]EnvelopeItemVo, len(amounts))
	for i:=0;i< len(amounts);i++ {
		itemVo := EnvelopeItemVo{
			EnvelopeId:     EnvelopeId, //所属大红包id
			EnvelopeItemId: uuid.New(), //红包item
			ItemAmount:     amounts[i], //单个红
		}
		EnvelopeItemList=append(EnvelopeItemList,itemVo)

	}
	return EnvelopeItemList
}
//每次抢的时候就--，抢的时候生成小红包的信息。
func randomMoney(restAmount int, restNumber int) decimal.Decimal{

		amount := rand.Intn(restAmount/restNumber*2-1) + 1
		restAmount-=amount
		v:=decimal.NewFromInt(int64(amount))
		fmt.Println(v)
		fmt.Println(amount)

	return v
}

func RandomCreateHongbao(hongbaoId string)  {
	//从缓存中获取当前的number和amount，
	mysql.RedisDB.Do("get"+ hongbaoId)
	//获取大红包id的库存
	//并生成randomMoney
	//在库存中number减去1，amount-randomMoney
}


//红包的首页：红包的金额、红包数量、发红包的人名字、红包的主福、当前剩余总书、:抢到的人的名字、金额
func (e *EnvelopeService) GetHongbaoMessage(ctx context.Context, envelopeId int) (*mysql.EnvelopeDo, error) {
	//获取大红包

	//获取红包小项们

	//查看缓存是否存
	//RedisEnvelope,err:=mysql.RedisDB.Do("get",envelopeId)

	//if err!=nil {
	//	////获取红包小项
	//	//RedisEnvelopeItem,err:=mysql.RedisDB.Do("get",envelopeId)
	//	//
	//	//return RedisEnvelope;
	//}

	//存在就在数据库取

	//不存在就去数据库取
	envelope, err := e.envelopeDao.SelectEnvelopeById(ctx, envelopeId)
	return envelope, err
}


