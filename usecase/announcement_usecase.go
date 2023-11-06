package usecase

import "real-estate/domain"

type announcementUseCase struct {
}

func NewAnnouncementUseCase() domain.AnnouncementUseCase {

	return announcementUseCase{}
}

func (a announcementUseCase) Create(request domain.AnnouncementCreateRequest) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) Update(request domain.AnnouncementUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) GetOne(ID int64) (domain.AnnouncementDetailedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) DeleteOne(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) ChangeStatus(ID int64, status domain.AnnouncementStatus) error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) FilterAll() error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) FilterPublic() error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) Publish(userId int64, adId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) Reject(userId int64, adId int64) error {
	//TODO implement me
	panic("implement me")
}

func (a announcementUseCase) Purchase(adId int64) error {
	//TODO implement me
	panic("implement me")
}
