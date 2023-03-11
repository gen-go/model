package dbstruct

import (
	"time"

	"gorm.io/gorm"
)

// StockLimit struct
type StockLimit struct {
	ID       uint64  `gorm:"->:false;<-:create;column:id;primaryKey;autoIncrement;" json:"-"`
	Day      string  `gorm:"type:varchar(32);not null;uniqueIndex:uni_day_code;comment:日期" json:"day"`
	Code     string  `gorm:"type:varchar(6);not null;index;uniqueIndex:uni_day_code;" json:"code"` // stock_code
	Name     string  `gorm:"type:varchar(32);not null;" json:"name"`
	RiseRate float64 `gorm:"type:decimal(5,2);not null;default:0;comment:涨幅" json:"rise_rate"`
	//HighPrice             float64        `gorm:"type:decimal(10,2);not null;default:0;comment:最高价" json:"high_price"`
	//LowPrice              float64        `gorm:"type:decimal(10,2);not null;default:0;comment:最低价" json:"low_price"`
	//OpenPrice             float64        `gorm:"type:decimal(10,2);not null;default:0;comment:开盘价" json:"open_price"`
	//ClosePrice            float64        `gorm:"type:decimal(10,2);not null;default:0;comment:收盘价" json:"close_price"`
	MainCapitalFlow       float64 `gorm:"type:decimal(10,2);not null;default:0;comment:主力资金流向(单位万)" json:"main_capital_flow"`
	CirculateMarketValue  float64 `gorm:"type:decimal(10,2);not null;default:0;comment:流通市值(单位亿)" json:"circulate_market_value"`
	TotalMarketValue      float64 `gorm:"type:decimal(10,2);not null;default:0;comment:总市值(单位亿)" json:"total_market_value"`
	Reason                string  `gorm:"type:varchar(255);not null;default:'';comment:涨停原因" json:"reason"`
	Industry              string  `gorm:"type:varchar(255);not null;default:'';comment:行业" json:"industry"`
	DaysLimits            string  `gorm:"type:varchar(30);not null;default:'';comment:几天几板" json:"days_limits"`
	ContinuousDays        int     `gorm:"type:tinyint;not null;default:0;comment:连续涨停天数" json:"continuous_days"`
	LimitType             string  `gorm:"type:varchar(30);not null;default:'';comment:涨停类型" json:"limit_type"`
	FirstLimitTime        string  `gorm:"type:varchar(20);not null;default:'';comment:首次涨停时间" json:"first_limit_time"`
	FinalLimitTime        string  `gorm:"type:varchar(20);not null;default:'';comment:最终涨停时间" json:"final_limit_time"`
	OpenTimes             int     `gorm:"type:tinyint;not null;default:0;comment:涨停开板次数" json:"open_times"`
	TurnoverRatio         float64 `gorm:"type:decimal(5,2);not null;default:0;comment:换手率" json:"turnover_ratio"`
	Turnover              float64 `gorm:"type:decimal(10,2);not null;default:0;comment:成交额(单位亿)" json:"turnover"`
	Volume                float64 `gorm:"type:decimal(10,2);not null;default:0;comment:成交量(单位万手)" json:"volume"`
	SealTurnover          float64 `gorm:"type:decimal(10,2);not null;default:0;comment:涨停封单额(单位亿)" json:"rise_rate"`
	SealVolume            float64 `gorm:"type:decimal(10,2);not null;default:0;comment:涨停封单量(单位万手)" json:"rise_rate"`
	SealVolumeByCirculate float64 `gorm:"type:decimal(10,2);not null;default:0;comment:涨停封单量占流通a股比" json:"rise_rate"`
	SealVolumeByVolume    float64 `gorm:"type:decimal(10,2);not null;default:0;comment:涨停封单量占成交量比" json:"rise_rate"`
	MarketCode            string  `gorm:"type:varchar(6);not null;" json:"market_code"` // market_code
	StockCode             string  `gorm:"type:varchar(10);not null;" json:"stock_code"` // stock_code

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName 表名
func (StockLimit) TableName() string {
	return "stock_limit"
}

func init() {
	set("StockLimit", StockLimit{})
}
