gen: protodep
	./scripts/gen.sh

protodep:
	protodep up -f -c -u

run:
	go run ./cmd/*.go