package service

import "context"

// NovelsService describes the service.
type NovelsService interface {
	// Browse the novels list
	Browse(ctx context.Context, request interface{}) (rs interface{}, err error)

	// Read a novel by primary key
	Read(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error)

	// Add a novel
	Add(ctx context.Context, request interface{}) (rs interface{}, err error)

	// Edit a novel by primary key
	Edit(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error)

	// Put the novel into trash the bin by primary key
	Trash(ctx context.Context, pk interface{}) (rs interface{}, err error)

	// Put the novel out of the trash bin by primary key
	Restore(ctx context.Context, pk interface{}) (rs interface{}, err error)

	// Destroy a novel by primary key
	Destroy(ctx context.Context, pk interface{}) (rs interface{}, err error)

	// Batch add novels
	BatchAdd(ctx context.Context, request interface{}) (rs interface{}, err error)

	// Batch edit novels
	BatchEdit(ctx context.Context, request interface{}) (rs interface{}, err error)

	// Batch put the novel into the trash bin by primary keys
	BatchTrash(ctx context.Context, pks interface{}) (rs interface{}, err error)

	// Batch put the novel out of the trash bin by primary keys
	BatchRestore(ctx context.Context, pks interface{}) (rs interface{}, err error)

	// Batch destroy novels by the primary keys
	BatchDestroy(ctx context.Context, pks interface{}) (rs interface{}, err error)
}
