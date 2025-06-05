``

# Changelog

## \[Unreleased v1] — Early Development Stage

* 🚧 Project just started, only the basic skeleton is ready.
* 🔤 Currently, only simple ASCII art and placeholder text are implemented.
* ⚙️ No real system information gathering code written yet.
* 🛠️ No actual functionality; just the foundation is being built.
* 📚 Documentation and config files are still drafts.
* ❗ This project is not ready for testing or use; major changes are expected.

## Planned Improvements

* Adding real system info fetching features.
* Developing cross-platform support and compatibility.
* Creating meaningful output for users.
* Implementing configuration and customization options.
* Starting testing and debugging phases.

----

## Unreleased v2

* **Project structure refactored**
* **`checker` package** added for configuration checks
* **`preferences` package** introduced for user settings
* File names and internal organization improved
* Main application moved to `src/main` directory
* bubleTea library added

---

## Unreleased v3
* Config go refactored
* automatic json generator implemented


---

## Unreleased v4

* Updated README.
* Added new **info** package.
* Added the following files to the `info` package:

    * `get_os.go`
    * `system_info.go`
* Added `get_banner.go` to the `app` package.
* Created new **info/util** subpackage.
* Added the following files to the `info/util` package:

    * `cpu_info.go`
    * `gpu_info.go`
    * `memory_info.go`
    * `local_ip.go`
    * `public_ip.go`

## Unreleasedv 5
* get_banner.go renamed to banner.go and moved to info/util package
* get_config_data.go renamed to config_data.go and moved to checker package
* Added logo color changing feature
