package interceptor

import (
	"log"

	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

func TestOne() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			log.Println("TestOne middleware initialized")
			next(c)
			log.Println("TestOne middleware exiting")
		}
	}
}

func TestTwo() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			log.Println("TestTwo middleware initialized")
			next(c)
			log.Println("TestTwo middleware exiting")
		}
	}
}

func TestThree() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			log.Println("TestThree middleware initialized")
			next(c)
			log.Println("TestThree middleware exiting")
		}
	}
}

func TestFour() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			log.Println("TestFour middleware initialized")
			next(c)
			log.Println("TestFour middleware exiting")
		}
	}
}

func TestFive() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			log.Println("TestFive middleware initialized")
			next(c)
			log.Println("TestFive middleware exiting")
		}
	}
}
