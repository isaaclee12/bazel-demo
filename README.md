# bazel-demo
A demo for a simple of deployment of Bazel, as learning practice.


## Build and Run
```
bazel build //:main
bazel run //:main
```

# Learning Phases

## Phase 1: Environment Setup
Set up DevContainer with Go and Bazel for reproducible development environment.

**Steps:**
- [ ] Create `.devcontainer/` folder
- [ ] Create `.devcontainer/devcontainer.json` with Go image + Bazel feature
- [ ] Open folder in VSCode
- [ ] Click "Reopen in Container" when prompted
- [ ] Verify installation: `bazel version` and `go version`

### Phase 2: Hello World with Bazel
Build and run a simple Go program with Bazel to understand basic workflow.

**Steps:**
- [ ] Create `WORKSPACE.bazel` (empty file)
- [ ] Create `MODULE.bazel` (declare module name and rules_go dependency)
- [ ] Create `main.go` with Hello World program
- [ ] Create `BUILD.bazel` with `go_binary` target
- [ ] Run `bazel build //:main`
- [ ] Run `bazel run //:main`
- [ ] Explore `bazel-bin/` output directory

### Phase 3: Dependency Graph
Create utility packages and observe how Bazel manages dependencies.

**Steps:**
- [ ] Create `stringutils/` with stringutils.go, BUILD.bazel, and tests
- [ ] Create `mathutils/` with mathutils.go, BUILD.bazel, and tests
- [ ] Update `main.go` to import both utility packages
- [ ] Update root `BUILD.bazel` with `deps` attribute
- [ ] Run `bazel build //:main`
- [ ] Run `bazel query --output=label_kind "deps(//:main)"` to verify dependency graph
- [ ] Observe all three targets in bazel-bin/

### Phase 4: Breaking and Observing
Experiment with incremental builds and caching to understand Bazel's behavior.

**Baseline Build:**
- [ ] Run `bazel clean` to clear cache
- [ ] Run `bazel build //:main` and note build time
- [ ] Run `bazel build //:main` again (should be instant - cache hit)

**Break stringutils:**
- [ ] Change `Reverse()` function signature to add a parameter
- [ ] Try `bazel build //:main` (should fail - main.go doesn't match)
- [ ] Fix main.go to match new signature
- [ ] Rebuild with `-s` flag: `bazel build -s //:main`
- [ ] Observe only stringutils + main recompile (not mathutils)

**Break mathutils:**
- [ ] Add new function `Subtract(a, b int) int` to mathutils.go
- [ ] Rebuild without using it in main.go
- [ ] Observe: mathutils rebuilds, but main doesn't (unused dependency change)

**Inspect Cache:**
- [ ] Check `bazel-bin/` for binary timestamps
- [ ] Run `bazel clean` vs `bazel clean --expunge` (compare what gets deleted)
- [ ] Use `--explain=log.txt` to see what changed between builds

**Test Selective Rebuilds:**
- [ ] Modify a test file (e.g., stringutils_test.go)
- [ ] Run `bazel build //:main`
- [ ] Observe: tests don't trigger main rebuild (separate dependency path)



# Optional Milestones (if time permits)

## Phase 5: API Features & Testing Integration
Create a REST API to simulate a DoorDash backend service. Add comprehensive tests and see how Bazel handles test targets separately from build targets.

**Create API Package:**
- [ ] Create `api/` directory with `api.go` and `BUILD.bazel`
- [ ] Implement REST handlers: `GET /health`, `POST /calculate` (uses mathutils), `POST /transform` (uses stringutils)
- [ ] Add `net/http` server logic
- [ ] Define `go_library` target for API package

**Update Main Binary:**
- [ ] Modify `main.go` to start HTTP server using api package
- [ ] Update root `BUILD.bazel` to depend on `//api`
- [ ] Run `bazel run //:main` and test endpoints with curl

**Add API Tests:**
- [ ] Create `api/api_test.go` with HTTP handler tests
- [ ] Define `go_test` target in `api/BUILD.bazel`
- [ ] Run `bazel test //api:api_test`

**Test All Packages:**
- [ ] Run `bazel test //...` to execute all tests in project
- [ ] Check `bazel-testlogs/` for detailed test outputs
- [ ] Verify test results in terminal output

**Observe Test/Build Separation:**
- [ ] Modify `stringutils_test.go` (change expected value)
- [ ] Run `bazel build //:main` - observe no rebuild triggered
- [ ] Run `bazel test //stringutils:stringutils_test` - only test runs
- [ ] Fix test, modify `stringutils.go` function
- [ ] Run `bazel test //stringutils:stringutils_test` - observe both code and test rebuild

**Test Dependency Impact:**
- [ ] Break a function in `mathutils.go`
- [ ] Run `bazel test //...` 
- [ ] Observe: mathutils tests fail, API tests fail (uses mathutils), stringutils tests pass
- [ ] Fix function, rerun tests, verify cascading success

Phase 6: Remote Cache Setup

Phase 7: Understanding Gazelle
