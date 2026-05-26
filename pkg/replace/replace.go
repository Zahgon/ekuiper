package replace

var (
	replaceURL      = []string{"url"}
	replacePassword = []string{"saslPassword"}
	replaceDuration = []string{"cacheTtl", "timeout", "expiration", "interval"}
	replaceAction   = map[string]struct{}{
		"kafka": {},
		"sql":   {},
	}
)

func ReplaceRuleJson(ruleJson string, isTesting bool) string { _ = "STUB: not implemented"; return "" }

// each action only have 1 type

func WithDisableReplaceDburl() ReplacePropsOption {
	_ = "STUB: not implemented"
	return *new(ReplacePropsOption)
}

func WithDisableReplacePassword() ReplacePropsOption {
	_ = "STUB: not implemented"
	return *new(ReplacePropsOption)
}

func ReplacePropsDBURL(props map[string]interface{}) (bool, map[string]interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func ReplacePassword(props map[string]interface{}) (bool, map[string]interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func ReplaceDuration(props map[string]interface{}) (bool, map[string]interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func ReplacePropsWithPlug(plug string, props map[string]interface{}) (bool, map[string]interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func ReplacePropsWithOption(props map[string]interface{}, opts ...ReplacePropsOption) (bool, map[string]interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

type ReplacePropsOption func(c *ReplacePropsConfig)

type ReplacePropsConfig struct {
	DisableReplaceDbUrl    bool
	DisableReplacePassword bool
	DisableReplaceDuration bool
}

var passwordDict = map[string]struct{}{
	"password":      {},
	"pass":          {},
	"token":         {},
	"access_token":  {},
	"refresh_token": {},
}

func HidePassword(props map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func hide(props map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
