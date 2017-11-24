set -ex

echo "mode: set" > acc.cover

for import_path in `go list ./...`; do
  go test -coverprofile=profile.cover $import_path
  if [ -f profile.cover ]; then
    cat profile.cover | grep -Fve "mode: set" >> acc.cover
    rm -f profile.cover
  fi
done

if [ -n "${COVERALLS_TOKEN}" ]; then
  goveralls -coverprofile=acc.cover -repotoken=$COVERALLS_TOKEN -service=wercker.com
fi  
rm -f acc.cover || :
