data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./internal/adapters/repository/postgre/loader/loader.go",
  ]
}

data "composite_schema" "app" {
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/000000000000_setup.sql"
  }
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/20260208041131_verification_type_enum.sql"
  }
  schema "public" {
    url = data.external_schema.gorm.url
  }
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/20260207184614_log_profile_changes.sql"
  }
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/20260208045016_audit_logs_with_partitioning.sql"
  }
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/20260210040756_renameUserProfilesTable.sql"
  }
}

env "local" {
  url = "postgres://admin:password@localhost:5432/myapp?sslmode=disable"

  src = data.composite_schema.app.url
  
  dev = "docker://postgres/17/dev?search_path=public"
  
  migration {
    dir = "file://internal/adapters/repository/postgre/migrations"
  }
  
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}