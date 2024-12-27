package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthors_Create_Success(t *testing.T) {
	req := dto.AuthorsCreateReq {
		Fullname: util.RandomStringAlpha(8),
		Gender: "m",
		BirthDateStr: "2006-01-02",
	}

	w := doTest("POST", server.RootAuthors, req, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 201, w.Code)

	var resp dto.SuccessResponse[dto.AuthorsCreateResp]

	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if assert.Nil(t, err) {
		data := resp.Data
		assert.Greater(t,data.ID, 0)
		assert.Equal(t, req.Fullname, data.Fullname)
		assert.Equal(t, req.Gender, data.Gender)

		item, _ := authorsRepo.GetByID(uint(data.ID))
		if assert.NotNil(t, item) {
			assert.Equal(t, req.Fullname, item.Fullname)
			assert.Equal(t, req.Gender, string(*item.Gender))
		}

		tdate, _ := time.Parse("2006-01-02", "2007-01-01")
		payload := dto.AuthorsUpdateReq{
			ID: item.ID,
			Fullname: util.RandomStringAlpha(8),
			Gender: "f",
			BirthDate: tdate,
		}

		editResp := authorsRepo.Update(&payload)
		assert.Equal(t, nil, editResp)

		deleteResp := authorsRepo.Delete(item.ID)
		assert.Equal(t, nil, deleteResp)
	}
}