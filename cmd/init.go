package cmd

import "github.com/jsphbtst/babelfish/pkg/types"

type Globals struct {
	History      *types.HistoryJson
	Explanations *types.BreakdownJson
	Configs      *types.Configs
}

var globals Globals

func InitHistory(history *types.HistoryJson) {
	globals.History = history
}

func InitBreakdowns(breakdowns *types.BreakdownJson) {
	globals.Explanations = breakdowns
}

func InitConfigs(configs *types.Configs) {
	globals.Configs = configs
}
