package registry

import (
	"context"
	"fmt"
	"sync"
)

//插件管理类
//* 可以用一个大map管理，key 字符串， value是 registry 接口对象
//* 用户自定义去调用，自定义插件
//* 实现注册中心的初始化，供系统使用

type PluginMgr struct {
	// map维护所有的插件
	plugins map[string]Registry
	lock    sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// 插件注册
func RegistryPlugin(registry Registry) (err error) {
	return pluginMgr.registryPlugin(registry)
}

// 注册插件
func (p *PluginMgr) registryPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	//先去看里面有没有
	if _, ok := p.plugins[plugin.Name()]; ok {
		err = fmt.Errorf("plugin named %s already exists", plugin.Name())
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegister(ctx, name, opts...)
}

func (p *PluginMgr) initRegister(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	plugin, ok := p.plugins[name]
	if ok {
		err = fmt.Errorf("plugin named %s not exists", name)
		return nil, err
	}
	// 存在，返回值赋值
	registry = plugin
	// 进行组件初始化
	err = plugin.Init(ctx, opts...)
	return
}
