package mt5

import (
	"encoding/json"
	"fmt"

	"github.com/HimanshuM/go-mt5/constants"
)

type Trade struct {
	ID                   string
	Login                string
	ExternalAccount      string
	Group                string
	Symbol               string
	Digits               string
	Action               string
	TimeExpiration       string
	Type                 string
	TypeFill             string
	TypeTime             string
	Flags                string
	Volume               string
	VolumeExt            string
	Order                string
	OrderExternalID      string
	PriceOrder           string
	PriceTrigger         string
	PriceSL              string
	PriceTP              string
	PriceDeviation       string
	PriceDeviationTop    string
	PriceDeviationBottom string
	Comment              string
	ResultRetcode        string
	ResultDealer         string
	ResultDeal           string
	ResultOrder          string
	ResultVolume         string
	ResultVolumeExt      string
	ResultPrice          string
	ResultDealerBid      string
	ResultDealerAsk      string
	ResultDealerLast     string
	ResultMarketBid      string
	ResultMarketAsk      string
	ResultMarketLast     string
	ResultComment        string
	IDClient             string
	IP                   string
	SourceLogin          string
	Position             string
	PositionBy           string
	PositionExternalID   string
	PositionByExternalID string
	APIData              []*APIData `json:"ApiData"`
}

type TradeResponse struct {
	ID string `json:"id"`
}

type APIData struct {
	AppID       string
	ID          string
	ValueInt    string
	ValueUInt   string
	ValueDouble string
}

type TradeConfirm struct {
	ID                 string
	DealID             string
	OrderID            string
	PositionExternalID string
	Retcode            string
	ExternalRetcode    string
	Volume             string
	VolumeExt          string
	Price              string
	PriceGateway       string
	TickBid            string
	TickAsk            string
	TickLast           string
	Flags              string
	Comment            string
}

type TradeResult struct {
	Result *TradeConfirm `json:"result"`
	Answer *Trade        `json:"answer"`
}

func (t *Trade) toJSON() (string, error) {
	body, err := json.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("error marshalling trade to JSON: %v", err)
	}
	return string(body), nil
}

func (t *Trade) consumeResponse(res *Response) error {
	if tradeID, present := res.Parameters[constants.ParamID]; present {
		t.ID = tradeID.(string)
		return nil
	}
	tradeResponse := &TradeResponse{}
	if err := json.Unmarshal([]byte(res.Data), tradeResponse); err != nil {
		return fmt.Errorf("error parsing JSON for trade response: %v", err)
	}
	if tradeResponse.ID == "" {
		return fmt.Errorf("error sending request: trade ID not received")
	}
	t.ID = tradeResponse.ID
	return nil
}

func consumeTradeResponse(res *Response, tradeID string) (*TradeResult, error) {
	responseMap := make(map[string][]interface{})
	if err := json.Unmarshal([]byte(res.Data), &responseMap); err != nil {
		return nil, fmt.Errorf("error parsing trade response: %v", err)
	}
	tradeResponseMapArr, present := responseMap[tradeID]
	if !present {
		return nil, fmt.Errorf("error getting trade response: Not Found")
	}
	if len(tradeResponseMapArr) != 2 {
		return nil, fmt.Errorf("error getting trade response: empty response for trade")
	}
	return getTradeResult(tradeResponseMapArr)
}

func getTradeResult(responseMap []interface{}) (*TradeResult, error) {
	responseObjects := []map[string]interface{}{
		responseMap[0].(map[string]interface{}),
		responseMap[1].(map[string]interface{}),
	}
	answerIndex, result, err := getResultFromTradeResult(responseObjects)
	if err != nil {
		return nil, err
	}
	answer, err := getAnswerFromTradeResult(responseObjects[answerIndex])
	if err != nil {
		return nil, err
	}
	return &TradeResult{
		Result: result,
		Answer: answer,
	}, nil
}

func getResultFromTradeResult(responseObjects []map[string]interface{}) (int, *TradeConfirm, error) {
	if result, present := responseObjects[0][constants.ParamTradeResult]; present {
		resultObj, err := getResultObject(result)
		return 1, resultObj, err
	} else if result, present := responseObjects[1][constants.ParamTradeResult]; present {
		resultObj, err := getResultObject(result)
		return 0, resultObj, err
	}
	return -1, nil, fmt.Errorf("error parsing trade response: result object not found")
}

func getResultObject(result interface{}) (*TradeConfirm, error) {
	resultMap, okay := result.(map[string]interface{})
	if !okay {
		return nil, fmt.Errorf("error parsing trade response: unknown object for result")
	}
	return getResultFromMap(resultMap), nil
}

func getAnswerFromTradeResult(responseObject map[string]interface{}) (*Trade, error) {
	answerRaw, present := responseObject[constants.ParamTradeAnswer]
	if !present {
		return nil, fmt.Errorf("error parsing trade response: answer object not found")
	}
	answerMap, okay := answerRaw.(map[string]interface{})
	if !okay {
		return nil, fmt.Errorf("error parsing trade response: unknown object for answer")
	}
	answer := getAnswerFromMap(answerMap)
	apiDataArr, err := getAPIDataArrayFromResponse(answerMap)
	if err != nil {
		return nil, err
	}
	answer.APIData = apiDataArr
	return answer, nil
}

