rm -rf ../docs
mkdir ../docs
cp assets/* ../docs/
go run render/render.go
