echo "===== Generating Mockfile for Repository ====="
mockgen -source=./internal/repository/postgres/init.go -destination=./mock/repository/postgres/init.go
mockgen -source=./internal/repository/elasticsearch/init.go -destination=./mock/repository/elasticsearch/init.go
echo "===== Mockfile for Repository Generated ====="

echo ""

echo "===== Generating Mockfile for Usecase======="
mockgen -source=./internal/usecase/init.go -destination=./mock/usecase/init.go
echo "===== Mockfile for Usecase Generated======="

echo ""


echo "===== Generating Mockfile for Handler ====="
mockgen -source=./internal/handler/api/init.go -destination=./mock/handler/api/init.go
echo "===== Mockfile for Handler Generated =====" 
