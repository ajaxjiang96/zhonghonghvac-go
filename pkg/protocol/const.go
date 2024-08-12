package protocol

import "fmt"

type FanSpeed byte // 风速，B19 and B27

const (
	FanSpeedAuto FanSpeed = 0x00 // 自动
	FanSpeedLow  FanSpeed = 0x04 // 低速
	FanSpeedMid  FanSpeed = 0x02 // 中速
	FanSpeedHigh FanSpeed = 0x01 // 高速

	// 以下为非常规风速
	FanSpeedMidHigh FanSpeed = 0x03 // 中高速
	FanSpeedMidLow  FanSpeed = 0x05 // 中低速
	FanSpeedBreeze  FanSpeed = 0x06 // 微风
	FanSpeedTurbo   FanSpeed = 0x07 // 超强
	FanSpeedStop    FanSpeed = 0x08 // 停止
)

func (fs FanSpeed) String() string {
	switch fs {
	case FanSpeedAuto:
		return "自动"
	case FanSpeedLow:
		return "低速"
	case FanSpeedMid:
		return "中速"
	case FanSpeedHigh:
		return "高速"
	case FanSpeedMidHigh:
		return "中高速"
	case FanSpeedMidLow:
		return "中低速"
	case FanSpeedBreeze:
		return "微风"
	case FanSpeedTurbo:
		return "超强"
	case FanSpeedStop:
		return "停止"
	default:
		return "未知"
	}
}

type ACModeB19 byte // 空调模式，B19

const (
	ACModeB19Cooling    ACModeB19 = 0x01 // 设定制冷
	ACModeB19Dehumidify ACModeB19 = 0x02 // 设定除湿
	ACModeB19Ventilate  ACModeB19 = 0x04 // 设定送风
	ACModeB19Heating    ACModeB19 = 0x08 // 设定制热

	// 以下为非常规模式
	ACModeB19Fresh          ACModeB19 = 0x03 // 设定清爽
	ACModeB19AutoDehumidify ACModeB19 = 0x05 // 设定自动除湿
	ACModeB19Sleep          ACModeB19 = 0x06 // 设定贴心睡眠
	ACModeB19FloorHeating   ACModeB19 = 0x09 // 设定地暖
	ACModeB19TurboHeating   ACModeB19 = 0x0A // 设定强热（地暖和制热同时开启）
)

func (am ACModeB19) String() string {
	switch am {
	case ACModeB19Heating:
		return "制热"
	case ACModeB19Cooling:
		return "制冷"
	case ACModeB19Ventilate:
		return "送风"
	case ACModeB19Dehumidify:
		return "除湿"
	case ACModeB19Fresh:
		return "清爽"
	case ACModeB19AutoDehumidify:
		return "自动除湿"
	case ACModeB19Sleep:
		return "贴心睡眠"
	case ACModeB19FloorHeating:
		return "地暖"
	case ACModeB19TurboHeating:
		return "强热"
	default:
		return fmt.Sprintf("未知(%d)", am)
	}
}

type ACModeB27 byte // 空调模式，B27

const (
	ACModeB27Heating    ACModeB27 = 0x01 // 设定制热
	ACModeB27Cooling    ACModeB27 = 0x02 // 设定制冷
	ACModeB27Ventilate  ACModeB27 = 0x04 // 设定送风
	ACModeB27Dehumidify ACModeB27 = 0x08 // 设定除湿
)

func (am ACModeB27) String() string {
	switch am {
	case ACModeB27Heating:
		return "制热"
	case ACModeB27Cooling:
		return "制冷"
	case ACModeB27Ventilate:
		return "送风"
	case ACModeB27Dehumidify:
		return "除湿"
	default:
		return fmt.Sprintf("未知(%d)", am)
	}
}

type ACMode byte // 空调模式，B19 and B27

const (
	ACModeHeating    ACMode = iota // 设定制热
	ACModeCooling                  // 设定制冷
	ACModeVentilate                // 设定送风
	ACModeDehumidify               // 设定除湿
	ACModeUnknown    ACMode = 0xFF // 未知或不支持
)

func (am ACMode) String() string {
	switch am {
	case ACModeHeating:
		return "制热"
	case ACModeCooling:
		return "制冷"
	case ACModeVentilate:
		return "送风"
	case ACModeDehumidify:
		return "除湿"
	default:
		return fmt.Sprintf("未知(%d)", am)
	}
}

