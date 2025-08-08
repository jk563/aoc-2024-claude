# Advent of Code 2024 - Code Improvements Task List

## Low Priority Issues (Future Improvements)

### 3.3 Standardize Naming Conventions
- **Category**: Code Style
- **Priority**: Low
- **File**: `internal/day02/day02.go`
- **Issue**: Inconsistent public vs private method decisions
- **Solution**: Review and standardize visibility
- **Effort**: 30 minutes
- [x] Review method visibility decisions
- [x] Make useful helpers public where appropriate (ParseInput made private for consistency)
- [x] Update documentation for public methods
- [x] Ensure consistent naming patterns

### 3.5 Add Build Optimizations
- **Category**: Performance
- **Priority**: Low
- **Issue**: Build command doesn't use optimization flags
- **Solution**: Add optimized build targets
- **Effort**: 15 minutes
- [x] Add `build-release` target to justfile
- [x] Add optimization flags (-ldflags="-s -w" -trimpath)
- [x] Add size comparison to build output (build-compare target)
- [x] Document build optimization options

## Notes

- All improvements should maintain backward compatibility
- Performance optimizations should be validated with benchmarks, creating benchmarks first to get a standard
- New utilities should have comprehensive test coverage
