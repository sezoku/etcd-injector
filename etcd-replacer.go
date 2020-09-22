package main

import (
	"strings"

	"github.com/ShotaKitazawa/etcd-replacer/pkg/etcdclient"
	"github.com/ShotaKitazawa/etcd-replacer/pkg/replacer"
	"github.com/ShotaKitazawa/etcd-replacer/pkg/rulesource/file"
)

func Run(c config) error {
	// generate src & dst etcd client
	srcClient, err := etcdclient.New(etcdclient.Config{
		Endpoints: c.SrcEndpoints,
	})
	if err != nil {
		return err
	}
	defer srcClient.Close()
	dstClient, err := etcdclient.New(etcdclient.Config{
		Endpoints: c.DstEndpoints,
	})
	if err != nil {
		return err
	}
	defer dstClient.Close()

	// load rules
	rules, err := file.GetRules(c.RulesFilepath)
	if err != nil {
		return err
	}

	// get values from source etcd
	keyValues, err := srcClient.LsRecursive(c.SrcDirectory)
	if err != nil {
		return err
	}

	// replace values based rule
	results, err := replacer.Replace(keyValues, rules)
	if err != nil {
		return err
	}

	// replace keys & set values to destination etcd
	for _, kv := range results {
		strings.Replace(kv.Key, c.SrcDirectory, c.DstDirectory, 1)

		err := dstClient.Put(kv)
		if err != nil {
			return err
		}
	}

	return nil
}
