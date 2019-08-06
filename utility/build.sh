rm -rf ../docs
mkdir ../docs
mkdir ../docs/writings
cp assets/* ../docs/
cp assets/* ../docs/writings/

# scp -r externalpages/* antonjr:~/assets/
# scp -r assets/site.css antonjr:~/assets/
go run render/render.go
