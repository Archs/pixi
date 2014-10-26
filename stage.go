package pixi

import "github.com/gopherjs/gopherjs/js"

type Stage struct {
	*DisplayObjectContainer
}

func NewStage(background uint32) *Stage {
	return wrapStage(pkg.Get("Stage").New(background))
}

func wrapStage(object js.Object) *Stage {
	return &Stage{DisplayObjectContainer: wrapDisplayObjectContainer(object)}
}
