package trade

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	nethttp "net/http"
	neturl "net/url"
	"upi/common/turn"
	"upi/expand/unionpay/core"

	"upi/expand/unionpay"
)

type WebPay unionpay.Service

const (
	UnifiedPay = "/n_web_pay.api" //支付请求
	GET_PAY_RESULT = "/n_get_pay_result.api" //支付状态查询
	PAY_CANCEL = "/n_pay_cancel.api" //取消订单请求
)

// UnifiedPay 统一支付
func (w *WebPay) UnifiedPay(ctx context.Context, req UnifiedPayRequest) (resp *UnifiedPayResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarHeaderParams.Add("charset", "UTF-8")

	urlPath := w.Config.Host + UnifiedPay

	err = w.BeforeUnionPay(ctx, &req)
	if err != nil {
		return nil, nil, err
	}
	localVarPostBody = req

	localVarHTTPContentType := "application/x-www-form-urlencoded"
	logx.Infof("银联支付::接口请求::接口%s,请求参数：%s", UnifiedPay, turn.StructToJson(localVarPostBody))
	result, err = w.Client.Request(ctx, localVarHTTPMethod, urlPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		logx.Infof("银联支付::失败返回::接口%s,错误信息：%s", UnifiedPay, err.Error())
		return nil, result, err
	}
	logx.Infof("银联支付::返回::接口%s,响应参数：%s", UnifiedPay, turn.StructToJson(result.Response))
	resp = new(UnifiedPayResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return nil, result, err
}

func (w *WebPay) CancelPay()  {

}



func (w *WebPay) QueryPay()  {

}
