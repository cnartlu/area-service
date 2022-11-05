package app

type Option func(*App)

// ----------------------------------------------------------------
// ---------------------------start function-----------------------
// ----------------------------------------------------------------

func WithStartFunc(fn func() error) Option {
	return func(a *App) {
		a.start = fn
	}
}

// ----------------------------------------------------------------
// ---------------------------start flag---------------------------
// ----------------------------------------------------------------

// WithFlag 设置标签
func WithFlag(flag Flag) Option {
	return func(a *App) {
		a.flag = flag
	}
}

func WithFlagHelp(help bool) Option {
	return func(a *App) {
		a.flag.Help = help
	}
}

func WithFlagVersion(version bool) Option {
	return func(a *App) {
		a.flag.Version = version
	}
}

func WithFlagTest(test bool) Option {
	return func(a *App) {
		a.flag.Test = test
	}
}

func WithFlagSignal(signal string) Option {
	return func(a *App) {
		a.flag.Signal = signal
	}
}

func WithFlagConfig(config string) Option {
	return func(a *App) {
		a.flag.Config = config
	}
}

// ----------------------------------------------------------------
// ---------------------------end flag---------------------------
// ----------------------------------------------------------------
