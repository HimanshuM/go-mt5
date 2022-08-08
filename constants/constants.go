package constants

import "time"

const (
	// Prefixes
	PrefixAPI    = "MT5WEBAPI%04x%04x"
	PacketFormat = "%04x%04x"

	// Configuration
	APIVersion  = 3270
	MetaSize    = 9
	MaxCommands = 16383

	// Authorization commands
	CmdAuthStart  = "AUTH_START"
	CmdAuthAnswer = "AUTH_ANSWER"

	// Ping command
	CmdPing  = "TEST_ACCESS"
	CmdClose = "QUIT"

	// Time commands
	CmdServerTime        = "TIME_SERVER"
	CmdServerTimeSetting = "TIME_GET"

	// Group commands
	CmdGroupTotal = "GROUP_TOTAL"
	CmdGroupIndex = "GROUP_NEXT"
	CmdGroupGet   = "GROUP_GET"

	// User commands
	CmdUserAdd = "USER_ADD"

	// Trade commands
	CmdTradeRequest = "DEALER_SEND"
	CmdTradeResult  = "DEALER_UPDATES"
	CmdTradeBalance = "TRADE_BALANCE"

	// Symbol commands
	CmdSymbolList = "SYMBOL_LIST"
	CmdSymbolGet  = "SYMBOL_GET"
	CmdSymbolNext = "SYMBOL_NEXT"

	// Tick commands
	CmdTickLast        = "TICK_LAST"
	CmdTickLastByGroup = "TICK_LAST_GROUP"
	CmdTickStats       = "TICK_STAT"

	/* PARAMETERS */
	// Common parameters
	ParamReturnCode = "RETCODE"
	ParamGroup      = "GROUP"
	ParamIndex      = "INDEX"
	ParamTotal      = "TOTAL"
	ParamID         = "ID"

	// Authorization parameters
	ParamAuthSrvRand       = "SRV_RAND"
	ParamAuthSrvRandAnswer = "SRV_RAND_ANSWER"
	ParamAuthCLIRand       = "CLI_RAND"
	ParamAuthCLIRandAnswer = "CLI_RAND_ANSWER"
	ParamAuthCryptRand     = "CRYPT_RAND"

	// Time parameters
	ParamServerTime              = "TIME"
	ParamServerTimeDaylight      = "Daylight"
	ParamServerTimeDaylightState = "DaylightState"
	ParamServerTimeTimezone      = "TimeZone"
	ParamServerTimeTimeServer    = "TimeServer"
	ParamServerTimeDays          = "Days"

	// User parameters
	ParamUserLogin          = "LOGIN"
	ParamUserLoginJSON      = "Login"
	ParamUserPassMain       = "PASS_MAIN"
	ParamUserPassInvestor   = "PASS_INVESTOR"
	ParamUserRights         = "RIGHTS"
	ParamUserGroup          = "GROUP"
	ParamUserName           = "NAME"
	ParamUserCompany        = "COMPANY"
	ParamUserLanguage       = "LANGUAGE"
	ParamUserCountry        = "COUNTRY"
	ParamUserCity           = "CITY"
	ParamUserState          = "STATE"
	ParamUserZipCode        = "ZIPCODE"
	ParamUserAddress        = "ADDRESS"
	ParamUserPhone          = "PHONE"
	ParamUserEmail          = "EMAIL"
	ParamUserID             = "ID"
	ParamUserStatus         = "STATUS"
	ParamUserComment        = "COMMENT"
	ParamUserColor          = "COLOR"
	ParamUserPassPhone      = "PASS_PHONE"
	ParamUserLeverage       = "LEVERAGE"
	ParamUserAgent          = "AGENT"
	ParamUserBalance        = "BALANCE"
	ParamUserRegitration    = "Registration"
	ParamUserLastAccess     = "LastAccess"
	ParamUserLastPassChange = "LastPassChange"

	// Trade parameters
	ParamTradeLogin       = "LOGIN"
	ParamTradeType        = "TYPE"
	ParamTradeBalance     = "BALANCE"
	ParamTradeComment     = "COMMENT"
	ParamTradeCheckMargin = "CHECK_MARGIN"
	ParamTradeTicket      = "TICKET"

	ParamTradeMarketOrder    = "200"
	ParamTradeLimitOrder     = "201"
	ParamTradeModifyPosition = "202"
	ParamTradeModifyOrder    = "202"
	ParamTradeCancelOrder    = "204"

	ParamTradeBuyOrder           = "0"
	ParamTradeSellOrder          = "1"
	ParamTradeBuyLimitOrder      = "2"
	ParamTradeSellLimitOrder     = "3"
	ParamTradeBuyStopOrder       = "4"
	ParamTradeSellStopOrder      = "5"
	ParamTradeBuyStopLimitOrder  = "6"
	ParamTradeSellStopLimitOrder = "7"
	ParamTradeCloseByOrder       = "8"

	ParamTradeResult = "result"
	ParamTradeAnswer = "answer"

	// Symbol parameters
	ParamSymbol = "SYMBOL"

	// Tick parameters
	ParamTickTransID = "TRANS_ID"

	/* CONSTANTS */
	// Trade constants
	ConstTradeBalance = 2

	// Misc
	WordAPI            = "WebAPI"
	WordManager        = "MANAGER"
	CryptMethodDefault = "AES256OFB"
	CryptMethodNone    = "NONE"
	// Defines how long a connection stays alive. Not used anywhere in the package
	KeeyAliveDuration = time.Second * 180
	// Defines how frequently ping command should be called. Available for use by package users. Not used anywhere in the package
	PingDuration = time.Second * 20
)
