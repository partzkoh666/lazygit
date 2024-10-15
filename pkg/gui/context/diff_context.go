package context

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type DiffContext struct {
	*SimpleContext
	*SearchTrait

	c *ContextCommon
}

var _ types.ISearchableContext = (*DiffContext)(nil)

func NewDiffContext(
	view *gocui.View,
	windowName string,
	key types.ContextKey,

	c *ContextCommon,
) *DiffContext {
	ctx := &DiffContext{
		SimpleContext: NewSimpleContext(
			NewBaseContext(NewBaseContextOpts{
				Kind:             types.MAIN_CONTEXT,
				View:             view,
				WindowName:       windowName,
				Key:              key,
				Focusable:        true,
				HighlightOnFocus: true,
			})),
		SearchTrait: NewSearchTrait(c),
		c:           c,
	}

	// TODO: copied from PatchExplorerContext. Do we need something like this?
	// ctx.GetView().SetOnSelectItem(ctx.SearchTrait.onSelectItemWrapper(
	// 	func(selectedLineIdx int) error {
	// 		ctx.GetMutex().Lock()
	// 		defer ctx.GetMutex().Unlock()
	// 		ctx.NavigateTo(ctx.c.Context().IsCurrent(ctx), selectedLineIdx)
	// 		return nil
	// 	}),
	// )

	return ctx
}

func (self *DiffContext) ModelSearchResults(searchStr string, caseSensitive bool) []gocui.SearchPosition {
	return nil
}
