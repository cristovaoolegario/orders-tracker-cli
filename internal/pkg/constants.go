package pkg

import "encoding/base64"

// CorreiosValidationUrl correios API validation URL
const CorreiosValidationUrl = "https://proxyapp.correios.com.br/v1/app-validation"

// ValidationData body that will be use in the autentication request
var ValidationData = []byte(`{"requestToken":"` + base64.StdEncoding.EncodeToString([]byte("android;br.com.correios.preatendimento;F32E29976709359859E0B97F6F8A483B9B953578")) + `"}`)

// CorreiosBaseURL correios API base URL
const CorreiosBaseURL = "https://proxyapp.correios.com.br/v1/sro-rastro/%s"

// IconDictionary represents a dictionary that map dto.Event Code to Icons
var IconDictionary = map[string]string{
	"BDE01": "🎁",
	"BDE20": "📪",
	"OEC01": "🙌",
	"DO01":  "🚚",
	"RO01":  "🚚",
	"PO01":  "📦",
	"PO09":  "💤",
	"PAR10": "✅",
	"PAR16": "🛬",
	"PAR17": "💸",
	"PAR18": "🗺",
	"PAR21": "🔎",
	"PAR24": "🔙",
	"PAR26": "🙅",
	"PAR31": "🤑",
	"":      "🚧",
}
