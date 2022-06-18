package cpp

import (
	"github.com/jiajiajia666/GoQt/core"
	"github.com/jiajiajia666/GoQt/network"
	"github.com/jiajiajia666/GoQt/qml"
)

type DiskCacheFactory struct {
	*qml.QQmlNetworkAccessManagerFactory
	mCacheSize int64
}

func NewDiskCacheFactoryWithCacheSize(cacheSize int64) *DiskCacheFactory {
	dcf := &DiskCacheFactory{qml.NewQQmlNetworkAccessManagerFactory(), cacheSize}
	dcf.ConnectCreate(dcf.create)
	return dcf
}

func (dcf *DiskCacheFactory) create(parent *core.QObject) *network.QNetworkAccessManager {
	nam := NewCachingNetworkAccessManager(parent)

	diskCache := network.NewQNetworkDiskCache(nam)
	cacheFolder := core.QStandardPaths_WritableLocation(core.QStandardPaths__CacheLocation)
	diskCache.SetCacheDirectory(cacheFolder)
	diskCache.SetMaximumCacheSize(dcf.mCacheSize)

	nam.SetCache(diskCache)

	println("installing network cache of", (dcf.mCacheSize / 1024), "KB in folder", cacheFolder)

	nam.clearIgnoredUrlsFromCache()

	return nam.QNetworkAccessManager_PTR()
}
