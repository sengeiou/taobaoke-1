// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dao

// Injectors from wire.go:

func newTestDao() (*dao, func(), error) {
	logger, cleanup, err := NewLogger()
	if err != nil {
		return nil, nil, err
	}
	redis, cleanup2, err := NewRedis()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	memcache, cleanup3, err := NewMC()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, cleanup4, err := NewDB()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, cleanup5, err := NewMongo()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	orderClient, err := NewOrderClient(client, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup6, err := newDao(logger, redis, memcache, db, client, orderClient)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return daoDao, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
