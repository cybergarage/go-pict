# Changelog

# v1.0.2 (2025-11-23)
### Added
- Support for `byte` type
- `Name()`, `Type()`, and `String()` methods to Param/Type
### Changed
- Renamed `CastType()` to `CastTo()` in Elem
- Improved manual type casting with go-safecast library

# v1.0.1 (2025-09-16)
- Updated to Go 1.25
- Fixed compile warnings
- Added example tests

# v1.0.0 (2023-09-02)
- Fix interface
- Fix compile warnings

# v0.8.2 (2023-05-05)
- Fix pict.CastTo() not to return *[]uint8 values for []byte elements

# v0.8.1 (2023-04-18)
-  Add Elem::Cast*() functions to cast case elements into the specified types.
-  Add DeepEqual().

# v0.8.0 (2023-04-14)
- Initial public release
