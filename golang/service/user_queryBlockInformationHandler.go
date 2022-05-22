package service

import (
	"A11Smile/eth"
	"A11Smile/serializer"
	"context"
	"github.com/gin-gonic/gin"
	"math"
	"math/big"
)

func user_queryBlockInformationHandler(c *gin.Context) {

	num,_ := eth.Client.BlockNumber(context.Background())

	var res []interface{}
	for i := int(num) - 1000; i < int(num); i++ {
		e,_ := eth.Client.BlockByNumber(context.Background(),big.NewInt(int64(i)))
		if e.Transactions().Len() != 0 {
			for _,v := range e.Transactions() {
				fBalance := new(big.Float)
				fBalance.SetString(v.Value().String())
				balanceEther := new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))
				fBalance.SetString(v.GasPrice().String())
				gasPrice := new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))

				r := struct {
					Value interface{}
					To interface{}
					GasPrice interface{}
					Hash interface{}
					BlockNum interface{}
					Nonce interface{}
				}{
					Value: balanceEther.String(),
					To: v.To(),
					GasPrice: gasPrice.String(),
					Hash: v.Hash(),
					BlockNum: i,
					Nonce: v.Nonce(),
				}

				res = append(res, r)
			}
		}
	}


	serializer.RespOK(c,res)


}