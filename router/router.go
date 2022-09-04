package router

import "lcy-faas/function_module"

type Router struct {
	Paths map[string]function_module.FunctionModule
}

func (this *Router) Insert(path string, f function_module.FunctionModule) {
	this.Paths[path] = f
}

func (this *Router) Delete(path string) (err bool) {
	_, ok := this.Paths[path]
	if ok {
		delete(this.Paths, path)
		return true
	} else {
		return false
	}
}

func (this *Router) Find(path string) (f function_module.FunctionModule, err bool) {
	fn, ok := this.Paths[path]
	return fn, ok
}
