package filesystem

type HookFunc func()

type HooksChain []HookFunc

type Hook map[string]HooksChain

func (h HooksChain) Len() int {
	return len(h)
}

func (h HooksChain) Use(key string, hooks ...HookFunc) {

}