func (am ACMode) ToB19() ACModeB19 {
	switch am {
	case ACModeHeating:
		return ACModeB19Heating
	case ACModeCooling:
		return ACModeB19Cooling
	case ACModeVentilate:
		return ACModeB19Ventilate
	case ACModeDehumidify:
		return ACModeB19Dehumidify
	default:
		return ACModeB19Cooling // 默认为制冷
	}
}

func (am ACMode) ToB27() ACModeB27 {
	switch am {
	case ACModeHeating:
		return ACModeB27Heating
	case ACModeCooling:
		return ACModeB27Cooling
	case ACModeVentilate:
		return ACModeB27Ventilate
	case ACModeDehumidify:
		return ACModeB27Dehumidify
	default:
		return ACModeB27Cooling // 默认为制冷
	}
}
func ACModeFromB19(b19 ACModeB19) ACMode {
	switch b19 {
	case ACModeB19Heating:
		return ACModeHeating
	case ACModeB19Cooling:
		return ACModeCooling
	case ACModeB19Ventilate:
		return ACModeVentilate
	case ACModeB19Dehumidify:
		return ACModeDehumidify
	default:
		return ACModeUnknown
	}
}

func ACModeFromB27(b27 ACModeB27) ACMode {
	switch b27 {
	case ACModeB27Heating:
		return ACModeHeating
	case ACModeB27Cooling:
		return ACModeCooling
	case ACModeB27Ventilate:
		return ACModeVentilate
	case ACModeB27Dehumidify:
		return ACModeDehumidify
	default:
		return ACModeUnknown
	}
}

type ACStatus byte // 空调状态，B19 and B27

const (
	ACStatusOnline    ACStatus = 0x01 // 在线
	ACStatusOffline   ACStatus = 0x02 // 离线
	ACStatusSearching ACStatus = 0x03 // 搜索中
)

func (as ACStatus) String() string {
	switch as {
	case ACStatusOnline:
		return "在线"
	case ACStatusOffline:
		return "离线"
	case ACStatusSearching:
		return "搜索中"
	default:
		return fmt.Sprintf("未知(%d)", as)
	}
}

type ACBrand byte // 空调品牌，B19 and B27

const (
	ACBrandHitachi            ACBrand = iota + 1 // 日立多联机（2 芯）
	ACBrandDaikin                                // 大金多联机（2 芯）
	ACBrandToshiba                               // 东芝多联机（2 芯）
	ACBrandMitsubishiHeavy                       // 三菱重工多联机（2 芯）
	ACBrandMitsubishiElectric                    // 三菱电机多联机（4 芯）
	ACBrandGree                                  // 格力多联机（2 芯）
	ACBrandHisense                               // 海信多联机（2 芯）
	_
	ACBrandHaier // 海尔多联机（3 芯）
	_
	_
	_
	_
	_
	ACBrandPanasonic // 松下多联机（2 芯）
	ACBrandYork      // 约克多联机（2 芯）
	_
	ACBrandToshibaDeducted   // 东芝风管机（4 芯）
	ACBrandPanasonicDeducted // 松下风管机（4 芯）
	ACBrandMideaW1W2         // 美的 W1W2（2 芯）
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ACBrandHitachiDeducted // 日立风管机（4 芯）
	_
	ACBrandGreeDeducted4Core // 格力风管机（4 芯）
	ACBrandGreeDeducted2Core // 格力风管机（2 芯）
	ACBrandMideaCN40         // 美的 CN40（2 芯）
	ACBrandDaikinDeductedMX  // 大金MX风管机（4 芯）
	ACBrandHaierDeducted     // 海尔风管机（3 芯）
	_
	ACBrandHisenseDeducted        // 海信风管机（4 芯）
	ACBrandMitsubishiHeavy3Core   // 三菱重工多联机（3 芯）
	ACBrandHaierDeducted2Remote   // 海尔风管机（3 芯）
	ACBrandCarrierDeducted        // 开利风管机（4 芯）
	ACBrandMideaCN20              // 美的 CN20（5 芯）
	ACBrandMideaCN40Bidirectional // 美的 CN40（双向共存）
	ACBrandMideaX1X2              // 美的 X1X2（2 芯）
	ACBrandMideaColmo             // 美的 Colmo（2 芯）
	_
	ACBrandFujitsuDeducted // 富士通风管机（3 芯）
	ACBrandEKDeducted      // 欧科风管机（4 芯）
	ACBrandYorkGuangzhou   // 广州约克多联机（4 芯）
	ACBrandYorkDeducted    // 约克风管机（4 芯）
	_
	_
	ACBrandPanasonicMountedHongKong      // 松下壁挂香港（4 芯）
	ACBrandEmulator                 = 88 // 模拟器
	ACBrandACTest                   = 99 // 空调Test
)

