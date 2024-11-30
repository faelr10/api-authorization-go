package database

import (
	"fmt"
	"log"
	"sync"

	database_pkg "github.com/faelr10/api-authorization-go/pkg/infra"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CockroachDBImpl struct{}

func CockroachDB() database_pkg.IDatabase {
	return &CockroachDBImpl{}
}

var (
	db    *gorm.DB
	once  sync.Once
)

// InitDatabase inicializa a conexão com o CockroachDB usando GORM
func (c *CockroachDBImpl) InitDatabase(cfg *database_pkg.Config) (*gorm.DB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

		var err error

		db, err = gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Erro ao conectar ao CockroachDB: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Erro ao obter a instância de DB: %v", err)
		}

		if err = sqlDB.Ping(); err != nil {
			log.Fatalf("Erro ao verificar a conexão com o CockroachDB: %v", err)
		}

		log.Println("Conexão com o CockroachDB estabelecida com sucesso.")
	})
	return db, nil
}

// RunMigrations executa as migrações do banco de dados usando o GORM
func (c *CockroachDBImpl) RunMigrations(db *gorm.DB, migrationsPath string) error {

	// Se você estiver usando migrações manuais com a lib 'golang-migrate'
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("erro ao obter a instância de DB: %v", err)
	}
	
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao criar o driver de migração: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath, 
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("erro ao criar a instância de migração: %v", err)
	}

	// Executando as migrações (subir migrações)
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("erro ao executar migrações: %v", err)
	}

	log.Println("Migrações aplicadas com sucesso no CockroachDB.")
	return nil
}

// SetupDatabase inicializa e configura o banco de dados
func (c *CockroachDBImpl) SetupDatabase(cfg *database_pkg.Config) (*gorm.DB, error) {
	
	// Inicializando o banco de dados com GORM
	db, err := c.InitDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("erro ao inicializar o banco de dados: %v", err)
	}

	// Rodando as migrações
	err = c.RunMigrations(db, "file://internal/infra/database/migrations")
	if err != nil {
		return nil, fmt.Errorf("erro ao rodar migrações: %v", err)
	}

	return db, nil
}
