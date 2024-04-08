package checkers

import (
	"log"
	"os"
)

func CheckEnv(keys []string, suppressMessage bool) {
	for _, key := range keys {
		_, exists := os.LookupEnv(key)
		if exists {
			if !suppressMessage {
				log.Printf("✅ Found %s\n", key)
			}
			continue
		}

		log.Fatalf("❌ Did not find %s in env\n", key)
		os.Exit(1)
	}
}
