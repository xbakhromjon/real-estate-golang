package domain

type Currency int8
type AnnouncementStatus int8

const (
	Uzs Currency = iota
	Usd
	Rub
)
const (
	AnnouncementPublished AnnouncementStatus = iota
	AnnouncementArchived
	AnnouncementDeleted
	AnnouncementRejected
	AnnouncementSubmitted
	AnnouncementResubmitted
)

type Announcement struct {
	ID          int64
	Address     string
	Description string
	Price       float32
	Currency    Currency
	Status      AnnouncementStatus
}

type AnnouncementDetailedResponse struct {
	ID          int64
	Address     string
	Description string
	Price       float32
	Currency    Currency
}

type AnnouncementHeaderResponse struct {
	ID          int64
	Address     string
	Description string
	Price       float32
	Currency    Currency
}

type AnnouncementSummaryResponse struct {
	ID          int64
	Address     string
	Description string
	Price       float32
	Currency    Currency
}

type AnnouncementCreateRequest struct {
	Address     string
	Description string
	Price       float32
	Currency    Currency
}

type AnnouncementUpdateRequest struct {
	Address     string
	Description string
	Price       float32
	Currency    Currency
}

type AnnouncementUseCase interface {
	Create(request AnnouncementCreateRequest) (int64, error)
	Update(request AnnouncementUpdateRequest) error
	GetOne(ID int64) (AnnouncementDetailedResponse, error)
	DeleteOne(ID int64) error
	ChangeStatus(ID int64, status AnnouncementStatus) error
	FilterAll() error // pageable, complex filter
	FilterPublic() error
	Publish(userId int64, adId int64) error
	Reject(userId int64, adId int64) error
	Purchase(adId int64) error
}
