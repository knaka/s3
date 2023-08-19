package s3clt

import (
	"testing"
)

func Test_getSession(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		args       args
		wantRegion string
		wantBucket string
		wantKey    string
	}{
		{
			"With region",
			args{
				[]string{"regionX", "bucketX", "keyX"},
			},
			"regionX",
			"bucketX",
			"keyX",
		},
		//{
		//	"Without region",
		//	args{
		//		[]string{"bucketX", "keyX"},
		//	},
		//	"",
		//	"bucketX",
		//	"keyX",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSess, gotBucket, gotKey := getSession(tt.args.args)
			if tt.wantRegion != "" && *(gotSess.Config.Region) != tt.wantRegion {
				t.Errorf("getSession() gotRegion = %v, want %v", *(gotSess.Config.Region), tt.wantRegion)
			}
			if gotBucket != tt.wantBucket {
				t.Errorf("getSession() gotBucket = %v, want %v", gotBucket, tt.wantBucket)
			}
			if gotKey != tt.wantKey {
				t.Errorf("getSession() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
