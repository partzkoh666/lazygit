package controllers

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type DiffController struct {
	baseController
	c *ControllerCommon

	context      types.Context
	otherContext types.Context
}

var _ types.IController = &DiffController{}

func NewDiffController(
	c *ControllerCommon,
	context types.Context,
	otherContext types.Context,
) *DiffController {
	return &DiffController{
		baseController: baseController{},
		c:              c,
		context:        context,
		otherContext:   otherContext,
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
