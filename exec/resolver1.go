package exec

// MultiResolver chains multiple Resolvers, symbol looking up is according to the order of resolvers.
// The first found symbol will be returned.
type MultiResolver []Resolver

// NewMultiResolver instance a MultiResolver from resolves
func NewMultiResolver(resolvers ...Resolver) MultiResolver {
	return resolvers
}

// ResolveFunc implements Resolver interface
func (m MultiResolver) ResolveFunc(module, name string) (interface{}, bool) {
	for _, r := range m {
		if f, ok := r.ResolveFunc(module, name); ok {
			return f, true
		}
	}
	return nil, false
}

// ResolveGlobal implements Resolver interface
func (m MultiResolver) ResolveGlobal(module, name string) (int64, bool) {
	for _, r := range m {
		if v, ok := r.ResolveGlobal(module, name); ok {
			return v, true
		}
	}
	return 0, false
}

type importFunc struct {
	module, name string
	body         interface{}
}

type resolverBridge struct {
	resolver Resolver
	funcmap  map[string]int
	funcs    []importFunc
}

func newResolverBridge(r Resolver) *resolverBridge {
	return &resolverBridge{
		resolver: r,
		funcmap:  make(map[string]int),
		funcs:    make([]importFunc, 1),
	}
}
