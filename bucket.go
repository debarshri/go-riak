package riak

import (
	"github.com/3XX0/pooly"
	storage "google.golang.org/api/storage/v1"
	"log"
	"google.golang.org/api/googleapi"
	"fmt"
	"bytes"
)

type GCloudFSClient struct {
	BucketNamePrefix string
	BucketName       string
	Gcloud           *storage.Service
	ContentType      string
	ClientSecret     string
}

// Performs a Riak Get Bucket request.
func (c *Conn) GetBucket(req *RpbGetBucketReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketReq, MsgRpbGetBucketResp, req, resp)
	return
}

// Performs a Riak Set Bucket request.
func (c *Conn) SetBucket(req *RpbSetBucketReq) error {
	return c.do(MsgRpbSetBucketReq, MsgRpbSetBucketResp, req, nil)
}

// Performs a Riak List Buckets request.
// The protobufs say that it will return multiple responses but it in fact does not.
func (c *Conn) ListBuckets(req *RpbListBucketsReq) (resp *RpbListBucketsResp, err error) {
	resp = new(RpbListBucketsResp)
	err = c.do(MsgRpbListBucketsReq, MsgRpbListBucketsResp, req, resp)
	return
}

// Performs a Riak List Keys request.
// Returns multiple list keys responses.
func (c *Conn) ListKeys(req *RpbListKeysReq) ([]*RpbListKeysResp, error) {
	var resps []*RpbListKeysResp

	if err := c.request(MsgRpbListKeysReq, req); err != nil {
		return nil, err
	}
	for {
		resp := new(RpbListKeysResp)
		if err := c.response(MsgRpbListKeysResp, resp); err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		if resp.GetDone() {
			break
		}
	}
	return resps, nil
}

func (c *Conn) ListKeysWithGcloud(req *RpbListKeysReq, gs GCloudFSClient) error {

	if err := c.request(MsgRpbListKeysReq, req); err != nil {
		return err
	}

	bucket := req.Bucket
	for {
		resp := new(RpbListKeysResp)
		if err := c.response(MsgRpbListKeysResp, resp); err != nil {
			return err
		}
		//resps = append(resps, resp)

		key := string(resp.Keys[1])
		if !gs.Exists(key) {

			log.Printf("%v key doesn't exist in gcloud", key)

			kv := RpbGetReq{Bucket: bucket, Key: []byte(key)}

			s := CreatePool()
			defer s.Close()

			con, err := s.GetConn()
			client := Client(con)

			resKV, err := client.Get(&kv)

			if err != nil {
				log.Printf("%v key not found with err %v", key, err)

			}
			log.Printf("Content length %v", len(resKV.Content[0].Value))
			objects, _ := gs.Gcloud.Objects.List("test").Do()

			objects.Items[0].Size
			data := resKV.Content[0].Value

			riakDataSize := len(data)

			fmt.Printf("Data length is %v",len(data))
			buf := bytes.NewBuffer(data)
			object := &storage.Object{Name: key}

			res, err := gs.Gcloud.Objects.Insert(gs.BucketNamePrefix+"-"+gs.BucketName, object).Media(buf).Do();
			if err == nil {
				log.Printf("Inserted Object %v to GCloud Store at location: %v", res.Name, res.SelfLink)

				if riakDataSize != res.Size {
					//retry
					return key, err

				} else {
					return key, nil
				}

			} else {
			    return key, err
			}


		} else {
			log.Printf("Key %v exist", key)

		}

		if resp.GetDone() {
			break
		}
	}
	return nil
}

func (gs GCloudFSClient) Exists(key string) bool {

	if _, err := gs.Gcloud.Objects.Get(gs.BucketNamePrefix+"-"+gs.BucketName, key).Do(); err == nil {
		return true
	} else {
		log.Print("Did not find " + key + " in Google")
		return false
	}
}

// Performs a Riak Reset Bucket request.
func (c *Conn) ResetBucket(req *RpbResetBucketReq) error {
	return c.do(MsgRpbResetBucketReq, MsgRpbResetBucketResp, req, nil)
}

// Performs a Riak Get Bucket Type request.
func (c *Conn) GetBucketType(req *RpbGetBucketTypeReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketTypeReq, MsgRpbGetBucketResp, req, resp)
	return
}

// Performs a Riak Set Bucket Type request.
func (c *Conn) SetBucketType(req *RpbSetBucketTypeReq) error {
	return c.do(MsgRpbSetBucketTypeReq, MsgRpbSetBucketResp, req, nil)
}

func CreatePool() *pooly.Service {
	conf := new(pooly.ServiceConfig)
	conf.Driver = NewDriver()

	s, _ := pooly.NewService("riak", conf)
	s.Add("localhost:8087")

	return s
}
