package framework

type IGroup interface {
	// HttpMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// 嵌套group
	Group(string) IGroup
}

type Group struct {
	core   *Core
	parent *Group
	prefix string //组名
}

// 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

// 实现Post方法
func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

// 实现Put方法
func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

// 实现Delete方法
func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

// 递归获取组名
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.getAbsolutePrefix() + g.prefix
}

// 实现 Group 方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
