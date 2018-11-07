package signature

import (
	"testing"
)

func TestGenerateSignature1(t *testing.T) {
	apiSecret := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	requestMethod := "GET"
	fullURLPath := `/api/v1/instrument`
	requestData := ``
	expireTimestamp := int64(1518064236)

	signature := GenerateSignature(apiSecret, requestMethod, fullURLPath, requestData, int64(expireTimestamp))

	if signature != "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00" {
		t.Log(signature)
		t.Error("TestGenerateSignature1 Failed")
	}
}

func TestGenerateSignature2(t *testing.T) {
	apiSecret := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	requestMethod := "GET"
	fullURLPath := `/api/v1/instrument?filter={"symbol": "XBTM15"}`
	requestData := ``
	expireTimestamp := int64(1518064237)

	signature := GenerateSignature(apiSecret, requestMethod, fullURLPath, requestData, int64(expireTimestamp))

	if signature != "e2f422547eecb5b3cb29ade2127e21b858b235b386bfa45e1c1756eb3383919f" {
		t.Log(signature)
		t.Error("TestGenerateSignature2 Failed")
	}
}

func TestGenerateSignature3(t *testing.T) {
	apiSecret := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	requestMethod := "POST"
	fullURLPath := `/api/v1/order`
	requestData := `{"symbol":"XBTM15","price":219.0,"clOrdID":"mm_bitmex_1a/oemUeQ4CAJZgP3fjHsA","orderQty":98}`
	expireTimestamp := int64(1518064238)

	signature := GenerateSignature(apiSecret, requestMethod, fullURLPath, requestData, expireTimestamp)

	if signature != "1749cd2ccae4aa49048ae09f0b95110cee706e0944e6a14ad0b3a8cb45bd336b" {
		t.Log(signature)
		t.Error("TestGenerateSignature3 Failed")
	}
}
