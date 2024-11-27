package in

import (
	"context"

	"github.com/techforge-lat/dafi/v2"
)

// UseCaseCommand is a composite interface that combines create, update, and delete operations
// for use cases following the Command pattern in Clean Architecture.
//
// Type Parameters:
//   - C: The type for creation operations (e.g., UserCreate)
//   - U: The type for update operations (e.g., UserUpdate)
//
// This interface is typically implemented by service structs that handle business logic
// for write operations in your domain.
//
// Example usage:
//
//	type UserService struct{}
//	func (s *UserService) implements UseCaseCommand[UserCreate, UserUpdate]
type UseCaseCommand[C, U any] interface {
	UseCaseCreate[C] // Embeds create operations
	UseCaseUpdate[U] // Embeds update operations
	UseCaseDelete    // Embeds delete operations
}

// UseCaseCreate defines the contract for creating new entities in the system.
// This interface represents the "C" in CRUD operations at the use case level.
//
// Type Parameters:
//   - T: The type of entity to be created (e.g., UserCreate)
//
// This interface should be implemented by services that need to handle
// the creation of new domain entities with associated business rules.
type UseCaseCreate[T any] interface {
	// Create handles the creation of a new entity, applying business rules
	// and validations before persisting it.
	//
	// Parameters:
	//   - ctx: Context for the operation, carrying deadlines, cancellation signals, etc.
	//   - entity: Pointer to the entity to be created
	//
	// Returns:
	//   - error: Any error that occurred during the creation process,
	//     including validation errors or business rule violations
	Create(ctx context.Context, entity T) error
}

// UseCaseUpdate defines the contract for updating existing entities in the system.
// This interface represents the "U" in CRUD operations at the use case level.
//
// Type Parameters:
//   - T: The type containing the update data (e.g., UserUpdate)
//
// The interface allows for flexible updates using filters to identify
// which entities should be modified.
type UseCaseUpdate[T any] interface {
	// Update modifies existing entities based on provided filters and update data.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - entity: The data to update
	//   - filters: Optional set of filters to determine which entities to update
	//
	// Returns:
	//   - error: Any error during the update process, including validation errors
	Update(ctx context.Context, entity T, filters ...dafi.Filter) error
}

// UseCaseDelete defines the contract for removing entities from the system.
// This interface represents the "D" in CRUD operations at the use case level.
//
// Unlike other interfaces, this one doesn't use generics as deletion typically
// only requires identifying which entities to remove via filters.
type UseCaseDelete interface {
	// Delete removes entities that match the given filters.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - filters: Optional set of filters to determine which entities to delete
	//
	// Returns:
	//   - error: Any error during the deletion process
	Delete(ctx context.Context, filters ...dafi.Filter) error
}

// UseCaseQuery is a composite interface for read operations in the system.
// It combines single-entity and collection query operations.
//
// Type Parameters:
//   - M: The single entity model type (e.g., User)
//   - MS: The slice/collection type of the entity model (e.g., []User)
//
// This interface follows the Query part of CQRS pattern, separating
// read operations from write operations.
type UseCaseQuery[M any] interface {
	UseCaseFindOne[M]   // For single entity queries
	UseCaseFindAll[[]M] // For collection queries
}

// UseCaseQueryRelation extends UseCaseQuery to handle queries that include
// related entities. This is useful for complex domain models with relationships.
//
// Type Parameters:
//   - M: The single entity model type with relations (e.g., UserWithPosts)
//   - MS: The slice type of the entity model with relations (e.g., []UserWithPosts)
type UseCaseQueryRelation[M any] interface {
	UseCaseFindOneRelation[M]   // For single entity queries with relations
	UseCaseFindAllRelation[[]M] // For collection queries with relations
}

// UseCaseFindOne defines the contract for retrieving a single entity.
//
// Type Parameters:
//   - T: The type of entity to retrieve
type UseCaseFindOne[T any] interface {
	// FindOne retrieves a single entity matching the given criteria.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - criteria: Search criteria including filters, sorting, etc.
	//
	// Returns:
	//   - T: The found entity
	//   - error: Any error during the query process
	FindOne(ctx context.Context, criteria dafi.Criteria) (T, error)
}

// UseCaseFindAll defines the contract for retrieving multiple entities.
//
// Type Parameters:
//   - T: The collection type to retrieve (e.g., []User)
type UseCaseFindAll[T any] interface {
	// FindAll retrieves all entities matching the given criteria.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - criteria: Search criteria including filters, sorting, pagination, etc.
	//
	// Returns:
	//   - T: Collection of found entities
	//   - error: Any error during the query process
	FindAll(ctx context.Context, criteria dafi.Criteria) (T, error)
}

// UseCaseFindOneRelation defines the contract for retrieving a single entity
// along with its related entities.
//
// Type Parameters:
//   - T: The entity type with loaded relations
type UseCaseFindOneRelation[T any] interface {
	// FindOneRelation retrieves a single entity with its relations.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - criteria: Search criteria with relationship specifications
	//
	// Returns:
	//   - T: The found entity with loaded relations
	//   - error: Any error during the query process
	FindOneRelation(ctx context.Context, criteria dafi.Criteria) (T, error)
}

// UseCaseFindAllRelation defines the contract for retrieving multiple entities
// along with their related entities.
//
// Type Parameters:
//   - T: The collection type with loaded relations (e.g., []UserWithPosts)
type UseCaseFindAllRelation[T any] interface {
	// FindAllRelation retrieves multiple entities with their relations.
	//
	// Parameters:
	//   - ctx: Context for the operation
	//   - criteria: Search criteria with relationship specifications
	//
	// Returns:
	//   - T: Collection of found entities with loaded relations
	//   - error: Any error during the query process
	FindAllRelation(ctx context.Context, criteria dafi.Criteria) (T, error)
}
