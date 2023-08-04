package session

import (
	"github.com/google/uuid"
	"goLearn/web/jie"
)

// Manager 胶水，为了提升用户体验
type Manager struct {
	Propagator
	Store
	CtxSessKey string
}

func (m *Manager) GetSession(ctx *jie.Context) (Session, error) {
	if ctx.UserValues == nil {
		ctx.UserValues = make(map[string]any, 1)
	}
	val, ok := ctx.UserValues[m.CtxSessKey]
	if ok {
		return val.(Session), nil
	}

	sessID, err := m.Extract(ctx.Req)
	if err != nil {
		return nil, err
	}
	sess, err := m.Get(ctx.Req.Context(), sessID)
	if err != nil {
		return nil, err
	}
	ctx.UserValues[m.CtxSessKey] = sess
	return sess, err
}

func (m *Manager) InitSession(ctx *jie.Context) (Session, error) {
	id := uuid.New().String()
	sess, err := m.Generate(ctx.Req.Context(), id)
	if err != nil {

		return nil, err
	}
	// 注入HTTP响应里
	err = m.Inject(id, ctx.Resp)
	return sess, err
}

func (m *Manager) RemoveSession(ctx *jie.Context) error {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return err
	}
	err = m.Store.Remove(ctx.Req.Context(), sess.ID())
	if err != nil {
		return err
	}
	return m.Propagator.Remove(ctx.Resp)
}

func (m *Manager) RefreshSession(ctx *jie.Context) error {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return err
	}
	return m.Refresh(ctx.Req.Context(), sess.ID())
}
