// package postgres реализует доступ к данным для расчёта статистики.
// Репозиторий читает задачи с опциональной фильтрацией — доменная логика подсчёта
// статистики находится в domain.CreateStatistics, не здесь.
package postgres

import core_postgres_pool "github.com/cunofou/golang_todoapp/internal/core/repository/postgres/pool"

// StatisticsRepository — репозиторий для получения данных, необходимых для статистики.
type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

// NewStatisticsRepository создаёт репозиторий статистики с переданным пулом соединений.
func NewStatisticsRepository(
	pool core_postgres_pool.Pool,
) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
