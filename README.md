# 压缩项目

tar -czvf source.tar.gz main.py

# post json

JSON=$(cat <<EOF
{
  "name": "ly",
  "language": "python",
  "source": "$(base64 -i 0 source.tar.gz)",
"method": "GET",
"path": "function-ly",
"cpu": "2",
"memory": "512m",
}
EOF
)

# run

curl -X POST -H "Content-Type:application/json" -d "$JSON" "http://localhost:8080/function"

# delete

curl -X DELETE -H "Content-Type:application/json" -d "$JSON" "http://localhost:8080/function"
