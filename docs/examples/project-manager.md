# Project Manager consumer example

A Project Manager host can represent one implementation artifact whose version
uses a GitHub pull-request provider reference. Separate version-pinned links can
attach that same content to a deliverable and a review. Later pull-request or
commit revisions create new `ArtifactVersion` records; approval links and
verification for the earlier revision remain pinned. Project, deliverable,
review, and implementation are host vocabulary, not core enums.
