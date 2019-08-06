rm -rf ../docs
mkdir ../docs
mkdir ../docs/writings
cp assets/* ../docs/

# scp -r externalpages/* antonjr:~/assets/
# scp -r assets/site.css antonjr:~/assets/
go run render/render.go
