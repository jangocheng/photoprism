package photoprism

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/nsfw"
)

func TestIndex_Start(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	conf.InitializeTestData(t)

	tf := classify.New(conf.AssetsPath(), conf.TensorFlowOff())
	nd := nsfw.New(conf.NSFWModelPath())
	convert := NewConvert(conf)

	ind := NewIndex(conf, tf, nd, convert, NewFiles())
	imp := NewImport(conf, ind, convert)
	opt := ImportOptionsMove(conf.ImportPath())

	imp.Start(opt)

	indexOpt := IndexOptionsAll()
	indexOpt.Rescan = false

	ind.Start(indexOpt)
}

func TestIndex_File(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	conf.InitializeTestData(t)

	tf := classify.New(conf.AssetsPath(), conf.TensorFlowOff())
	nd := nsfw.New(conf.NSFWModelPath())
	convert := NewConvert(conf)

	ind := NewIndex(conf, tf, nd, convert, NewFiles())

	err := ind.File("xxx")
	assert.Equal(t, IndexFailed, err.Status)
}
