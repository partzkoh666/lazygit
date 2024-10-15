package controllers

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/tasks"
)

type DiffController struct {
	baseController
	c *ControllerCommon

	context      types.Context
	otherContext types.Context

	viewBufferManagerMap *map[string]*tasks.ViewBufferManager
}

var _ types.IController = &DiffController{}

func NewDiffController(
	c *ControllerCommon,
	context types.Context,
	otherContext types.Context,
	viewBufferManagerMap *map[string]*tasks.ViewBufferManager,
) *DiffController {
	return &DiffController{
		baseController:       baseController{},
		c:                    c,
		context:              context,
		otherContext:         otherContext,
		viewBufferManagerMap: viewBufferManagerMap,
	}
}

func (self *DiffController) GetKeybindings(opts types.KeybindingsOpts) []*types.Binding {
	return []*types.Binding{
		{
			Key:             opts.GetKey(opts.Config.Universal.TogglePanel),
			Handler:         self.TogglePanel,
			Description:     self.c.Tr.ToggleStagingView,
			Tooltip:         self.c.Tr.ToggleStagingViewTooltip,
			DisplayOnScreen: true,
		},
		{
			Key:         opts.GetKey(opts.Config.Universal.Return),
			Handler:     self.Escape,
			Description: self.c.Tr.ExitCustomPatchBuilder,
		},
	}
}

func (self *DiffController) Context() types.Context {
	return self.context
}

func (self *DiffController) GetMouseKeybindings(opts types.KeybindingsOpts) []*gocui.ViewMouseBinding {
	return []*gocui.ViewMouseBinding{}
}

func (self *DiffController) GetOnFocus() func(types.OnFocusOpts) {
	return func(opts types.OnFocusOpts) {
		self.c.Helpers().Diff.RenderFilesViewDiff(self.c.MainViewPairs().Diff)
		if opts.ClickedWindowName == "main" {
			if manager, ok := (*self.viewBufferManagerMap)[self.context.GetViewName()]; ok {
				// TODO: doesn't work the first time after launching. Need to
				// find a way to construct the ViewBufferManager for this view
				// earlier.
				manager.ReadLines(opts.ClickedViewLineIdx - self.context.GetView().LinesHeight() + 1)
			}
			self.context.GetView().FocusPoint(0, opts.ClickedViewLineIdx)
		}
	}
}

func (self *DiffController) GetOnFocusLost() func(types.OnFocusLostOpts) {
	return func(opts types.OnFocusLostOpts) {
	}
}

func (self *DiffController) TogglePanel() error {
	if self.otherContext.GetView().Visible {
		self.c.Context().Push(self.otherContext)
	}

	return nil
}

func (self *DiffController) Escape() error {
	self.c.Context().Pop()
	return nil
}