func getAPIDataArrayFromResponse(answerMap map[string]interface{}) ([]*APIData, error) {
	var apiDataArr []*APIData
	apiDataRaw, present := answerMap["ApiData"]
	if !present {
		return apiDataArr, nil
	}
	apiDataRawArr, okay := apiDataRaw.([]interface{})
	if okay {
		apiDataArr := make([]*APIData, len(apiDataRawArr))
		for i, apiDataRaw := range apiDataRawArr {
			apiDataMap, okay := apiDataRaw.(map[string]interface{})
			if !okay {
				return nil, fmt.Errorf("error parsing trade response: unknown object for API data answer.%d.ApiData", i)
			}
			apiDataArr[i] = getAPIDataFromMap(apiDataMap)
		}
	}
	return apiDataArr, nil
}

func getResultFromMap(result map[string]interface{}) *TradeConfirm {
	return &TradeConfirm{
		ID:                 result["ID"].(string),
		DealID:             result["DealID"].(string),
		OrderID:            result["OrderID"].(string),
		PositionExternalID: result["PositionExternalID"].(string),
		Retcode:            result["Retcode"].(string),
		ExternalRetcode:    result["ExternalRetcode"].(string),
		Volume:             result["Volume"].(string),
		VolumeExt:          result["VolumeExt"].(string),
		Price:              result["Price"].(string),
		PriceGateway:       result["PriceGateway"].(string),
		TickBid:            result["TickBid"].(string),
		TickAsk:            result["TickAsk"].(string),
		TickLast:           result["TickLast"].(string),
		Flags:              result["Flags"].(string),
		Comment:            result["Comment"].(string),
	}
}

func getAnswerFromMap(answer map[string]interface{}) *Trade {
	return &Trade{
		ID:                   answer["ID"].(string),
		Login:                answer["Login"].(string),
		ExternalAccount:      answer["ExternalAccount"].(string),
		Group:                answer["Group"].(string),
		Symbol:               answer["Symbol"].(string),
		Digits:               answer["Digits"].(string),
		Action:               answer["Action"].(string),
		TimeExpiration:       answer["TimeExpiration"].(string),
		Type:                 answer["Type"].(string),
		TypeFill:             answer["TypeFill"].(string),
		TypeTime:             answer["TypeTime"].(string),
		Flags:                answer["Flags"].(string),
		Volume:               answer["Volume"].(string),
		VolumeExt:            answer["VolumeExt"].(string),
		Order:                answer["Order"].(string),
		OrderExternalID:      answer["OrderExternalID"].(string),
		PriceOrder:           answer["PriceOrder"].(string),
		PriceTrigger:         answer["PriceTrigger"].(string),
		PriceSL:              answer["PriceSL"].(string),
		PriceTP:              answer["PriceTP"].(string),
		PriceDeviation:       answer["PriceDeviation"].(string),
		PriceDeviationTop:    answer["PriceDeviationTop"].(string),
		PriceDeviationBottom: answer["PriceDeviationBottom"].(string),
		Comment:              answer["Comment"].(string),
		ResultRetcode:        answer["ResultRetcode"].(string),
		ResultDealer:         answer["ResultDealer"].(string),
		ResultDeal:           answer["ResultDeal"].(string),
		ResultOrder:          answer["ResultOrder"].(string),
		ResultVolume:         answer["ResultVolume"].(string),
		ResultVolumeExt:      answer["ResultVolumeExt"].(string),
		ResultPrice:          answer["ResultPrice"].(string),
		ResultDealerBid:      answer["ResultDealerBid"].(string),
		ResultDealerAsk:      answer["ResultDealerAsk"].(string),
		ResultDealerLast:     answer["ResultDealerLast"].(string),
		ResultMarketBid:      answer["ResultMarketBid"].(string),
		ResultMarketAsk:      answer["ResultMarketAsk"].(string),
		ResultMarketLast:     answer["ResultMarketLast"].(string),
		ResultComment:        answer["ResultComment"].(string),
		IDClient:             answer["IDClient"].(string),
		IP:                   answer["IP"].(string),
		SourceLogin:          answer["SourceLogin"].(string),
		Position:             answer["Position"].(string),
		PositionBy:           answer["PositionBy"].(string),
		PositionExternalID:   answer["PositionExternalID"].(string),
		PositionByExternalID: answer["PositionByExternalID"].(string),
	}
}

func getAPIDataFromMap(apiData map[string]interface{}) *APIData {
	return &APIData{
		AppID:       apiData["AppID"].(string),
		ID:          apiData["ID"].(string),
		ValueInt:    apiData["ValueInt"].(string),
		ValueUInt:   apiData["ValueUInt"].(string),
		ValueDouble: apiData["ValueDouble"].(string),
	}
}
