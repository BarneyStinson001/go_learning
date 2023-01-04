package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop-api/inventory_srv/proto"
	"sync"
)


var invClient proto.InventoryClient
var conn *grpc.ClientConn


func TestSetInv(goodsId,num int32){
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num: num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}
//InvDetail
//Sell
//Reback

func Init(){
	var err error
	conn, err = grpc.Dial("127.0.0.1:3276", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}

func TestSell(wg *sync.WaitGroup){
	defer wg.Done()
	_,err:=invClient.Sell(context.Background(),&proto.SellInfo{
		GoodsInvInfo: []*proto.GoodsInvInfo{
			{GoodsId: 505,Num: 1},
		},
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println("库存扣减成功")
}


func main() {
	Init()
	//TestCreateUser()
	//TestGetGoodsList()
	//TestBatchGetGoods()
	//var  i int32
	//for i= 500;i<=1000;i++{
	//	TestSetInv(i,100)
	//}
	//wg防止主进程退出
	var  wg sync.WaitGroup
	wg.Add(10)

	for i:=0;i<10;i++{
		//TestSell(&wg)//顺序执行
		go TestSell(&wg)//协程并发

	}

	wg.Wait()

	conn.Close()
}
