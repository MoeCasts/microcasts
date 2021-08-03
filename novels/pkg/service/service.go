package service

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang/protobuf/ptypes"
)

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

type basicNovelsService struct {
	db *gorm.DB
}

func (b *basicNovelsService) Browse(ctx context.Context, request interface{}) (rs interface{}, err error) {
	data := []*pb.Novel{
		{
			Id:        1,
			Uuid:      uuid.New().String(),
			Name:      "我的青春恋爱物语果然有问题",
			Slug:      "oregairu",
			Author:    "渡航",
			Summary:   "我的本命番剧～",
			Cover:     "https://rin.linovel.net/article_cover/1049857_837d7b4fd14145e3818b48d36a53f0e2.jpg",
			Finished:  true,
			State:     pb.State_PUBLISHED,
			CreatedAt: ptypes.TimestampNow(),
			UpdatedAt: ptypes.TimestampNow(),
		},
		{
			Id:        2,
			Uuid:      uuid.New().String(),
			Name:      "化物语",
			Slug:      "Bakemonogatari",
			Author:    "西尾维新",
			Summary:   "高中三年级学生・阿良良木历在5月的某天，在学校意外得知同班两年、从未对话的同学战场原黑仪的秘密。那就是她身上几乎没有体重。",
			Cover:     "https://upload.wikimedia.org/wikipedia/zh/5/54/%E5%8C%96%E7%89%A9%E8%AA%9ELogo1.png",
			Finished:  true,
			State:     pb.State_PUBLISHED,
			CreatedAt: ptypes.TimestampNow(),
			UpdatedAt: ptypes.TimestampNow(),
		},
	}

	req := request.(*pb.BrowseRequest)
	page := req.Page
	if page == 0 {
		page = 1
	}

	paginate := req.Paginate
	if paginate == 0 {
		paginate = 2
	}

	from := (page - 1) * paginate
	to := page * paginate
	rs = data[from:to]

	return rs, err
}
func (b *basicNovelsService) Read(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Read
	return rs, err
}
func (b *basicNovelsService) Add(ctx context.Context, request interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Add
	return rs, err
}
func (b *basicNovelsService) Edit(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Edit
	return rs, err
}
func (b *basicNovelsService) Trash(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Trash
	return rs, err
}
func (b *basicNovelsService) Restore(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Restore
	return rs, err
}
func (b *basicNovelsService) Destroy(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of Destroy
	return rs, err
}
func (b *basicNovelsService) BatchAdd(ctx context.Context, request interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of BatchAdd
	return rs, err
}
func (b *basicNovelsService) BatchEdit(ctx context.Context, request interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of BatchEdit
	return rs, err
}
func (b *basicNovelsService) BatchTrash(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of BatchTrash
	return rs, err
}
func (b *basicNovelsService) BatchRestore(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of BatchRestore
	return rs, err
}
func (b *basicNovelsService) BatchDestroy(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	// TODO implement the business logic of BatchDestroy
	return rs, err
}

// NewBasicNovelsService returns a naive, stateless implementation of NovelsService.
func NewBasicNovelsService() NovelsService {
	dsn := "host=localhost user=default password=secret dbname=microcasts port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&pb.NovelORM{})

	return &basicNovelsService{
		db: db,
	}
}

// New returns a NovelsService with all of the expected middleware wired in.
func New(middleware []Middleware) NovelsService {
	var svc NovelsService = NewBasicNovelsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
