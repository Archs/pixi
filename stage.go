package pixi

type Stage struct {
	*DisplayObjectContainer
}

func NewStage(background uint32) Stage {
	stage := pkg.Get("Stage").New(background)
	return Stage{DisplayObjectContainer: wrapDisplayObjectContainer(stage)}
}
