package limiter

import (
	"sync"
	"time"
)

type storageItem struct {
	start time.Time
	count int
	LimiterParams
}

type mapLimiter struct {
	storage map[string]storageItem
	mtx     sync.Mutex
}

func NewMapLimiter() Limiter {
	return &mapLimiter{
		storage: make(map[string]storageItem),
		mtx:     sync.Mutex{},
	}
}

func (m *mapLimiter) newStorageItem(key string, start time.Time, params LimiterParams) {
	m.storage[key] = storageItem{
		start:         start,
		count:         1,
		LimiterParams: params,
	}
}

func (m *mapLimiter) Check(key string, timeStamp time.Time, params LimiterParams) (bool, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	item, ok := m.storage[key]

	if ok {
		//out of interval, clear counter
		if timeStamp.Sub(item.start) > params.TimeInterval {
			m.newStorageItem(key, timeStamp, params)
			return true, nil
		}
		//limit is over
		if item.count >= params.CountLimit {
			return false, nil
		}
		//increase counter
		item.count++
		m.storage[key] = item
		return true, nil
	}
	//create new item
	m.newStorageItem(key, timeStamp, params)
	return true, nil
}
