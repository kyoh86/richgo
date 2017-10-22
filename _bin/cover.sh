echo "mode: set" > acc.cover
for import_path in `go list ./... | grep -Fve 'sample'`; do
  directory=$GOPATH/src/$import_path
  if ls $directory/*.go &> /dev/null; then
    go test -coverprofile=profile.cover $import_path
    if [ -f profile.cover ]; then
      cat profile.cover | grep -Fve "mode: set" >> acc.cover
    fi
  fi
done

rm -f profile.cover
if [ -n "${COVERALLS_TOKEN}" ]; then
  goveralls -coverprofile=acc.cover -repotoken=$COVERALLS_TOKEN -service=wercker.com
fi  
rm -f acc.cover
