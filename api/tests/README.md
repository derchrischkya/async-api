# How to run the tests

1. Push your changes to a branch (e.g. feature/AmazingFeature)

- This is going to run `.github/workflows/api.yaml` which is going to run the tests

2. Run go locally

- `cd ./api/src && make all && cd ../..`
- `./api/linux_x86_64/bin/apiserver & disown`
- `sh ./api/tests/run.sh`