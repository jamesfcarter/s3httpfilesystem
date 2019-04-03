package s3httpfilesystem_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"

	"github.com/jamesfcarter/s3httpfilesystem"
	"github.com/jamesfcarter/s3httpfilesystem/mock/mock_s3iface"
)

func setupMock(t *testing.T) (*s3httpfilesystem.Filesystem, *mock_s3iface.MockS3API, func()) {
	ctrl := gomock.NewController(t)
	mock := mock_s3iface.NewMockS3API(ctrl)
	s3f := s3httpfilesystem.NewWithS3("test", mock)
	return s3f, mock, ctrl.Finish
}

func TestReaddir(t *testing.T) {
	s3f, mock, finish := setupMock(t)
	defer finish()

	mock.EXPECT().
		ListObjects((&s3.ListObjectsInput{}).
			SetBucket("test").
			SetPrefix("testdir").
			SetDelimiter("/")).
		Return((&s3.ListObjectsOutput{}).
			SetContents([]*s3.Object{
				(&s3.Object{}).
					SetKey("testdir/testfile.tst").
					SetSize(123),
			}).
			SetCommonPrefixes([]*s3.CommonPrefix{
				(&s3.CommonPrefix{}).
					SetPrefix("testdir/subdir/"),
			}), nil)

	f, err := s3f.Open("/testdir")
	if err != nil {
		t.Fatal(err)
	}

	files, err := f.Readdir(0)
	if err != nil {
		t.Fatal(err)
	}

	expected := []struct {
		name string
		size int64
	}{
		{"subdir", 0},
		{"testfile.tst", 123},
	}

	if len(files) != len(expected) {
		t.Fatalf("unexpected number of returned files: %d\n", len(files))
	}

	for i, tf := range expected {
		t.Run(tf.name, func(t *testing.T) {
			name := files[i].Name()
			if name != tf.name {
				t.Errorf("unexpected name %s\n", name)
			}
			size := files[i].Size()
			if size != tf.size {
				t.Errorf("unexpected size %d\n", size)
			}
		})
	}

}

func TestStat(t *testing.T) {
	s3f, mock, finish := setupMock(t)
	defer finish()

	mock.EXPECT().
		ListObjects((&s3.ListObjectsInput{}).
			SetBucket("test").
			SetPrefix("testdir/testfile.tst").
			SetDelimiter("/").
			SetMaxKeys(1)).
		Return((&s3.ListObjectsOutput{}).
			SetContents([]*s3.Object{
				(&s3.Object{}).
					SetKey("testdir/testfile.tst").
					SetSize(123),
			}), nil)

	f, err := s3f.Open("/testdir/testfile.tst")
	if err != nil {
		t.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	if fi.Name() != "testfile.tst" {
		t.Errorf("unexpected name %s\n", fi.Name())
	}
	if fi.Size() != 123 {
		t.Errorf("unexpected size %d\n", fi.Size())
	}
}
