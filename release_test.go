package relic

import (
	"testing"

	"fmt"
)

func TestRegisterReleasesMonotonicAndUnique(t *testing.T) {
	_, err := NewHistory("Test").RecordReleases(
		Release{
			Version: parseVersion(t, "2.1.1"),
			Notes:   `Everything fixed`,
		},
		"2.1.0",
		`Everything broken`,
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.1",
		`Marvelous advances were made`,
	)
	if err == nil {
		t.Errorf("error expected")
	}

	history, err := NewHistory("Test").RecordReleases(
		Release{
			Version: parseVersion(t, "2.1.1"),
			Notes:   `Everything fixed`,
		},
		"2.1.0",
		`Everything broken`,
		"2.0.0",
		`Wonderful things were achieved`,
		"1.0.0",
		`Wonderful things were achieved`,
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.1",
		`Marvelous advances were made`,
	)
	if err != nil {
		t.Error(err)
	}
	changelog, err := history.Changelog()
	if err != nil {
		t.Error(err)
	}

	assertChangelog(t, "# Test Changelog\n## Version 2.1.1\nEverything fixed\n\n## Version 2.1.0\nEverything broken\n\n## Version 2.0.0\nWonderful things were achieved\n\n## Version 1.0.0\nWonderful things were achieved\n\n## Version 0.0.2\nWonderful things were achieved\n\n## Version 0.0.1\nMarvelous advances were made\n",
		changelog)

	// Fail gap
	_, err = NewHistory("Test").RecordReleases(
		Release{
			Version: parseVersion(t, "1.0.3"),
			Notes:   `Wonderful things were achieved`,
		},
		"0.0.2",
		`Wonderful things were achieved`,
		Release{
			Version: parseVersion(t, "0.0.1"),
			Notes:   `Marvelous advances were made`,
		},
	)
	if err == nil {
		t.Errorf("error expected")
	}

	history, err = NewHistory("Test").RecordReleases(
		Release{
			Version: parseVersion(t, "1.0.3"),
			Notes:   `Wonderful things were achieved`,
		},
		"1.0.2",
		`Hotfix`,
		"1.0.1",
		`Hotfix`,
		"1.0.0",
		`Wonderful things were achieved`,
		Release{
			Version: parseVersion(t, "0.0.1"),
			Notes:   `Marvelous advances were made`,
		},
	)
	if err != nil {
		t.Error(err)
	}
	changelog, err = history.Changelog()
	if err != nil {
		t.Error(err)
	}
	assertChangelog(t, "# Test Changelog\n## Version 1.0.3\nWonderful things were achieved\n\n## Version 1.0.2\nHotfix\n\n## Version 1.0.1\nHotfix\n\n## Version 1.0.0\nWonderful things were achieved\n\n## Version 0.0.1\nMarvelous advances were made\n",
		changelog)

	_, err = NewHistory("Test").RecordReleases(
		"0.1.3",
		`Wonderful things were achieved`,
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.1",
		`Marvelous advances were made`,
	)
	if err == nil {
		t.Errorf("error expected")
	}

	history, err = NewHistory("Test").RecordReleases(
		"0.0.3",
		`Wonderful things were achieved`,
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.1",
		`Marvelous advances were made`,
	)
	if err != nil {
		t.Error(err)
	}
	changelog, err = history.Changelog()
	if err != nil {
		t.Error(err)
	}
	assertChangelog(t, "# Test Changelog\n## Version 0.0.3\nWonderful things were achieved\n\n## Version 0.0.2\nWonderful things were achieved\n\n## Version 0.0.1\nMarvelous advances were made\n",
		changelog)

	_, err = NewHistory("Test").RecordReleases(
		"0.0.3",
		`Wonderful things were achieved`,
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.1",
	)
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = NewHistory("Test").RecordReleases(
		"0.0.2",
		`Wonderful things were achieved`,
		"0.0.3",
		`Wonderful things were achieved`,
		"0.0.1",
		`Marvelous advances were made`,
	)
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestMultipleRecordReleases(t *testing.T) {

	history, err := NewHistory("Test").RecordReleases(
		"0.1.0",
		"Basic functionality",
		"0.0.2",
		"Build scripts",
		"0.0.1",
		"Proof of concept",
	)
	if err != nil {
		t.Error(err)
	}
	changelog, err := history.Changelog()
	if err != nil {
		t.Error(err)
	}
	assertChangelog(t, "# Test Changelog\n## Version 0.1.0\nBasic functionality\n\n## Version 0.0.2\nBuild scripts\n\n## Version 0.0.1\nProof of concept\n",
		changelog)

	history1, err := history.RecordReleases(
		"1.0.0",
		"finally",
		"0.2.1",
		"Patch",
		"0.2.0",
		"Came after 0.1.0",
	)
	if err != nil {
		t.Error(err)
	}
	if history1 != history {
		fmt.Errorf("history1 and history should be a pointer to the same object")
	}

	changelog, err = history1.Changelog()
	if err != nil {
		t.Error(err)
	}
	assertChangelog(t, "# Test Changelog\n## Version 1.0.0\nfinally\n\n## Version 0.2.1\nPatch\n\n## Version 0.2.0\nCame after 0.1.0\n\n## Version 0.1.0\nBasic functionality\n\n## Version 0.0.2\nBuild scripts\n\n## Version 0.0.1\nProof of concept\n",
		changelog)

	_, err = history.RecordReleases(
		"0.1.3",
		`New newness`,
		"0.1.2",
		`Added blockchain`,
		"0.1.1",
		"Dried apricot",
	)
	if err == nil {
		t.Errorf("error expected")
	}
}

func assertChangelog(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected changelog:\n%s\n\nBut actual changelog was:\n\n%s\nActual (yankable):\n%#v",
			expected, actual, actual)
	}
}

func parseVersion(t *testing.T, versionString string) Version {
	version, err := ParseVersion(versionString)
	if err != nil {
		t.Error(err)
	}
	return version
}
