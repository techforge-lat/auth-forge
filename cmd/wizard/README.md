# Hexagonal Architecture Generator

A code generator tool for creating modules following the hexagonal architecture pattern.

## Installation

```bash
go install github.com/yourusername/hexagonal-generator@latest
```

## Usage

1. Create a configuration file (e.g., `module-config.json`) with your module specification:

```json
{
  "module_name": "YourModuleName",
  "database_name": "your_database",
  "package_path": "your-project-path",
  "fields": [
    {
      "name": "FieldName",
      "type": "string",
      "is_nullable": false
    }
  ]
}
```

2. Run the generator:

```bash
hexagonal-generator -config ./module-config.json
```

## Configuration Options

### Module Name
- The name of your module (e.g., "Product", "User", "Order")
- Will be used to generate all related types and files
- Should be in PascalCase

### Database Name
- Name of your database for repository configuration
- Used in import paths and repository setup

### Package Path
- The base import path for your project
- Used to generate correct import statements

### Fields
Each field requires:
- `name`: Field name in PascalCase
- `type`: Go type for the field
- `is_nullable`: Boolean indicating if the field can be null

Supported field types:
- `string`
- `int`
- `float64`
- `bool`
- `time.Time`
- `uuid.UUID`
- Custom types (will be used as-is)

## Generated Structure

The generator creates the following structure:

```
internal/
  core/
    {module}/
      application/
        usecase.go
      domain/
        command.go
        query.go
      infrastructure/
        in/
          httprest/
            handler.go
            routes.go
        out/
          repository/
            postgres/
              psql.go
              query.go
pkg/
  di/
    {module}.go
```

## Features

Generated code includes:
- Full CRUD operations
- Request validation
- Error handling with errortrace
- Database operations with pgxscan
- Query building with sqlcraft
- Dependency injection setup
- HTTP handlers with Echo
- Repository implementation
- Transaction support
- Domain-driven design patterns

## Example Usage

1. Create a new module config:

```json
{
  "module_name": "Customer",
  "database_name": "cloud_crm",
  "package_path": "cloud-crm-backend",
  "fields": [
    {
      "name": "FirstName",
      "type": "string",
      "is_nullable": false
    },
    {
      "name": "LastName",
      "type": "string",
      "is_nullable": false
    },
    {
      "name": "Email",
      "type": "string",
      "is_nullable": false
    },
    {
      "name": "Phone",
      "type": "string",
      "is_nullable": true
    }
  ]
}
```

2. Generate the module:

```bash
hexagonal-generator -config ./customer-config.json
```

3. Add the generated routes to your main router:

```go
// in your router setup
if err := customerRoutes(server); err != nil {
    return err
}
```

4. Register the module dependencies:

```go
// in your dependency setup
if err := initCustomerDependencies(container); err != nil {
    return err
}
```

## Post-Generation Steps

After generating your module:

1. Add any custom business logic to the usecase
2. Add custom validations if needed
3. Add any additional repository methods
4. Update your database schema
5. Add the module's routes to your main router
6. Register the module's dependencies

## Best Practices

1. Keep module names singular (e.g., "Product" not "Products")
2. Use PascalCase for field names
3. Consider which fields should be nullable
4. Add custom business logic after generation
5. Add custom validations for complex rules
6. Keep modules focused and single-purpose
7. Use UUID for IDs when possible
8. Follow the established error handling patterns

## Common Field Types

```json
{
  "fields": [
    {"name": "Name", "type": "string", "is_nullable": false},
    {"name": "Age", "type": "int", "is_nullable": false},
    {"name": "Balance", "type": "float64", "is_nullable": false},
    {"name": "IsActive", "type": "bool", "is_nullable": false},
    {"name": "CreatedAt", "type": "time.Time", "is_nullable": false},
    {"name": "ExternalID", "type": "uuid.UUID", "is_nullable": true}
  ]
}
```
