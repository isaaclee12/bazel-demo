# bazel-demo
A demo for a simple of deployment of Bazel

Game Plan: Bazel + DevContainers Learning Project (6-8 hours)

Phase 1: Environment Setup (1 hour)
You'll create the DevContainer configuration and get your IDE working inside the container. This 
phase is about proving that you can open VSCode, have it detect the devcontainer config, rebuild the container, and land you in a fully-equipped Go and Bazel development environment. By the end, you should be able to run bazel version and go version from your integrated terminal without having installed either tool on your host machine.

Phase 2: Hello World with Bazel (1-1.5 hours)
You'll write the simplest possible Go program and create your first Bazel BUILD file to compile it. This teaches you the basic Bazel syntax and workflow. You'll learn what bazel build, bazel run, and bazel test do, and you'll see the output directories where Bazel stores artifacts. The goal here is demystifying what happens when you execute a Bazel command.


Build and Run

 bazel build //:main
 bazel run //:main
 Explore bazel-bin/ output directory

Verify

 Check that binary exists in bazel-bin/main_/main
 Run bazel clean and rebuild to see full build process


Phase 3: Dependency Graph (2 hours)
This is where it gets interesting. You'll create two utility packages (let's call them stringutils and mathutils) and have your main service import both of them. You'll write BUILD files for each package, declaring the dependency relationships. Then you'll intentionally change one utility package and observe that Bazel only rebuilds what's affected. This is where the hashing and incremental build concepts click into place.

Create Utility Packages

 Create stringutils/ directory with stringutils.go
 Create stringutils/BUILD.bazel defining go_library
 Create mathutils/ directory with mathutils.go
 Create mathutils/BUILD.bazel defining go_library

Update Main to Use Utils

 Modify main.go to import and call both utility packages
 Update root BUILD.bazel to declare dependencies on utils

Build and Observe

 Run bazel build //:main
 Check bazel-bin/ for all three targets

Test Incremental Builds

 Change a function in stringutils
 Rebuild, observe only affected targets rebuild
 Change mathutils, observe different rebuild pattern

Verify Dependency Graph
 bazel query --output=graph //:main
 Run bazel query --output=graph //:main (optional, shows visual dependency graph)

Ready for file structure and contents?

Phase 4: Breaking and Observing (2-3 hours)
This is the most valuable 
phase for interview prep. You'll deliberately mess with things to see how Bazel responds. Change a function signature in stringutils and watch the build fail because the main service is now broken. Fix it and observe the rebuild. Add print statements to see which packages are actually recompiling. Check the Bazel cache directories to see the hashed artifacts. Run bazel clean and rebuild everything to compare the time difference. This hands-on experimentation will give you the intuition you need to talk confidently about build systems.

Phase 5: Testing Integration (1 hour)
You'll add unit tests to each package and configure Bazel to run them. This shows you how testing fits into the build graph. You'll see that changing a test file doesn't trigger a rebuild of the actual code, which demonstrates Bazel's understanding of the dependency direction.


Optional

Phase 6: Remote Cache Setup (if time permits)
If you finish early and want to push further, we can set up a local remote cache server and configure Bazel to use it. This would let you experience the "cache hit from a remote source" workflow that happens at DoorDash scale.
Does this progression make sense? The key is that each 
phase builds on the previous one, and by 
Phase 4 you're not following instructions anymore—you're experimenting and discovering how the system behaves. That's what will make the concepts stick and give you real stories to tell in the interview.