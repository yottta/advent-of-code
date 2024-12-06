package state_machine

import (
	"bytes"
)

type StateMachineCmd int

const (
	StateMachineCmdReset = iota
	StateMachineCmdStoreToBuffer
	StateMachineCmdWriteAndReset
	StateMachineCmdLock
	StateMachineCmdUnlock
)

type StateMachine struct {
	startMatcher MatcherFunc
	nextMatcher  MatcherFunc

	buf bytes.Buffer
	res []string

	locked bool
}

func New(rules MatcherFunc, locked bool) *StateMachine {
	return &StateMachine{
		startMatcher: rules,
		locked:       locked,
	}
}

type MatcherFunc func(r rune) (StateMachineCmd, MatcherFunc)

func Or(matchers ...MatcherFunc) MatcherFunc {
	return func(r rune) (StateMachineCmd, MatcherFunc) {
		for _, m := range matchers {
			if res, next := m(r); res == 1 {
				return res, next
			}
		}
		return StateMachineCmdReset, nil
	}
}

func RuneMatcher(exp rune, next MatcherFunc) MatcherFunc {
	return func(r rune) (StateMachineCmd, MatcherFunc) {
		if exp == r {
			return StateMachineCmdStoreToBuffer, next
		}
		return StateMachineCmdReset, nil
	}
}

func LockOnRune(exp rune) MatcherFunc {
	return func(r rune) (StateMachineCmd, MatcherFunc) {
		if exp == r {
			return StateMachineCmdLock, nil
		}
		return StateMachineCmdReset, nil
	}
}

func StopOnRune(exp rune) MatcherFunc {
	return func(r rune) (StateMachineCmd, MatcherFunc) {
		if exp == r {
			return StateMachineCmdWriteAndReset, nil
		}
		return StateMachineCmdReset, nil
	}
}

func UnlockOnRune(exp rune) MatcherFunc {
	return func(r rune) (StateMachineCmd, MatcherFunc) {
		if exp == r {
			return StateMachineCmdUnlock, nil
		}
		return StateMachineCmdReset, nil
	}
}

func (sm *StateMachine) Consume(r rune) StateMachineCmd {
	if sm.nextMatcher == nil {
		sm.nextMatcher = sm.startMatcher
	}
	res, next := sm.nextMatcher(r)
	switch res {
	case StateMachineCmdReset:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	case StateMachineCmdStoreToBuffer:
		sm.buf.WriteRune(r)
		if next != nil {
			sm.nextMatcher = next
		}
	case StateMachineCmdWriteAndReset:
		sm.buf.WriteRune(r)
		if sm.locked {
			sm.nextMatcher = sm.startMatcher
			sm.buf.Reset()
			break
		}
		sm.res = append(sm.res, sm.buf.String())
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	case StateMachineCmdLock:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
		sm.locked = true
	case StateMachineCmdUnlock:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
		sm.locked = false
	}
	return res
}

func (sm *StateMachine) Results() []string {
	return sm.res
}
