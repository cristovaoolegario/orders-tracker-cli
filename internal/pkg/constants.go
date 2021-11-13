package pkg

// CorreiosBaseURL correios API base URL
const CorreiosBaseURL = "https://correios.contrateumdev.com.br/api/rastreio"

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
