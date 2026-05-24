//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what scanOverridesFor — override 스캔 (both/notNull only/nullable only/none/mismatch) 검증

package query

import "testing"

func TestScanOverridesFor(t *testing.T) {
	const pgxImport = "github.com/jackc/pgx/v5/pgtype"

	makeConfig := func(entries ...sqlcOverrideEntry) sqlcOverridesConfig {
		var cfg sqlcOverridesConfig
		type sqlEntry struct {
			Gen struct {
				Go struct {
					Overrides []sqlcOverrideEntry `yaml:"overrides"`
				} `yaml:"go"`
			} `yaml:"gen"`
		}
		var s struct {
			Gen struct {
				Go struct {
					Overrides []sqlcOverrideEntry `yaml:"overrides"`
				} `yaml:"go"`
			} `yaml:"gen"`
		}
		s.Gen.Go.Overrides = entries
		cfg.SQL = append(cfg.SQL, s)
		return cfg
	}

	makeEntry := func(dbType string, nullable bool) sqlcOverrideEntry {
		var e sqlcOverrideEntry
		e.DBType = dbType
		e.Nullable = nullable
		e.GoType.Import = pgxImport
		e.GoType.Type = "UUID"
		return e
	}

	t.Run("both overrides found", func(t *testing.T) {
		cfg := makeConfig(makeEntry("uuid", false), makeEntry("uuid", true))
		notNull, nullable := scanOverridesFor(cfg, "uuid", "pgtype", "UUID")
		if !notNull || !nullable {
			t.Errorf("expected both true, got notNull=%v nullable=%v", notNull, nullable)
		}
	})

	t.Run("only notNull found", func(t *testing.T) {
		cfg := makeConfig(makeEntry("uuid", false))
		notNull, nullable := scanOverridesFor(cfg, "uuid", "pgtype", "UUID")
		if !notNull || nullable {
			t.Errorf("expected notNull=true nullable=false, got notNull=%v nullable=%v", notNull, nullable)
		}
	})

	t.Run("only nullable found", func(t *testing.T) {
		cfg := makeConfig(makeEntry("uuid", true))
		notNull, nullable := scanOverridesFor(cfg, "uuid", "pgtype", "UUID")
		if notNull || !nullable {
			t.Errorf("expected notNull=false nullable=true, got notNull=%v nullable=%v", notNull, nullable)
		}
	})

	t.Run("no matching overrides", func(t *testing.T) {
		cfg := makeConfig(makeEntry("numeric", false))
		notNull, nullable := scanOverridesFor(cfg, "uuid", "pgtype", "UUID")
		if notNull || nullable {
			t.Errorf("expected both false, got notNull=%v nullable=%v", notNull, nullable)
		}
	})

	t.Run("empty config", func(t *testing.T) {
		var cfg sqlcOverridesConfig
		notNull, nullable := scanOverridesFor(cfg, "uuid", "pgtype", "UUID")
		if notNull || nullable {
			t.Errorf("expected both false, got notNull=%v nullable=%v", notNull, nullable)
		}
	})
}
