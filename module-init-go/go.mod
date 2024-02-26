module example.com/learn

go 1.22.0

// in here all the required dependency module and module's dependency
// module will be present . it may as // indirect

require (
	github.com/aasourav/crypto v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/aasourav/crypto => ../package-create-crypto
