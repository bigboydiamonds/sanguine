package testbridgev1

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../../contracts/testcontracts/TestSynapseBridgeV1.sol --pkg testbridgev1 --sol-version 0.6.12 --filename testbridgev1
//go:generate go run github.com/vburenin/ifacemaker -f testbridgev1.abigen.go -s TestSynapseBridgeV1Filterer  -i ITestSynapseBridgeV1Filterer  -p testbridgev1  -o filterer_generated.go -c "autogenerated file"

// ignore this line: go:generate cannot be the last line of a file
