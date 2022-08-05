package constants

import "time"

const (
	// Prefixes
	PREFIX_API    = "MT5WEBAPI%04x%04x"
	PACKET_FORMAT = "%04x%04x"

	// Configuration
	API_VERSION  = 3270
	META_SIZE    = 9
	MAX_COMMANDS = 16383

	// Authorization commands
	CMD_AUTH_START  = "AUTH_START"
	CMD_AUTH_ANSWER = "AUTH_ANSWER"

	// Ping command
	CMD_PING  = "TEST_ACCESS"
	CMD_CLOSE = "QUIT"

	// Time commands
	CMD_SERVER_TIME         = "TIME_SERVER"
	CMD_SERVER_TIME_SETTING = "TIME_GET"

	// Group commands
	CMD_GROUP_TOTAL = "GROUP_TOTAL"
	CMD_GROUP_INDEX = "GROUP_NEXT"
	CMD_GROUP_GET   = "GROUP_GET"

	// User commands
	CMD_USER_ADD = "USER_ADD"

	// Trade commands
	CMD_TRADE_BALANCE = "TRADE_BALANCE"

	// Symbol commands
	CMD_SYMBOL_LIST = "SYMBOL_LIST"
	CMD_SYMBOL_GET  = "SYMBOL_GET"
	CMD_SYMBOL_NEXT = "SYMBOL_NEXT"

	// Tick commands
	CMD_TICK_LAST          = "TICK_LAST"
	CMD_TICK_LAST_BY_GROUP = "TICK_LAST_GROUP"
	CMD_TICK_STATS         = "TICK_STAT"

	/* PARAMETERS */
	// Common parameters
	PARAM_RETURN_CODE = "RETCODE"
	PARAM_GROUP       = "GROUP"
	PARAM_INDEX       = "INDEX"
	PARAM_TOTAL       = "TOTAL"

	// Authorization parameters
	PARAM_AUTH_SRV_RAND        = "SRV_RAND"
	PARAM_AUTH_SRV_RAND_ANSWER = "SRV_RAND_ANSWER"
	PARAM_AUTH_CLI_RAND        = "CLI_RAND"
	PARAM_AUTH_CLI_RAND_ANSWER = "CLI_RAND_ANSWER"
	PARAM_AUTH_CRYPT_RAND      = "CRYPT_RAND"

	// Time parameters
	PARAM_SERVER_TIME               = "TIME"
	PARAM_SERVER_TIME_DAYLIGHT      = "Daylight"
	PARAM_SERVER_TIME_DAYLIGHTSTATE = "DaylightState"
	PARAM_SERVER_TIME_TIMEZONE      = "TimeZone"
	PARAM_SERVER_TIME_TIMESERVER    = "TimeServer"
	PARAM_SERVER_TIME_DAYS          = "Days"

	// User parameters
	PARAM_USER_LOGIN            = "LOGIN"
	PARAM_USER_LOGIN_JSON       = "Login"
	PARAM_USER_PASS_MAIN        = "PASS_MAIN"
	PARAM_USER_PASS_INVESTOR    = "PASS_INVESTOR"
	PARAM_USER_RIGHTS           = "RIGHTS"
	PARAM_USER_GROUP            = "GROUP"
	PARAM_USER_NAME             = "NAME"
	PARAM_USER_COMPANY          = "COMPANY"
	PARAM_USER_LANGUAGE         = "LANGUAGE"
	PARAM_USER_COUNTRY          = "COUNTRY"
	PARAM_USER_CITY             = "CITY"
	PARAM_USER_STATE            = "STATE"
	PARAM_USER_ZIPCODE          = "ZIPCODE"
	PARAM_USER_ADDRESS          = "ADDRESS"
	PARAM_USER_PHONE            = "PHONE"
	PARAM_USER_EMAIL            = "EMAIL"
	PARAM_USER_ID               = "ID"
	PARAM_USER_STATUS           = "STATUS"
	PARAM_USER_COMMENT          = "COMMENT"
	PARAM_USER_COLOR            = "COLOR"
	PARAM_USER_PASS_PHONE       = "PASS_PHONE"
	PARAM_USER_LEVERAGE         = "LEVERAGE"
	PARAM_USER_AGENT            = "AGENT"
	PARAM_USER_BALANCE          = "BALANCE"
	PARAM_USER_REGISTRATION     = "Registration"
	PARAM_USER_LAST_ACCESS      = "LastAccess"
	PARAM_USER_LAST_PASS_CHANGE = "LastPassChange"

	// Trade parameters
	PARAM_TRADE_LOGIN        = "LOGIN"
	PARAM_TRADE_TYPE         = "TYPE"
	PARAM_TRADE_BALANCE      = "BALANCE"
	PARAM_TRADE_COMMENT      = "COMMENT"
	PARAM_TRADE_CHECK_MARGIN = "CHECK_MARGIN"
	PARAM_TRADE_TICKET       = "TICKET"

	// Symbol parameters
	PARAM_SYMBOL = "SYMBOL"

	// Tick parameters
	PARAM_TICK_TRANS_ID = "TRANS_ID"

	/* CONSTANTS */
	// Trade constants
	CONST_TRADE_BALANCE = 2

	// Misc
	WORD_API             = "WebAPI"
	WORD_MANAGER         = "MANAGER"
	CRYPT_METHOD_DEFAULT = "AES256OFB"
	CRYPT_METHOD_NONE    = "NONE"
	// Defines how long a connection stays alive. Not used anywhere in the package
	KEEP_ALIVE_DURATION = time.Second * 180
	// Defines how frequently ping command should be called. Available for use by package users. Not used anywhere in the package
	PING_DURATION = time.Second * 20
)
