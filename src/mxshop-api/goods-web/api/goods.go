package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"net/http"
	"strconv"
	"strings"
)

//这个函数如果放在initialize的validator里面，就会涉及到循环导入的问题，这边要引入initialize包，显示 import cycle not allowed
func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		idx := strings.Index(field, ".")
		rsp[field[idx+1:]] = err
	}
	return rsp
}

func HandlerValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误" + e.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					//"msg":"用户服务不可用",
					"msg": e.Message(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
		}
	}
}

func GetGoodsList(c *gin.Context)  {

	//向service层 构建请求
	request:=&proto.GoodsFilterRequest{}

	//商品过滤,值从url后面的参数获取 类型转换
	priceMin := c.DefaultQuery("pmin","0")
	printMinInt,_:= strconv.Atoi(priceMin) //priceMin=abc  	//要不要报错，可以忽略，默认值为0
	request.PriceMin= int32(printMinInt)

	priceMax := c.DefaultQuery("pmax","0")
	printMaxInt,_:= strconv.Atoi(priceMax) //priceMin=abc
	request.PriceMax= int32(printMaxInt)

	isHot:=c.DefaultQuery("ih","0")
	if isHot == "1"{
		request.IsHot= true
	}
	isNew:=c.DefaultQuery("in","0")
	if isNew== "1"{
		request.IsNew= true
	}
	isTab:=c.DefaultQuery("it","0")
	if isTab== "1"{
		request.IsTab= true
	}

	//category需要转换
	categoryId:= c.DefaultQuery("c","0")//注意service层的判断
	categoryIdInt,_:=strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages:= c.DefaultQuery("p","0")
	pagesInt,_:=strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	numsPerPage:= c.DefaultQuery("pnum","10")
	numsPerPageInt,_:=strconv.Atoi(numsPerPage)
	request.PagePerNums = int32(numsPerPageInt)

	keywords:= c.DefaultQuery("q","")
	request.KeyWords=keywords

	brandId:= c.DefaultQuery("b","0")
	brandIdInt,_:=strconv.Atoi(brandId)
	request.Pages = int32(brandIdInt)

	//请求goods-service层

	rsp,err:=global.GoodsSrvClient.GoodsList(context.Background(),request)
	//出错处理，从user那边考过来
	if err != nil {
		zap.S().Errorw("[GetGoodsList] 查询【商品列表】 失败")
		HandlerGrpcErrorToHttp(err, c)
		return
	}

	//处理返回数据，有业务决定哪些要返回，哪些不返回
	//等于转发    spring-cloud没有分层
	reMap:=map[string]interface{}{
		"total":rsp.Total,

		//"data":rsp.Data,//GoodsInfoResponse grpc自动改标签
	}
	//换手动拼接,后期维护成本会低一些,根据前端文档.符合文档规范
	goodsList:=make([]interface{},0)
	for _,value :=range rsp.Data{
		goodsList = append(goodsList,map[string]interface{}{
			"id":value.Id,
			"name": value.Name,
			"goods_brief":value.GoodsBrief,
			"desc":value.GoodsDesc,
			"ship_free": value.ShipFree,
			"images": value.Images,
			"desc_image":value.DescImages,
			"front_image":value.GoodsFrontImage,
			"shop_price":value.ShopPrice,
			"categpry":map[string]interface{}{
				"id":value.Category.Id,
				"name":value.Category.Name,
			},
			"brand":map[string]interface{}{
				"id":value.Brand.Id,
				"name":value.Brand.Name,
				"logo":value.Brand.Logo,
			},
			"is_hot":value.IsHot,
			"is_new":value.IsNew,
			"on_sale":value.OnSale,
		})
	}
	reMap["data"]=goodsList

	c.JSON(http.StatusOK,reMap)
}