type ACWindDir byte // 风向

const (
	ACWindDirNoDir ACWindDir = 0x00 // 无风向
	ACWindDirPos1  ACWindDir = 0x01 // 位置1
	ACWindDirPos2  ACWindDir = 0x02 // 位置2
	ACWindDirPos3  ACWindDir = 0x03 // 位置3
	ACWindDirPos4  ACWindDir = 0x04 // 位置4
	ACWindDirPos5  ACWindDir = 0x05 // 位置5
	ACWindDirPos6  ACWindDir = 0x06 // 位置6
	ACWindDirPos7  ACWindDir = 0x07 // 位置7
	ACWindDirAuto  ACWindDir = 0xFF // 自动摆动
)

func (wd ACWindDir) String() string {
	switch wd {
	case ACWindDirNoDir:
		return "无风向"
	case ACWindDirPos1:
		return "位置1"
	case ACWindDirPos2:
		return "位置2"
	case ACWindDirPos3:
		return "位置3"
	case ACWindDirPos4:
		return "位置4"
	case ACWindDirPos5:
		return "位置5"
	case ACWindDirPos6:
		return "位置6"
	case ACWindDirPos7:
		return "位置7"
	case ACWindDirAuto:
		return "自动摆动"
	default:
		return fmt.Sprintf("未知(%d)", wd)
	}
}

type FuncCode byte

const (
	// Bit access
	FuncCodeReadGateway                        FuncCode = 0xB0
	FuncCodeEditGateway                        FuncCode = 0xB1
	FuncCodeGatewayOnOff                       FuncCode = 0x31
	FuncCodeGatewayTemp                        FuncCode = 0x32
	FuncCodeGatewayControl                     FuncCode = 0x33
	FuncCodeGatewayWindSpeed                   FuncCode = 0x34
	FuncCodeACStatus                           FuncCode = 0x50
	FuncCodeGatewayWindDir                     FuncCode = 0x71
	FuncCodeGatewayNewAirOnOff                 FuncCode = 0x72
	FuncCodeGatewayNewAirMode                  FuncCode = 0x73
	FuncCodeGatewayNewAirSpeed                 FuncCode = 0x74
	FuncCodeGatewayFloorHeatingOnOff           FuncCode = 0x81
	FuncCodeGatewayFloorHeatingTemp            FuncCode = 0x82
	FuncCodeGatewayFloorHeatingControl         FuncCode = 0x83
	FuncCodeGatewayFloorHeatingAntiFreezeOnOff FuncCode = 0x84
	FuncCodePerformanceCheck                   FuncCode = 0x01
	FuncCodeStatusCheck                        FuncCode = 0x02
	FuncCodeOnOff                              FuncCode = 0x03
	FuncCodeControl                            FuncCode = 0x03
	FuncCodeErrorCheck                         FuncCode = 0x04
	FuncCodeFreshAirStatus                     FuncCode = 0x11
	FuncCodeFreshAirPerformance                FuncCode = 0x12
	FuncCodeFreshAirControl                    FuncCode = 0x13
	FuncCodeFreshAirErrorCheck                 FuncCode = 0x14
	FuncCodeFloorHeatingPerformance            FuncCode = 0x21
	FuncCodeFloorHeatingStatusCheck            FuncCode = 0x22
	FuncCodeFloorHeatingControlCheck           FuncCode = 0x24
	FuncCodeFloorHeatingOnOff                  FuncCode = 0x23
)

const (
	HeadCode            byte = 0xDD
	HeadCodeGateway     byte = 0xDD
	HeadCodeReadGateway byte = 0xFF
	ON                  byte = 0x01
	OFF                 byte = 0x00
)

type Validation uint

const (
	ValidationNone Validation = iota
	ValidationOdd
	ValidationEven
)
