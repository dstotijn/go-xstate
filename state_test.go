package xstate_test

import (
	"testing"

	"github.com/dstotijn/go-xstate"
)

func TestMatchesState(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		parent xstate.StateValue
		child  xstate.StateValue
		want   bool
	}{
		{
			name: "equal atomic states",
			parent: xstate.StateValue{
				"foo": nil,
			},
			child: xstate.StateValue{
				"foo": nil,
			},
			want: true,
		},
		{
			name: "unequal atomic states",
			parent: xstate.StateValue{
				"foo": nil,
			},
			child: xstate.StateValue{
				"bar": nil,
			},
			want: false,
		},
		{
			name: "equal nested states",
			parent: xstate.StateValue{
				"foo": xstate.StateValue{
					"bar": nil,
				},
			},
			child: xstate.StateValue{
				"foo": xstate.StateValue{
					"bar": nil,
				},
			},
			want: true,
		},
		{
			name: "unequal nested states",
			parent: xstate.StateValue{
				"foo": xstate.StateValue{
					"bar": nil,
				},
			},
			child: xstate.StateValue{
				"foo": xstate.StateValue{
					"wrong": nil,
				},
			},
			want: false,
		},
		{
			name: "child is substate of parent",
			parent: xstate.StateValue{
				"foo": nil,
			},
			child: xstate.StateValue{
				"foo": xstate.StateValue{
					"bar": nil,
				},
				"baz": nil,
			},
			want: true,
		},
		{
			name: "parent is more specific than child",
			parent: xstate.StateValue{
				"foo": xstate.StateValue{
					"bar": nil,
				},
			},
			child: xstate.StateValue{
				"foo": nil,
			},
			want: false,
		},
		{
			name:   "parent is empty",
			parent: xstate.StateValue{},
			child: xstate.StateValue{
				"foo": nil,
			},
			want: true,
		},
		{
			name: "child is empty",
			parent: xstate.StateValue{
				"foo": nil,
			},
			child: xstate.StateValue{},
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := xstate.MatchesState(tt.parent, tt.child)
			if got != tt.want {
				t.Errorf("unexpected match result (want: %v, got: %v)", tt.want, got)
			}
		})
	}
}
