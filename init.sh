# shellcheck disable=SC2162
read -p 'Init service name: ' service
mkdir "$service"
	cd "$service" || exit
	mkdir api cmd config service sql
	cd api || exit
	touch api.proto data.proto
	cd ../sql || exit
	mkdir migrations queries
