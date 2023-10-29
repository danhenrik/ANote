package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

type CommunityRepository struct {
	DBConn interfaces.DBConnection
}

func NewCommunityRepository(
	DBConn interfaces.DBConnection,
) CommunityRepository {
	return CommunityRepository{DBConn: DBConn}
}

func (this CommunityRepository) Create(community *domain.Community) *errors.AppError {
	err := this.DBConn.Exec(
		"INSERT INTO communities (id, name) VALUES ($1, $2)",
		community.Id,
		community.Name,
	)
	if err != nil {
		log.Println("[CommunityRepo] Error on insert new community:", err)
		return err
	}
	return nil
}

func (this CommunityRepository) GetAll() ([]domain.Community, *errors.AppError) {
	objType := reflect.TypeOf(domain.Community{})
	res, err := this.DBConn.QueryMultiple(objType, "SELECT * FROM communities")
	if err != nil {
		log.Println("[CommunityRepo] Error on get all communities:", err)
		return nil, err
	}

	if community, ok := res.([]domain.Community); ok {
		return community, nil
	}
	return []domain.Community{}, nil
}

func (this CommunityRepository) GetById(id string) (*domain.Community, *errors.AppError) {
	objType := reflect.TypeOf(domain.Community{})
	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM communities WHERE id = $1", id)
	if err != nil {
		log.Println("[CommunityRepo] Error on get community by id:", err)
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	if community, ok := res.(domain.Community); ok {
		return &community, nil
	}
	return nil, nil
}

func (this CommunityRepository) GetByNoteId(noteId string) ([]domain.Community, *errors.AppError) {
	objType := reflect.TypeOf(domain.Community{})
	res, err := this.DBConn.QueryMultiple(
		objType,
		`SELECT communities.* FROM  communities
		INNER JOIN community_notes	ON community_notes.community_id = communities.id
		WHERE community_notes.note_id = $1`,
		noteId,
	)

	if err != nil {
		log.Println("[CommunityRepo] Error on get communities by note id:", err)
		return nil, err
	}

	if communities, ok := res.([]domain.Community); ok {
		return communities, nil
	}
	return []domain.Community{}, nil
}

func (this CommunityRepository) GetByUserId(userId string) ([]domain.Community, *errors.AppError) {
	objType := reflect.TypeOf(domain.Community{})
	res, err := this.DBConn.QueryMultiple(
		objType,
		`SELECT communities.* FROM  communities
		INNER JOIN community_members ON community_members.community_id = communities.id
		WHERE community_members.user_id = $1`,
		userId,
	)

	if err != nil {
		log.Println("[CommunityRepo] Error on get communities by note id:", err)
		return nil, err
	}

	if communities, ok := res.([]domain.Community); ok {
		return communities, nil
	}
	return []domain.Community{}, nil
}

func (this CommunityRepository) GetMembers(communityId string) ([]domain.User, *errors.AppError) {
	return nil, nil
}

func (this CommunityRepository) CheckMember(communityId string, userId string) (bool, *errors.AppError) {
	objType := reflect.TypeOf(struct{ Exists bool }{})
	res, err := this.DBConn.QueryOne(
		objType,
		`SELECT EXISTS (
			SELECT * FROM community_members
			WHERE community_id = $1
			AND user_id = $2
		)`,
		communityId,
		userId,
	)
	if err != nil {
		log.Println("[CommunityRepo] Error on check member:", err)
		return false, err
	}
	if res == nil {
		return false, nil
	}
	return res.(struct{ Exists bool }).Exists, nil
}

func (this CommunityRepository) AddMember(communityId string, userId string) *errors.AppError {
	err := this.DBConn.Exec("INSERT INTO community_members (id, community_id, user_id) VALUES ($1, $2, $3)",
		helpers.NewUUID(),
		communityId,
		userId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (this CommunityRepository) RemoveMember(communityId string, userId string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM community_members WHERE community_id = $1 AND user_id = $2", communityId, userId)
	if err != nil {
		log.Println("[CommunityRepo] Error on delete member:", err)
		return err
	}
	return nil
}

func (this CommunityRepository) Delete(id string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM communities WHERE id = $1", id)
	if err != nil {
		log.Println("[CommunityRepo] Error on delete community:", err)
		return err
	}
	return nil
}

func (this CommunityRepository) SetBackground(communityId string, filename string) *errors.AppError {
	var fname *string = nil
	if filename != "" {
		fname = &filename
	}

	err := this.DBConn.Exec("UPDATE communities SET background = $1 WHERE id = $2", fname, communityId)
	if err != nil {
		log.Println("[CommunityRepo] Error on set background:", err)
		return err
	}
	return nil
}
