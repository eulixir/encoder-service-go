package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	fmt.Print(video)

	repositories.VideoRepositoryDb{Db: db}.Insert(video)

	output := "banana"
	status := "started"

	job, err := domain.NewJob(output, status, video)
	require.Nil(t, err)

	repo := repositories.JobRepositoryDb{Db: db}
	repo.Insert(job)

	j, err := repo.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Nil(t, j)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}
