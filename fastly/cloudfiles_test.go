package fastly

import "testing"

func TestClient_Cloudfiles(t *testing.T) {
	t.Parallel()

	var err error
	var tv *Version
	record(t, "cloudfiles/version", func(c *Client) {
		tv = testVersion(t, c)
	})

	// Create
	var cloudfiles *Cloudfiles
	record(t, "cloudfiles/create", func(c *Client) {
		cloudfiles, err = c.CreateCloudfiles(&CreateCloudfilesInput{
			Service:         testServiceID,
			Version:         tv.Number,
			Name:            String("test-cloudfiles"),
			User:            String("user"),
			AccessKey:       String("secret-key"),
			BucketName:      String("bucket-name"),
			Path:            String("/path"),
			Region:          String("DFW"),
			Period:          Uint(12),
			GzipLevel:       Uint(9),
			Format:          String("format"),
			FormatVersion:   Uint(1),
			TimestampFormat: String("%Y"),
			MessageType:     String("classic"),
			Placement:       String("waf_debug"),
			PublicKey:       String(pgpPublicKey()),
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	// Ensure deleted
	defer func() {
		record(t, "cloudfiles/cleanup", func(c *Client) {
			c.DeleteCloudfiles(&DeleteCloudfilesInput{
				Service: testServiceID,
				Version: tv.Number,
				Name:    "test-cloudfiles",
			})

			c.DeleteCloudfiles(&DeleteCloudfilesInput{
				Service: testServiceID,
				Version: tv.Number,
				Name:    "new-test-cloudfiles",
			})
		})
	}()

	if cloudfiles.Name != "test-cloudfiles" {
		t.Errorf("bad name: %q", cloudfiles.Name)
	}
	if cloudfiles.User != "user" {
		t.Errorf("bad user: %q", cloudfiles.User)
	}
	if cloudfiles.BucketName != "bucket-name" {
		t.Errorf("bad bucket_name: %q", cloudfiles.BucketName)
	}
	if cloudfiles.AccessKey != "secret-key" {
		t.Errorf("bad access_key: %q", cloudfiles.AccessKey)
	}
	if cloudfiles.Path != "/path" {
		t.Errorf("bad path: %q", cloudfiles.Path)
	}
	if cloudfiles.Region != "DFW" {
		t.Errorf("bad region: %q", cloudfiles.Region)
	}
	if cloudfiles.Period != 12 {
		t.Errorf("bad period: %q", cloudfiles.Period)
	}
	if cloudfiles.GzipLevel != 9 {
		t.Errorf("bad gzip_level: %q", cloudfiles.GzipLevel)
	}
	if cloudfiles.Format != "format" {
		t.Errorf("bad format: %q", cloudfiles.Format)
	}
	if cloudfiles.FormatVersion != 1 {
		t.Errorf("bad format_version: %q", cloudfiles.FormatVersion)
	}
	if cloudfiles.TimestampFormat != "%Y" {
		t.Errorf("bad timestamp_format: %q", cloudfiles.TimestampFormat)
	}
	if cloudfiles.MessageType != "classic" {
		t.Errorf("bad message_type: %q", cloudfiles.MessageType)
	}
	if cloudfiles.Placement != "waf_debug" {
		t.Errorf("bad placement: %q", cloudfiles.Placement)
	}
	if cloudfiles.PublicKey != pgpPublicKey() {
		t.Errorf("bad public_key: %q", cloudfiles.PublicKey)
	}

	// List
	var lc []*Cloudfiles
	record(t, "cloudfiles/list", func(c *Client) {
		lc, err = c.ListCloudfiles(&ListCloudfilesInput{
			Service: testServiceID,
			Version: tv.Number,
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(lc) < 1 {
		t.Errorf("bad cloudfiles: %v", lc)
	}

	// Get
	var ncloudfiles *Cloudfiles
	record(t, "cloudfiles/get", func(c *Client) {
		ncloudfiles, err = c.GetCloudfiles(&GetCloudfilesInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-cloudfiles",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if cloudfiles.Name != ncloudfiles.Name {
		t.Errorf("bad name: %q", cloudfiles.Name)
	}
	if cloudfiles.User != ncloudfiles.User {
		t.Errorf("bad user: %q", cloudfiles.User)
	}
	if cloudfiles.BucketName != ncloudfiles.BucketName {
		t.Errorf("bad bucket_name: %q", cloudfiles.BucketName)
	}
	if cloudfiles.AccessKey != ncloudfiles.AccessKey {
		t.Errorf("bad access_key: %q", cloudfiles.AccessKey)
	}
	if cloudfiles.Path != ncloudfiles.Path {
		t.Errorf("bad path: %q", cloudfiles.Path)
	}
	if cloudfiles.Region != ncloudfiles.Region {
		t.Errorf("bad region: %q", cloudfiles.Region)
	}
	if cloudfiles.Period != ncloudfiles.Period {
		t.Errorf("bad period: %q", cloudfiles.Period)
	}
	if cloudfiles.GzipLevel != ncloudfiles.GzipLevel {
		t.Errorf("bad gzip_level: %q", cloudfiles.GzipLevel)
	}
	if cloudfiles.Format != ncloudfiles.Format {
		t.Errorf("bad format: %q", cloudfiles.Format)
	}
	if cloudfiles.FormatVersion != ncloudfiles.FormatVersion {
		t.Errorf("bad format_version: %q", cloudfiles.FormatVersion)
	}
	if cloudfiles.TimestampFormat != ncloudfiles.TimestampFormat {
		t.Errorf("bad timestamp_format: %q", cloudfiles.TimestampFormat)
	}
	if cloudfiles.MessageType != ncloudfiles.MessageType {
		t.Errorf("bad message_type: %q", cloudfiles.MessageType)
	}
	if cloudfiles.Placement != ncloudfiles.Placement {
		t.Errorf("bad placement: %q", cloudfiles.Placement)
	}
	if cloudfiles.PublicKey != ncloudfiles.PublicKey {
		t.Errorf("bad public_key: %q", cloudfiles.PublicKey)
	}

	// Update
	var ucloudfiles *Cloudfiles
	record(t, "cloudfiles/update", func(c *Client) {
		ucloudfiles, err = c.UpdateCloudfiles(&UpdateCloudfilesInput{
			Service:       testServiceID,
			Version:       tv.Number,
			Name:          "test-cloudfiles",
			NewName:       String("new-test-cloudfiles"),
			Period:        Uint(0),
			GzipLevel:     Uint(0),
			FormatVersion: Uint(2),
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if ucloudfiles.Name != "new-test-cloudfiles" {
		t.Errorf("bad name: %q", ucloudfiles.Name)
	}
	if ucloudfiles.GzipLevel != 0 {
		t.Errorf("bad gzip_level: %q", ucloudfiles.GzipLevel)
	}
	if ucloudfiles.Period != 0 {
		t.Errorf("bad period: %q", ucloudfiles.Period)
	}
	if ucloudfiles.FormatVersion != 2 {
		t.Errorf("bad format_version: %q", ucloudfiles.FormatVersion)
	}

	// Delete
	record(t, "cloudfiles/delete", func(c *Client) {
		err = c.DeleteCloudfiles(&DeleteCloudfilesInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "new-test-cloudfiles",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListCloudfiles_validation(t *testing.T) {
	var err error
	_, err = testClient.ListCloudfiles(&ListCloudfilesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.ListCloudfiles(&ListCloudfilesInput{
		Service: "foo",
		Version: 0,
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_CreateCloudfiles_validation(t *testing.T) {
	var err error
	_, err = testClient.CreateCloudfiles(&CreateCloudfilesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.CreateCloudfiles(&CreateCloudfilesInput{
		Service: "foo",
		Version: 0,
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_GetCloudfiles_validation(t *testing.T) {
	var err error
	_, err = testClient.GetCloudfiles(&GetCloudfilesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetCloudfiles(&GetCloudfilesInput{
		Service: "foo",
		Version: 0,
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetCloudfiles(&GetCloudfilesInput{
		Service: "foo",
		Version: 1,
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_UpdateCloudfiles_validation(t *testing.T) {
	var err error
	_, err = testClient.UpdateCloudfiles(&UpdateCloudfilesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateCloudfiles(&UpdateCloudfilesInput{
		Service: "foo",
		Version: 0,
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateCloudfiles(&UpdateCloudfilesInput{
		Service: "foo",
		Version: 1,
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DeleteCloudfiles_validation(t *testing.T) {
	var err error
	err = testClient.DeleteCloudfiles(&DeleteCloudfilesInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteCloudfiles(&DeleteCloudfilesInput{
		Service: "foo",
		Version: 0,
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteCloudfiles(&DeleteCloudfilesInput{
		Service: "foo",
		Version: 1,
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}
