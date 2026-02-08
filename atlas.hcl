data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./internal/adapters/repository/postgre/loader/main.go",
  ]
}

data "composite_schema" "app" {
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/000000000000_setup.sql"
  }
  schema "public" {
    url = data.external_schema.gorm.url
  }
  schema "public" {
    url = "file://internal/adapters/repository/postgre/migrations/20260207184614_log_profile_changes.sql"
  }
}

env "local" {
  # 2. Reference the composite schema instead of the array
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