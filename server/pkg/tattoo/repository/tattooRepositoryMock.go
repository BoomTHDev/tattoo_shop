// server/pkg/tattoo/repository/tattooRepositoryMock.go

package repository

import (
	"sync"
	"time"

	"github.com/BoomTHDev/tattoo_port/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tattooRepositoryMock struct {
	mu   sync.RWMutex
	data map[uuid.UUID]entities.Tattoo
}

// NewTattooRepositoryMock สร้าง mock repo ว่าง ๆ
func NewTattooRepositoryMock() TattooRepository {
	return &tattooRepositoryMock{
		data: make(map[uuid.UUID]entities.Tattoo),
	}
}

// NewTattooRepositoryMockWithSeed สร้าง mock repo พร้อมข้อมูลเริ่มต้น
func NewTattooRepositoryMockWithSeed() TattooRepository {
	r := &tattooRepositoryMock{
		data: make(map[uuid.UUID]entities.Tattoo),
	}

	seed := []entities.Tattoo{
		{
			ID:    uuid.New(),
			Title: "Traditional Dragon Sleeve",
			ImageURL: []string{
				"https://images.unsplash.com/photo-1625314868143-20c4a50e26ac?q=80&w=800",
				"https://images.unsplash.com/photo-1616422287073-5b8961a72b8a?q=80&w=800",
			},
			CreatedAt: time.Now().AddDate(0, -1, 0),
			UpdatedAt: time.Now().AddDate(0, 0, -10),
		},
		{
			ID:    uuid.New(),
			Title: "Minimalist Line Art",
			ImageURL: []string{
				"https://images.unsplash.com/photo-1595526114035-efad3f9c49a6?q=80&w=800",
			},
			CreatedAt: time.Now().AddDate(0, -2, 0),
			UpdatedAt: time.Now().AddDate(0, 0, -5),
		},
		{
			ID:    uuid.New(),
			Title: "Japanese Koi Fish",
			ImageURL: []string{
				"https://images.unsplash.com/photo-1618221299795-d748f0ec5df8?q=80&w=800",
				"https://images.unsplash.com/photo-1604654894613-9db12e3d4df4?q=80&w=800",
				"https://images.unsplash.com/photo-1599442494544-2b43f5201fa2?q=80&w=800",
			},
			CreatedAt: time.Now().AddDate(0, -3, 0),
			UpdatedAt: time.Now(),
		},
	}

	for _, t := range seed {
		r.data[t.ID] = t
	}

	return r
}

func (r *tattooRepositoryMock) Create(tattoo *entities.Tattoo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// ถ้าไม่ได้ set ID มา ให้สร้างใหม่
	if tattoo.ID == uuid.Nil {
		tattoo.ID = uuid.New()
	}

	now := time.Now()
	if tattoo.CreatedAt.IsZero() {
		tattoo.CreatedAt = now
	}
	tattoo.UpdatedAt = now

	// เก็บลง map (ทับค่าเดิมถ้า key ซ้ำ)
	r.data[tattoo.ID] = *tattoo
	return nil
}

func (r *tattooRepositoryMock) GetAll() ([]entities.Tattoo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]entities.Tattoo, 0, len(r.data))
	for _, t := range r.data {
		out = append(out, t)
	}
	return out, nil
}

func (r *tattooRepositoryMock) GetById(id string) (*entities.Tattoo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	uid, err := uuid.Parse(id)
	if err != nil {
		// ให้พฤติกรรมเหมือน gorm หาไม่เจอ
		return nil, gorm.ErrRecordNotFound
	}

	if t, ok := r.data[uid]; ok {
		// สร้าง copy เพื่อกัน data race ภายนอก
		cpy := t
		return &cpy, nil
	}

	return nil, gorm.ErrRecordNotFound
}
