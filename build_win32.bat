@echo off
cd execs
cd corevm
echo Build CoreVM
go build -o ../../build/gsvm.exe
echo done!
cd ..
cd core_service
echo Build CoreService
go build -o ../../build/core_service.exe
echo done!
cd ..
cd ..
cd bundle
echo Build BundleTool
go build -o ../build/bundle.exe
echo done!