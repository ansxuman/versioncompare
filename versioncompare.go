// Package versioncompare provides functions for comparing version strings.
package versioncompare

import (
	"strconv"
	"strings"
)

// compareVersions compares two version strings.
// Returns 1 if currentVersion is greater, -1 if requiredVersion is greater, 0 if equal, and any error encountered.
func compareVersions(currentVersion, requiredVersion string) (int, error) {
	currentParts := strings.Split(currentVersion, ".")
	requiredParts := strings.Split(requiredVersion, ".")

	for i := 0; i < len(currentParts) && i < len(requiredParts); i++ {
		currentNum, err := strconv.Atoi(currentParts[i])
		if err != nil {
			return 0, err
		}

		requiredNum, err := strconv.Atoi(requiredParts[i])
		if err != nil {
			return 0, err
		}

		if currentNum < requiredNum {
			return -1, nil
		} else if currentNum > requiredNum {
			return 1, nil
		}
	}

	// Check if any remaining parts in required version
	for i := len(currentParts); i < len(requiredParts); i++ {
		requiredNum, err := strconv.Atoi(requiredParts[i])
		if err != nil {
			return 0, err
		}

		if requiredNum > 0 {
			return -1, nil
		}
	}

	// Check if any remaining parts in current version
	for i := len(requiredParts); i < len(currentParts); i++ {
		currentNum, err := strconv.Atoi(currentParts[i])
		if err != nil {
			return 0, err
		}

		if currentNum > 0 {
			return 1, nil
		}
	}

	return 0, nil // Versions are equal
}

// IsGreater checks if currentVersion is greater than requiredVersion.
func IsGreater(currentVersion, requiredVersion string) (bool, error) {
	result, err := compareVersions(currentVersion, requiredVersion)
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// IsLess checks if currentVersion is less than requiredVersion.
func IsLess(currentVersion, requiredVersion string) (bool, error) {
	result, err := compareVersions(currentVersion, requiredVersion)
	if err != nil {
		return false, err
	}
	return result == -1, nil
}

// IsEqual checks if currentVersion is equal to requiredVersion.
func IsEqual(currentVersion, requiredVersion string) (bool, error) {
	result, err := compareVersions(currentVersion, requiredVersion)
	if err != nil {
		return false, err
	}
	return result == 0, nil
}
