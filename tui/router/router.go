// Package router defines the navigation primitives for the TUI.
//
// Every screen implements Page. When a key is pressed, Update returns
// an Action that tells the app what to do next:
//
//   Stay{p}  – stay on this screen (p is the updated page state)
//   Push{p}  – navigate into a sub-page (adds to the stack)
//   Pop{}    – go back (removes top of stack)
//   Quit{}   – exit the app
package router

// Page is the interface every screen must implement.
type Page interface {
	View() string
	Update(key string) Action
}

// Action is returned by Page.Update to tell the router what to do.
type Action interface{ isAction() }

type (
	Stay struct{ Page } // update current page state
	Push struct{ Page } // navigate into a sub-page
	Pop  struct{}       // go back
	Quit struct{}       // exit the app
)

func (Stay) isAction() {}
func (Push) isAction() {}
func (Pop) isAction()  {}
func (Quit) isAction() {}
