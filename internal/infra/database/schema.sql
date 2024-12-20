-- adicionando pluguin para uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- adicionando a tabela de endpoints
CREATE TABLE  IF NOT EXISTS "endpoints" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "service_name" text NOT NULL,
  "name" text NOT NULL UNIQUE,
  "description" text NOT NULL,
  "method" varchar(10) NOT NULL,
  "url" text NOT NULL,
  "path" text NOT NULL,
  "headers" text NOT NULL,
  "body" text NOT NULL,
  "sensitive_api" boolean NOT NULL DEFAULT false,
  "active" boolean NOT NULL DEFAULT true,
  "deleted_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)

