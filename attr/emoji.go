package attr

import (
	"log/slog"
)

func EmojiLevelFunc() func(groups []string, a slog.Attr) slog.Attr {

	fn := func(groups []string, a slog.Attr) slog.Attr {

		if a.Key == slog.LevelKey {

			// Handle custom level values.
			level := a.Value.Any().(slog.Level)

			// 🪵
			// 🧯

			switch {
			case level < LevelDebug:
				a.Value = slog.StringValue("TRACE ⁉️ ")
			case level < LevelInfo:
				a.Value = slog.StringValue("DEBUG 🔍")
			case level < LevelNotice:
				a.Value = slog.StringValue("INFO 💬")
			case level < LevelWarning:
				a.Value = slog.StringValue("WARNING ⚠️ ")
			case level < LevelError:
				a.Value = slog.StringValue("ERROR 🔥")
			case level < LevelEmergency:
				a.Value = slog.StringValue("EMERGENCY 💥")
			default:
				a.Value = slog.StringValue("EMERGENCY 💥")
			}
		}

		return a

	}

	return fn
}
