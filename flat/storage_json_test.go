package main

// To write tests for the JSON storage type, we'd need the actual db *scribble.Driver mocked.
// Otherwise we're writing actual files to disk and might have to clear out the data dir between runs,
// which is awkward.