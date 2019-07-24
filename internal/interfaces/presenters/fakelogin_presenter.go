package presenters

import "net/http"

// FakeLoginPresenter BadgerHole Presenter
type FakeLoginPresenter interface {
	Response(http.ResponseWriter, map[string]string)
}
