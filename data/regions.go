package data

type region struct {
	Name, Letter string
}

var Regions = map[string]region{
	"01": region{"Stockholms län", "AB"},
	"03": region{"Uppsala län", "C"},
	"04": region{"Södermanlands län", "D"},
	"05": region{"Östergötlands län", "E"},
	"06": region{"Jönköpings län", "F"},
	"07": region{"Kronobergs län", "G"},
	"08": region{"Kalmar län", "H"},
	"09": region{"Gotlands län", "I"},
	"10": region{"Blekinge län", "K"},
	"12": region{"Skåne län", "M"},
	"13": region{"Hallands län", "N"},
	"14": region{"Västra Götalands län", "O"},
	"17": region{"Värmlands län", "S"},
	"18": region{"Örebro län", "T"},
	"19": region{"Västmanlands län", "U"},
	"20": region{"Dalarnas län", "W"},
	"21": region{"Gävleborgs län", "X"},
	"22": region{"Västernorrlands län", "Y"},
	"23": region{"Jämtlands län", "Z"},
	"24": region{"Västerbottens län", "AC"},
	"25": region{"Norrbottens län", "BD"},
}
