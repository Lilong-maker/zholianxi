package service

import (
	"context"
	"fmt"
	"zholianxi/src/basic/config"
	"zholianxi/src/handler/model"

	__ "zholianxi/src/basic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedGoodsServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {

	var goods model.Goods
	err := goods.FindGoods(config.DB, in.Name)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "参数错误",
			Code: 400,
		}, nil
	}
	Good := model.Goods{
		Name:  in.Name,
		Price: float32(in.Price),
		Num:   int(in.Num),
		Color: in.Color,
		Sign:  in.Sign,
	}
	err = Good.GoodsAdd(config.DB)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "商品添加失败",
			Code: 400,
		}, nil
	}
	go func() {
		esAdd := map[string]interface{}{
			"Name":  in.Name,
			"Price": in.Price,
			"Num":   in.Num,
			"Color": in.Color,
			"Sign":  in.Sign,
		}
		_, err = config.Elastic.Index().Index("good").BodyJson(esAdd).Do(context.Background())
		if err != nil {
			return
		}
		fmt.Println("es同步失败")

	}()
	return &__.GoodsAddResp{
		Msg:  "商品添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.LoginResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	if user.Password != in.Password {
		return &__.LoginResp{
			Msg:  "密码错误",
			Code: 400,
		}, nil
	}
	//handler, err := middleware.TokenHandler(strconv.Itoa(int(user.ID)))
	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
		Id:   int64(user.ID),
	}, nil
}
